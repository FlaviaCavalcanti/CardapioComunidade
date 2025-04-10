package repository

import (
	"aplicacao/entity"
	"context"
	"database/sql"
	"fmt"
)

func (r *Repo) InserirCardapio(ctx context.Context, card entity.Cardapio) (int, error) {
	query := `INSERT INTO cardapios (data, descricao, cho, lip, ptn, kcal, peso_prato)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := r.DB.Exec(query, card.Data, card.Descricao, card.CHO, card.LIP, card.PTN, card.KCAL, card.PesoPrato)
	if err != nil {
		return 0, fmt.Errorf("eRRO AQUI: %w", err)
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *Repo) ListarTodosCardapios() ([]entity.Cardapio, error) {
	query := `SELECT id, data, descricao, cho, lip, ptn, kcal, peso_prato FROM cardapios`
	rows, erro := r.DB.Query(query)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var cardapios []entity.Cardapio
	for rows.Next() {
		var card entity.Cardapio
		if erro := rows.Scan(&card.ID, &card.Data, &card.Descricao, &card.CHO, &card.LIP, &card.PTN, &card.KCAL, &card.PesoPrato); erro != nil {
			return nil, erro
		}
		cardapios = append(cardapios, card)
	}
	return cardapios, nil
}

func (r *Repo) AtualizarCardapioR(id int, atuali entity.Cardapio) error {
	query := `UPDATE cardapios SET descricao = ?, cho = ?, lip = ?, ptn = ?, kcal = ?, peso_prato = ? WHERE id = ?`
	_, erro := r.DB.Exec(query, atuali.Descricao, atuali.CHO, atuali.LIP, atuali.PTN, atuali.KCAL, atuali.PesoPrato, id)
	return erro
}

func (r *Repo) DeletarCardapioR(id int) error {
	query := `DELETE FROM cardapios WHERE id = ?`
	_, erro := r.DB.Exec(query, id)
	return erro
}

func (r *Repo) BuscarCardapioIDr(id int) (*entity.Cardapio, error) {
	query := `SELECT id, data, descricao, cho, lip, ptn, kcal, peso_prato FROM cardapios WHERE id = ?`
	row := r.DB.QueryRow(query, id)

	var card entity.Cardapio

	if erro := row.Scan(&card.ID, &card.Data, &card.Descricao, &card.CHO, &card.LIP, &card.PTN, &card.KCAL, &card.PesoPrato); erro != nil {
		if erro == sql.ErrNoRows {
			return nil, nil
		}
		return nil, erro
	}
	return &card, nil
}

func (r *Repo) BuscarPorDataR(data string) (entity.Cardapio, error) {
	query := `SELECT id, data, descricao, cho, lip, ptn, kcal, peso_prato FROM cardapios WHERE data = ?`
	row := r.DB.QueryRow(query, data)

	var carda entity.Cardapio

	erro := row.Scan(&carda.ID, &carda.Data, &carda.Descricao, &carda.CHO, &carda.LIP, &carda.PTN, &carda.KCAL, &carda.PesoPrato)
	if erro != nil {
		return entity.Cardapio{}, erro
	}
	return carda, nil
}

func (r *Repo) BuscarQuantidadeCardapio(cardapioID int) ([]entity.QuantidadeAlmoco, error) {
	query := `SELECT id, cardapio_id, local, quantidade FROM quantidade_almoços WHERE cardapio_id = ?`
	rows, erro := r.DB.Query(query, cardapioID)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var quantidads []entity.QuantidadeAlmoco

	for rows.Next() {
		var qtdTemp entity.QuantidadeAlmoco

		if erro := rows.Scan(&qtdTemp.ID, &qtdTemp.CardapioID, &qtdTemp.Local, &qtdTemp.Quantidade); erro != nil {
			return nil, erro
		}
		quantidads = append(quantidads, qtdTemp)
	}

	return quantidads, nil
}

// func (r *Repo) BuscarPorIntervaloDeData(inicio, fim string) ([]entity.Cardapio), error{

// }
