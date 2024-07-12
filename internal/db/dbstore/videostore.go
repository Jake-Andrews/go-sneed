package dbstore

import (
	"context"
	"fmt"
	store "go-sneed/internal/db"
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

func (v *videoRepo) CreateVideo(ctx context.Context, Video *store.Video) error {
    log.Printf("Creating video: %v", Video)
	sql := `INSERT INTO videos (user_id, title, description, duration, file_path, thumbnail_path, quality)
            VALUES (@user_id, @title, @description, @duration, @file_path, @thumbnail_path, @quality)`
	args := pgx.NamedArgs{
		"user_id":        Video.UserID,
		"title":          Video.Title,
		"description":    Video.Description,
		"duration":       Video.Duration,
		"file_path":      Video.FilePath,
		"thumbnail_path": Video.ThumbnailPath,
		"quality":        Video.Quality,
	}

	_, err := v.db.Exec(ctx, sql, args)
	if err != nil {
        return fmt.Errorf("Error creating video: %v\nVideo: %v", err, v)
	}
    return nil
}

func (v *videoRepo) GetVideo(ctx context.Context, videoID uuid.UUID) (store.Video, error) {
	sql := `SELECT video_id, user_id, title, description, duration, file_path, thumbnail_path, quality, views, likes, dislikes, created_at
            FROM videos
            WHERE video_id = @video_id`

	args := pgx.NamedArgs{
		"video_id": videoID,
	}

	rows, err := v.db.Query(ctx, sql, args)
    if err != nil {
        return store.Video{}, fmt.Errorf("Query GetVideo error: %v", err)
    }

	video, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[store.Video])
	if err != nil {
		if err == pgx.ErrNoRows {
			return store.Video{}, fmt.Errorf("video not found")
		}
		return store.Video{}, fmt.Errorf("error collecting video row: %v", err)
	}

	return video, nil
}

    //_, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[store.User])
func (v *videoRepo) GetRandomVideos(ctx context.Context, numVideos int) ([]store.Video, error) {
	sql := `SELECT video_id, user_id, title, description, duration, file_path, thumbnail_path, quality, views, likes, dislikes, created_at
            FROM videos
            ORDER BY RANDOM()
            LIMIT @limit`

	args := pgx.NamedArgs{
		"limit": numVideos,
	}

	rows, err := v.db.Query(ctx, sql, args)
	if err != nil {
		return nil, fmt.Errorf("Error querying random videos: %v", err)
	}
	defer rows.Close()

	videos, err := pgx.CollectRows(rows, pgx.RowToStructByName[store.Video])
	if err != nil {
		return nil, fmt.Errorf("Error collecting video rows: %v", err)
	}

	return videos, nil
}
