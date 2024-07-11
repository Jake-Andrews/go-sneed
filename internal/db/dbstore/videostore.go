package dbstore

import (
	"context"
	"fmt"
	store "go-sneed/internal/db"
	"go-sneed/internal/models"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type videoRepo struct {
	db           *pgxpool.Pool
}

func NewVideoStore(DB *pgxpool.Pool) store.VideoStore {
	return &videoRepo{
        db: DB,
	}
}

func (v *videoRepo) CreateVideo(ctx context.Context, videoData *models.VideoData) error {
    log.Printf("Creating video: %v", videoData)
	sql := `INSERT INTO videos (user_id, title, description, duration, file_path, thumbnail_path, quality)
            VALUES (@user_id, @title, @description, @duration, @file_path, @thumbnail_path, @quality)`
	args := pgx.NamedArgs{
		"user_id":        videoData.UserID,
		"title":          videoData.Title,
		"description":    videoData.Description,
		"duration":       videoData.Duration,
		"file_path":      videoData.FilePath,
		"thumbnail_path": videoData.ThumbnailPath,
		"quality":        videoData.Quality,
	}

	_, err := v.db.Exec(ctx, sql, args)
	if err != nil {
        return fmt.Errorf("Error creating video: %v\nVideo: %v", err, videoData)
	}
    return nil
}

func (v *videoRepo) GetVideo(ctx context.Context, videoID uuid.UUID) (store.Video, error) {
    return store.Video{}, nil
}
