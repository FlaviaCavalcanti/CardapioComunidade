package service

import (
	"aplicacao/entity"
	"context"
)

func (s *ServiceStruct) CriarQtdS(ctx context.Context, qtd entity.QuantidadeAlmoco) (int, error) {
	return s.repo.InserirQtd(ctx, qtd)
}

func (s *ServiceStruct) BuscarQuantidadesS() ([]entity.QuantidadeAlmoco, error) {
	return s.repo.BuscarTdQtd()
}

func (s *ServiceStruct) AtualizarQuantidadeS(id int, modes entity.QuantidadeAlmoco) error {
	return s.repo.AtualizarQtdR(id, modes)
}

func (s *ServiceStruct) DeletarQuantidadeS(id int) error {
	return s.repo.DeletarQtdR(id)
}
