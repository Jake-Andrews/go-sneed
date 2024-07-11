package dbstore

import (
	"context"
	"fmt"
	store "go-sneed/internal/db"
	"go-sneed/internal/hash"
    "go-sneed/internal/models"
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

// userFormData (Set fields): username string, email string, password string
func (u *userRepo) CreateUser(ctx context.Context, userFormData *models.FormData) error {
    log.Printf("CreateUser fn start (before hash): %+v\n", userFormData)
    sql := `INSERT INTO users (username, email, password) VALUES (@username, @email, @password)`

    if userExists := u.UserExists(ctx, userFormData.Email, userFormData.Username); userExists {
        return fmt.Errorf("User with email %q or username %q already exists", userFormData.Email, userFormData.Username)
    }

    hashedPassword, err := u.passwordhash.GenerateFromPassword(userFormData.Password)
    if err != nil {
        log.Printf("Error hashing users password: %v", err)
        return err
    }

    args := pgx.NamedArgs{
        "username": userFormData.Username,
        "email":    userFormData.Email,
        "password": hashedPassword,
    }

    _, err = u.db.Exec(ctx, sql, args)
    if err != nil {
        log.Printf("unable to insert row: %v", err)
        return err
    }

    userFormData.Password = ""
    return nil
}

func (u *userRepo) GetUser(ctx context.Context, email string) (store.User, error) {
    sql := "SELECT * FROM users WHERE email = @email"
    log.Printf("GetUser with email: %q", email)

    args := pgx.NamedArgs{
        "email": email,
    }

    rows, err := u.db.Query(ctx, sql, args)
    if err != nil {
        log.Printf("Query GetUser error: %v", err)
        return store.User{}, err
    }

    userFormData, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[store.User])
    if err != nil {
        log.Printf("Error finding userFormData: %v", err)
        return store.User{}, err
    }

    return userFormData, nil
}

func (u *userRepo) UserExists(ctx context.Context, email string, username string) bool {
    sql := "SELECT user_id FROM users WHERE email = @email OR username = @username"

    args := pgx.NamedArgs{
        "email":    email,
        "username": username,
    }

    rows, err := u.db.Query(ctx, sql, args)
    if err != nil {
        log.Printf("Query UserExists error: %v", err)
        return false
    }

    _, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[store.User])
    if err != nil {
        log.Printf("UserExists email %q username %q\n", email, username)
        log.Println(err)
        return false
    }
    return true
}

