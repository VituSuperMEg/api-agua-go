package estoque

type Estoque struct {
	ID             string `json:"id"`
	Marca          string `json:"marca"`
	Quantidade     string `json:"quantidade"`
	ValorUnitario  string `json:"valor_unitario"`
	PrecoDeRevenda string `json:"preco_de_revenda"`
}

type EstoqueRepository interface {
	Save(estoque *Estoque) error
	FindAll() ([]*Estoque, error)
}
