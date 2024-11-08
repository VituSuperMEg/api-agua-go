package http

import (
	"api-agua/internal/application/clientes"
	model "api-agua/internal/domain/clientes"
	"encoding/json"
	"net/http"
)

type ClienteHandler struct {
	useCase *clientes.ClienteUseCase
}

func NewClienteHandler(uc *clientes.ClienteUseCase) *ClienteHandler {
	return &ClienteHandler{useCase: uc}
}

func (h *ClienteHandler) GetAllClientes(w http.ResponseWriter, r *http.Request) {
	clientesList, err := h.useCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clientesList)
}

func (h *ClienteHandler) CreateCliente(w http.ResponseWriter, r *http.Request) {
	var c model.Cliente

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
