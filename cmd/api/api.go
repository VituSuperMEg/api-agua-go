package api

import (
	"api-agua/internal/routes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func InitApi(db *sql.DB) {

	// Inicilizador de Rotas
	routes.InitRouterClientes(db)
	routes.InitRouterEstoque(db)
	// Iniciar o servidor
	fmt.Println("Banco de dados conectado com sucesso!")
	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
