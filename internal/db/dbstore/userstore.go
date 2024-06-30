package dbstore

import (
	"context"
	store "go-sneed/internal/db"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserStore(DB *pgxpool.Pool) store.UserStore {
	return &userRepo{
		DB,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, email string, password string) error {
    sql := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

    _, err1 := u.db.Exec(ctx, sql, "sneed", email, password)
    if err1 != nil {
        log.Printf("unable to insert row: %v", err1)
        return err1
    }
    return nil
}

func (u *userRepo) GetUser(ctx context.Context, email string) (*store.User, error) {
    sql := "SELECT * FROM users WHERE email (email) VALUES ($1)"
    log.Printf("GetUser with email: %s", email)

    rows, err := u.db.Query(context.Background(), sql, email)
    if err != nil {
        log.Printf("Query GetUser error: %v", err)
    }

    user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[store.User])
    if err != nil {
        log.Printf("Error finding user: %v", err)
        return &store.User{}, err
    }

    return &user, nil
}
