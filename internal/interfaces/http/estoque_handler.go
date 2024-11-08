package http

import (
	"api-agua/internal/application/estoque"
	model "api-agua/internal/domain/estoque"
	"encoding/json"
	"net/http"
)

type EstoqueHandler struct {
	useCase *estoque.EstoqueUseCase
}

func NewEstoqueHandler(uc *estoque.EstoqueUseCase) *EstoqueHandler {
	return &EstoqueHandler{useCase: uc}
}

func (h *EstoqueHandler) CreateEstoque(w http.ResponseWriter, r *http.Request) {
	var c model.Estoque

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.useCase.Create(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
	w.WriteHeader(http.StatusCreated)
}
