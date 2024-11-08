package routes

import (
	u "api-agua/internal/application/estoque"
	"api-agua/internal/domain/estoque"
	"api-agua/internal/infra/respository"
	httpfaces "api-agua/internal/interfaces/http"
	"database/sql"
	"net/http"
)

func InitRouterEstoque(db *sql.DB) {
	repo := respository.NewEstoqueRespositorySqlite(db)
	service := estoque.NewEstoqueService(repo)
	useCase := u.NewEstoqueUseCase(service)
	handler := httpfaces.NewEstoqueHandler(useCase)

	// Definir rotas
	http.HandleFunc("/estoque-list", handler.GetAll)
	http.HandleFunc("/estoque-create", handler.CreateEstoque)
	http.HandleFunc("/estoque-update-quantity", handler.UpdateQuantity)
}
