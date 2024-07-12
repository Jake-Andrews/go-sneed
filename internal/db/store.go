package store

import (
	"context"
	"go-sneed/internal/models"
	"time"

	"github.com/google/uuid"
)

type User struct {
    ID         uuid.UUID `db:"user_id"`
    Username   string    `db:"username" validate:"required,lte=50"`
    Email      string    `db:"email" validate:"required,lte=50"`
	Password   string    `db:"password" validate:"required,gte=8,lte=50"`
    Created_at time.Time `db:"created_at"`
    Last_login time.Time `db:"last_login"`
}

type Video struct {
    ID             uuid.UUID              `db:"video_id"`
    UserID         uuid.UUID              `db:"user_id" validate:"required"`
    Title          string                 `db:"title" validate:"required"`
    Description    string                 `db:"description"`
    Duration       time.Duration          `db:"duration"`
    FilePath       string                 `db:"file_path" validate:"required"`
    ThumbnailPath  string                 `db:"thumbnail_path" validate:"required"`
    Quality        map[string]interface{} `db:"quality" validate:"required"`
    Views          int                    `db:"views"`
    Likes          int                    `db:"likes"`
    Dislikes       int                    `db:"dislikes"`
    CreatedAt      time.Time              `db:"created_at"`
}

type UserStore interface {
	CreateUser(ctx context.Context, user *models.FormData) error
	GetUser(ctx context.Context, email string) (User, error)
    UserExists(ctx context.Context, email string, username string) bool
}

type VideoStore interface {
    CreateVideo(ctx context.Context, video *Video) error
    GetVideo(ctx context.Context, videoID uuid.UUID) (Video, error)
    GetRandomVideos(ctx context.Context, numVideos int) ([]Video, error)
}
