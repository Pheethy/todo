package database

import (
	"context"
	"github/pheethy/todo/config"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	PGX  = "pgx"
	SQLX = "sqlx"
)

func DBConnect(ctx context.Context, cfg config.IDbConfig) *sqlx.DB {
	//connect
	db, err := sqlx.ConnectContext(ctx, PGX, cfg.Url())
	if err != nil {
		log.Fatalf("connect to database failed: %v", err)
	}

	db.DB.SetMaxOpenConns(cfg.MaxConns())

	return db
}
