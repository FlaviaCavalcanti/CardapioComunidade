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

	// Buscando cardapio por data
	cardapio, erro := s.repo.BuscarPorDataR(data) // Agora cardapio pode ser nil
	if erro != nil {
		return nil, erro // Se for um erro real do banco de dados
	}

	if cardapio == nil { // Se o cardapio não foi encontrado para a data
		return map[string]interface{}{
			"cardapio":   nil,                         // Explicitamente nil
			"quantidade": []entity.QuantidadeAlmoco{}, // Retorna lista vazia de quantidades
		}, nil // E nenhum erro
	}

	// Buscar as quantidades de almoço para o cardapio
	quantidades, erro := s.repo.BuscarQuantidadeCardapio(cardapio.ID)
	if erro != nil {
		// Se houver um erro ao buscar quantidades para um cardapio encontrado,
		// podemos retornar o cardapio com quantidades vazias ou retornar o erro.
		// Para este caso, vamos retornar o cardapio com quantidades vazias para não bloquear a exibição do cardapio.
		return map[string]interface{}{
			"cardapio":   cardapio,
			"quantidade": []entity.QuantidadeAlmoco{}, // Retorna lista vazia em caso de erro na busca de quantidades
		}, nil
	}

	juncao := map[string]interface{}{
		"cardapio":   cardapio,
		"quantidade": quantidades,
	}

	return juncao, nil
}
