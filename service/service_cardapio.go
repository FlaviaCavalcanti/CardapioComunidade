package service

import (
	"aplicacao/entity"
	"context"
)

func (s *ServiceStruct) CriarCardapio(ctx context.Context, card entity.Cardapio) (int, error) {
	return s.repo.InserirCardapio(ctx, card)
}

func (s *ServiceStruct) ListarTudoCardapio() ([]entity.Cardapio, error) {
	return s.repo.ListarTodosCardapios()
}

func (s *ServiceStruct) AtualizarCardapioS(id int, modelo entity.Cardapio) error {
	return s.repo.AtualizarCardapioR(id, modelo)
}

func (s *ServiceStruct) DeletarCardapioS(id int) error {
	return s.repo.DeletarCardapioR(id)
}

func (s *ServiceStruct) BuscarIDCardapio(id int) (*entity.Cardapio, error) {
	return s.repo.BuscarCardapioIDr(id)
}

func (s *ServiceStruct) BuscarCardapioEQtd(data string) (map[string]interface{}, error) {

	//Buscando cardapio por data
	cardapio, erro := s.repo.BuscarPorDataR(data)
	if erro != nil {
		return nil, erro
	}
	//Buscar as quantidades de almoço para o cardapio
	quantidades, erro := s.repo.BuscarQuantidadeCardapio(cardapio.ID)
	if erro != nil {
		return nil, erro
	}
	juncao := map[string]interface{}{
		"cardapio":   cardapio,
		"quantidade": quantidades,
	}

	return juncao, nil

}
