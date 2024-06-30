package store

import "context"

type User struct {
	ID       uint
	Email    string
	Password string
    Username string
}

type UserStore interface {
	CreateUser(ctx context.Context, email string, password string) error
	GetUser(ctx context.Context, mail string) (*User, error)
}

