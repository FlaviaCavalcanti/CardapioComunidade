package repository

import (
	"context"
	"database/sql"
)

type Repo struct {
	DB *sql.DB
}

func NewRepository(ctx context.Context, db *sql.DB) (*Repo, error) {
	return &Repo{DB: db}, nil
}
