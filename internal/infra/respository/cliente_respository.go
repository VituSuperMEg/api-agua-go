package respository

import (
	"api-agua/internal/domain/clientes"
	"database/sql"

	"github.com/google/uuid"
)

type ClienteRespositorySqlite struct {
	db *sql.DB
}

func NewClienteRespositorySqlite(db *sql.DB) *ClienteRespositorySqlite {
	return &ClienteRespositorySqlite{db: db}
}

func (r *ClienteRespositorySqlite) Save(cliente *clientes.Cliente) error {
	// Gera um novo UUID para o ID do cliente
	id := uuid.New().String()

	_, err := r.db.Exec(
		"INSERT INTO clientes (id, name, celular, endereco, rua, numero, bairro) VALUES (?, ?, ?, ?, ?, ?, ?)",
		id, cliente.Name, cliente.Celular, cliente.Endereco, cliente.Rua, cliente.Numero, cliente.Bairro,
	)
	return err
}

func (r *ClienteRespositorySqlite) FindAll() ([]*clientes.Cliente, error) {
	rows, err := r.db.Query("SELECT * FROM clientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientesList []*clientes.Cliente

	for rows.Next() {
		var cliente clientes.Cliente
		err := rows.Scan(&cliente.ID, &cliente.Name, &cliente.Celular, &cliente.Endereco, &cliente.Rua, &cliente.Numero, &cliente.Bairro)
		if err != nil {
			return nil, err
		}
		clientesList = append(clientesList, &cliente)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clientesList, nil
}
