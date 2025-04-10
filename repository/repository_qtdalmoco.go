package repository

import (
	"aplicacao/entity"
	"context"
)

func (r *Repo) InserirQtd(ctx context.Context, qtd entity.QuantidadeAlmoco) (int, error) {
	query := `INSERT INTO quantidade_almoços (cardapio_id, local, quantidade) 
	VALUES (?, ?, ?)`

	result, err := r.DB.Exec(query, qtd.CardapioID, qtd.Local, qtd.Quantidade)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *Repo) BuscarTdQtd() ([]entity.QuantidadeAlmoco, error) {
	query := `SELECT id, cardapio_id, local, quantidade FROM quantidade_almoços`
	rows, erro := r.DB.Query(query)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var quantidades []entity.QuantidadeAlmoco
	for rows.Next() {
		var quant entity.QuantidadeAlmoco
		if erro := rows.Scan(&quant.ID, &quant.CardapioID, &quant.Local, &quant.Quantidade); erro != nil {
			return nil, erro
		}
		quantidades = append(quantidades, quant)
	}
	return quantidades, nil
}

func (r *Repo) AtualizarQtdR(id int, atualizar entity.QuantidadeAlmoco) error {
	query := `UPDATE quantidade_almoços SET cardapio_id = ?, local = ?, quantidade = ? WHERE id = ?`
	_, erro := r.DB.Exec(query, atualizar.CardapioID, atualizar.Local, atualizar.Quantidade, atualizar.ID)
	return erro
}

func (r *Repo) DeletarQtdR(id int) error {
	query := `DELETE FROM quantidade_almoços WHERE id = ?`
	_, erro := r.DB.Exec(query, id)
	return erro
}
