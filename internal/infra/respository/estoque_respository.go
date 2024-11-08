package respository

import (
	"api-agua/internal/domain/estoque"
	"database/sql"

	"github.com/google/uuid"
)

type EstoqueRespositorySqlite struct {
	db *sql.DB
}

func NewEstoqueRespositorySqlite(db *sql.DB) *EstoqueRespositorySqlite {
	return &EstoqueRespositorySqlite{db: db}
}

func (r *EstoqueRespositorySqlite) FindAll() ([]*estoque.Estoque, error) {
	rows, err := r.db.Query("select * from estoque")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var estoqueList []*estoque.Estoque

	for rows.Next() {
		var e estoque.Estoque
		err := rows.Scan(&e.ID, &e.Marca, &e.PrecoDeRevenda, &e.Quantidade, &e.ValorUnitario)
		if err != nil {
			return nil, err
		}
		estoqueList = append(estoqueList, &e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return estoqueList, nil
}

func (r *EstoqueRespositorySqlite) Save(estoque *estoque.Estoque) error {
	// Gera um novo UUID para o ID do estoque
	id := uuid.New().String()

	_, err := r.db.Exec(
		"INSERT INTO estoque (id, marca, quantidade, valor_unitario, preco_de_revenda) VALUES (?,?,?,?,?)",
		id, estoque.Marca, estoque.Quantidade, estoque.ValorUnitario, estoque.PrecoDeRevenda,
	)

	return err
}
