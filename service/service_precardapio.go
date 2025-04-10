package service

import (
	"aplicacao/entity"
	"context"
)

func (s *ServiceStruct) CriarPre(ctx context.Context, modelo entity.PreCardapio) (int, error) {
	return s.repo.Inserir(ctx, modelo)
}

func (s *ServiceStruct) ListarTodosPre() ([]entity.PreCardapio, error) {
	return s.repo.ListarTodos()
}

func (s *ServiceStruct) BuscarPorID(id int) (*entity.PreCardapio, error) {
	return s.repo.BuscarID(id)
}

func (s *ServiceStruct) AtualizarPre(id int, modelo entity.PreCardapio) error {
	return s.repo.Atualizar(id, modelo)
}

func (s *ServiceStruct) DeletarPre(id int) error {
	return s.repo.Deletar(id)
}
