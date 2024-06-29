package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

func NewPostgresDB(cfgStr string) *pgxpool.Pool {
    log.Printf("db_conn config string: %s", cfgStr)
	pgxConfig, err := pgxpool.ParseConfig(cfgStr)
	if err != nil {
		panic(err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	pgxPool, err := pgxpool.NewWithConfig(context.TODO(), pgxConfig)
	if err != nil {
		panic(err)
	}

    if err := pgxPool.Ping(context.Background()); err != nil {
        panic(err)
    }
    log.Print("DB pinged successfully!")

	return pgxPool
}
