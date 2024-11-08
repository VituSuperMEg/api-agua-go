package clientes

type Cliente struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Celular  string `json:"celular"`
	Endereco string `json:"endereco"`
	Rua      string `json:"rua"`
	Numero   string `json:"numero"`
	Bairro   string `json:"bairro"`
}

type ClienteRespository interface {
	Save(cliente *Cliente) error
	FindAll() ([]*Cliente, error)
}
