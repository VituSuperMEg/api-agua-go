package clientes

import "errors"

type ClienteService struct {
	respository ClienteRespository
}

func NewClienteService(res ClienteRespository) *ClienteService {
	return &ClienteService{respository: res}
}

func (s *ClienteService) Create(cliente *Cliente) error {
	if cliente.Name == "" {
		return errors.New("Name is required")
	}
	return s.respository.Save(cliente)
}

func (s *ClienteService) FindAll() ([]*Cliente, error) {
	return s.respository.FindAll()
}
