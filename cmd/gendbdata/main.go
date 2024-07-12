package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go-sneed/internal/config"
	store "go-sneed/internal/db"
	"go-sneed/internal/db/dbstore"
	"go-sneed/internal/db/postgres"
	"go-sneed/internal/hash/passwordhash"
	"go-sneed/internal/models"

	"github.com/google/uuid"
)

// Set env VIDEO_FOLDER_PATH or change var below to folder path
// Give each video its own folder with a screenshot
// Generates a default user into the db then adds each video
// to the db with userID as the foreign key
var defaultVideoPath = "./static/videos"

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func createVideos(userUUID uuid.UUID) ([]store.Video, error) {
	folderPath := getEnv("VIDEO_FOLDER_PATH", defaultVideoPath)
	var videos []store.Video

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != folderPath {
			video, err := processSubfolder(path, userUUID)
			if err != nil {
				return err
			}
			videos = append(videos, video)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func processSubfolder(subfolderPath string, userUUID uuid.UUID) (store.Video, error) {
	var videoFile, thumbnailFile string
	err := filepath.Walk(subfolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
				thumbnailFile = strings.Replace(path, "./static", "http://localhost:5151/static", 1)
			} else if ext == ".mp4" || ext == ".avi" || ext == ".mov" || ext == ".mkv" {
				videoFile = strings.Replace(path, "./static", "http://localhost:5151/static", 1)
			}
		}
		return nil
	})

	if err != nil {
		return store.Video{}, err
	}

	if videoFile == "" || thumbnailFile == "" {
		return store.Video{}, fmt.Errorf("video or thumbnail file not found in subfolder: %s", subfolderPath)
	}

	video := store.Video{
		ID:             uuid.New(),
		UserID:         userUUID,
		Title:          filepath.Base(subfolderPath),
		Description:    "Description",
		Duration:       time.Duration(0),
		FilePath:       videoFile,
		ThumbnailPath:  thumbnailFile,
		Quality:        map[string]interface{}{},
		Views:          0,
		Likes:          0,
		Dislikes:       0,
		CreatedAt:      time.Now(),
	}

	return video, nil
}

// set a folder path
func main() {
    cfg := config.LoadConfig()
    db := postgres.NewPostgresDB(cfg.PG_URI)
    videoStore := dbstore.NewVideoStore(db)
    passHash := passwordhash.NewHPasswordHash()
    userStore := dbstore.NewUserStore(db, passHash)

    user := models.FormData{
        Username: "testname",
        Email: "sneed@email.com",
        Password: "testing",
    }
    if err := userStore.CreateUser(context.Background(), &user); err != nil {
        log.Fatalf("Error making user: %v\nUser: %v", err, user)
    }
    log.Printf("Created user: %v", user)
    userStruct, err := userStore.GetUser(context.Background(), "sneed@email.com")
    if err != nil {
        log.Fatal("Error getting user!")
    }

	videos, err := createVideos(userStruct.ID)
	if err != nil {
		log.Fatalf("Error finding videos: %v", err)
	}
    for i, video := range videos {
        if err := videoStore.CreateVideo(context.Background(), &video); err != nil {
            log.Fatalf("Error making video: %v\nVideo: %v", err, video)
        }
		log.Printf("Created Video #%d: %+v\n", i, video)
	}
}

