package dbstore

import (
	"context"
	"fmt"
	store "go-sneed/internal/db"
	"go-sneed/internal/hash"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db           *pgxpool.Pool
    passwordhash hash.PasswordHash
}

func NewUserStore(DB *pgxpool.Pool, PasswordHash hash.PasswordHash) store.UserStore {
	return &userRepo{
        db:           DB,
        passwordhash: PasswordHash,
	}
}

// user (Set fields): username string, email string, password string
func (u *userRepo) CreateUser(ctx context.Context, user *store.User) error {
    log.Printf("CreateUser fn start (before hash): %+v\n", user)
    sql := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

    if userExists := u.UserExists(ctx, user.Email, user.Username); userExists {
        return fmt.Errorf("User with email %q or username %q already exists", user.Email, user.Username)
    }

    hashedPassword, err := u.passwordhash.GenerateFromPassword(user.Password)
	if err != nil {
        log.Printf("Error hashing users password: %v", err)
		return err
	}

    _, err1 := u.db.Exec(ctx, sql, user.Username, user.Email, hashedPassword)
    if err1 != nil {
        log.Printf("unable to insert row: %v", err1)
        return err1
    }

    user.Password = ""
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

func (u *userRepo) UserExists(ctx context.Context, email string, username string) bool {
    sql := "SELECT user_id FROM users WHERE email = $1 or username = $2"

    rows, err := u.db.Query(context.Background(), sql, email, username)
    if err != nil {
        log.Printf("Query UserExists error: %v", err)
    }

    _, err1 := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[store.User])
    if err1 != nil {
        log.Printf("UserExists email %q username %q\n", email, username)
        log.Println(err1)
        return false
    }
    return true
}
