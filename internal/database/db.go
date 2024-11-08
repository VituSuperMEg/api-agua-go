package database

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "../internal/database/db.db")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS estoque (
			id TEXT PRIMARY KEY,
			marca TEXT,
			quantidade TEXT UNIQUE,
			valor_unitario TEXT,
			preco_de_revenda TEXT
	)
	`)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
