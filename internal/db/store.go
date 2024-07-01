package store

import (
	"context"
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
type UserStore interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, email string) (User, error)
    UserExists(ctx context.Context, email string, username string) bool
}

