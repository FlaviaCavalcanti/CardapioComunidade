package repository

import (
	"context"
	"database/sql"
)

// Repo - Struct que guarda a conexão com o banco de dados
type Repo struct {
	DB *sql.DB
}

// NewRepository - Função que inicializa e retornar o ponteiro de Repo com o db dado.
func NewRepository(ctx context.Context, db *sql.DB) (*Repo, error) {
	return &Repo{DB: db}, nil
}
