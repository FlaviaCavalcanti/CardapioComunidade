package service

import (
	"aplicacao/entity"
	"aplicacao/repository"
	"context"
)

type Service interface {
	// =============== Pré-Cardapios ==================

	CriarPre(ctx context.Context, modelo entity.PreCardapio) (int, error)
	ListarTodosPre() ([]entity.PreCardapio, error)
	BuscarPorID(id int) (*entity.PreCardapio, error)
	AtualizarPre(id int, modelo entity.PreCardapio) error
	DeletarPre(id int) error

	// =============== Cardapios ======================

	CriarCardapio(ctx context.Context, card entity.Cardapio) (int, error)
	ListarTudoCardapio() ([]entity.Cardapio, error)
	AtualizarCardapioS(id int, modelo entity.Cardapio) error
	DeletarCardapioS(id int) error
	BuscarIDCardapio(id int) (*entity.Cardapio, error)

	// =============== Quantidade Almoços ======================

	CriarQtdS(ctx context.Context, qtd entity.QuantidadeAlmoco) (int, error)
	BuscarQuantidadesS() ([]entity.QuantidadeAlmoco, error)
	AtualizarQuantidadeS(id int, modes entity.QuantidadeAlmoco) error
	DeletarQuantidadeS(id int) error

	BuscarCardapioEQtd(data string) (map[string]interface{}, error)
}

type ServiceStruct struct {
	repo *repository.Repo
}

func NewService(ctx context.Context, repo *repository.Repo) Service {
	return &ServiceStruct{repo: repo}
}
