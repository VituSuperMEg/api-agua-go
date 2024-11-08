package main

import (
	"api-agua/cmd/api"
	"api-agua/internal/database"

	_ "modernc.org/sqlite"
)

func main() {
	db := database.InitDB() // inicia o banco de dados
	defer db.Close()        // fecha o banco de dados
	api.InitApi(db)         // inicia a api
}
