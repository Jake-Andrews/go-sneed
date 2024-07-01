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

// user (Set fields): username string, email string, password string
func (u *userRepo) CreateUser(ctx context.Context, user store.User) error {
    log.Printf("Creating user: %+v\n", user)
    sql := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

    _, err1 := u.db.Exec(ctx, sql, user.Username, user.Email, user.Password)
    if err1 != nil {
        log.Printf("unable to insert row: %v", err1)
        return err1
    }
    return nil
}

func (u *userRepo) GetUser(ctx context.Context, email string) (store.User, error) {
    sql := "SELECT * FROM users WHERE email (email) VALUES ($1)"
    log.Printf("GetUser with email: %q", email)

    rows, err := u.db.Query(context.Background(), sql, email)
    if err != nil {
        log.Printf("Query GetUser error: %v", err)
    }

    user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[store.User])
    if err != nil {
        log.Printf("Error finding user: %v", err)
        return store.User{}, err
    }

    return user, nil
}
