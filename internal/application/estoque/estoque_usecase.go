package estoque

import "api-agua/internal/domain/estoque"

type EstoqueUseCase struct {
	service *estoque.EstoqueService
}

func NewEstoqueUseCase(service *estoque.EstoqueService) *EstoqueUseCase {
	return &EstoqueUseCase{service: service}
}

func (uc *EstoqueUseCase) Create(estoque *estoque.Estoque) error {
	return uc.service.Create(estoque)
}
