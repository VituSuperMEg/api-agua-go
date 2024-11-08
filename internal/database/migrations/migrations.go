package migrations

import (
	"database/sql"
)

func ClientesMigrations(db *sql.DB) {
	// err := db.Ping()
	// // _, err = db.Exec("DROP TABLE IF EXISTS clientes")
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// _, err = db.Exec(`
	// CREATE TABLE IF NOT EXISTS clientes (
	// 		id TEXT PRIMARY KEY,
	// 		name TEXT,
	// 		celular TEXT UNIQUE,
	// 		endereco TEXT,
	// 		rua TEXT,
	// 		numero TEXT,
	// 		bairro TEXT
	// )
	// `)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
