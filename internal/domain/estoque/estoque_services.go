package estoque

type EstoqueService struct {
	repo EstoqueRepository
}

func NewEstoqueService(repo EstoqueRepository) *EstoqueService {
	return &EstoqueService{repo: repo}
}

func (s *EstoqueService) Create(estoque *Estoque) error {
	return s.repo.Save(estoque)
}

func (s *EstoqueService) FindAll() ([]*Estoque, error) {
	return s.repo.FindAll()
}
