package routes

import (
	u "api-agua/internal/application/clientes"
	"api-agua/internal/domain/clientes"
	"api-agua/internal/infra/respository"
	httpfaces "api-agua/internal/interfaces/http"

	"database/sql"
	"net/http"
)

func InitRouterClientes(db *sql.DB) {
	repo := respository.NewClienteRespositorySqlite(db)
	service := clientes.NewClienteService(repo)
	useCase := u.NewClienteUseCase(service)
	handler := httpfaces.NewClienteHandler(useCase)

	// Definir rotas
	http.HandleFunc("/clientes", handler.GetAllClientes)
	http.HandleFunc("/clientes-post", handler.CreateCliente)
}
