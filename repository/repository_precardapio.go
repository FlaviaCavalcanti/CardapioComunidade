package repository

import (
	"aplicacao/entity"
	"context"
	"database/sql"
	"fmt"
)

// Insirir um pre-cardapio no banco
func (r *Repo) Inserir(ctx context.Context, pre entity.PreCardapio) (int, error) {
	query := `INSERT INTO pre_cardapios (nome, descricao, cho, lip, ptn, kcal, peso_prato)
			 VALUES(?,?,?,?,?,?,?)`
	result, err := r.DB.Exec(query, pre.Nome, pre.Descricao, pre.CHO, pre.LIP, pre.PTN, pre.KCAL, pre.PesoPrato)
	if err != nil {
		return 0, fmt.Errorf("erro ao criar a tarefa:%w", err)
	}

	id, erro := result.LastInsertId()
	return int(id), erro
}

// Listar todos os pre-cardapios
func (r *Repo) ListarTodos() ([]entity.PreCardapio, error) {
	query := `SELECT id, nome, descricao, cho, lip, ptn, kcal, peso_prato FROM pre_cardapios`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modelos []entity.PreCardapio
	for rows.Next() {
		var pre entity.PreCardapio
		if err := rows.Scan(&pre.ID, &pre.Nome, &pre.Descricao, &pre.CHO, &pre.LIP, &pre.PTN, &pre.KCAL, &pre.PesoPrato); err != nil {
			return nil, err
		}
		modelos = append(modelos, pre)
	}

	return modelos, nil
}

// Função para listar pre-cardapios por id
func (r *Repo) BuscarID(id int) (*entity.PreCardapio, error) {
	query := `SELECT id, nome, descricao, cho, lip, ptn, kcal, peso_prato FROM pre_cardapios WHERE id = ?`
	row := r.DB.QueryRow(query, id)

	var pre entity.PreCardapio

	if erro := row.Scan(&pre.ID, &pre.Nome, &pre.Descricao, &pre.CHO, &pre.LIP, &pre.PTN, &pre.KCAL, &pre.PesoPrato); erro != nil {
		if erro == sql.ErrNoRows {
			return nil, nil
		}
		return nil, erro
	}
	return &pre, nil
}

//Função para atualizar um id

func (r *Repo) Atualizar(id int, pre entity.PreCardapio) error {
	query := `UPDATE pre_cardapios SET nome = ?, descricao = ?, cho = ?, lip = ?, ptn = ?, kcal = ?, peso_prato = ? WHERE id = ?`

	_, err := r.DB.Exec(query, pre.Nome, pre.Descricao, pre.CHO, pre.LIP, pre.PTN, pre.KCAL, pre.PesoPrato, id)
	return err
}

// Função para deletar um pre-cardapio
func (r *Repo) Deletar(id int) error {
	query := `DELETE FROM pre_cardapios WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
