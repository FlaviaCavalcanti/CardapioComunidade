package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Dbconect() *sql.DB {
	sc := "root:rootpassword@tcp(localhost:3307)/aplicacao"

	db, err := sql.Open("mysql", sc)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com banco: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar com banco: %v", err)
	}

	fmt.Println("🛸 A Conexão com o banco foi um sucesso 🛸")

	return db
}
