package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Conn() (*sql.DB, error) {
	db, erro := sql.Open("postgres", "host=172.18.0.3 port=5432 user=postgres password=p3ttinar1 database=BaseDeConhecimento sslmode=disable")
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		log.Fatal(erro)
	}
	return db, nil

}
