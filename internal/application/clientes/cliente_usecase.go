package clientes

import "api-agua/internal/domain/clientes"

type ClienteUseCase struct {
	service *clientes.ClienteService
}

func NewClienteUseCase(service *clientes.ClienteService) *ClienteUseCase {
	return &ClienteUseCase{service: service}
}
func (uc *ClienteUseCase) FindAll() ([]*clientes.Cliente, error) {
	return uc.service.FindAll()
}

func (uc *ClienteUseCase) Create(cliente *clientes.Cliente) error {
	return uc.service.Create(cliente)
}
