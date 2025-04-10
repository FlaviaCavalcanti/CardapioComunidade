package endpoint

import (
	"aplicacao/entity"
	"aplicacao/service"
	"context"

	"github.com/go-kit/log"

	"github.com/go-kit/kit/endpoint"
)

func MakeCriarPreEndpoint(s service.Service, logger log.Logger) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.PreCardapio)
		id, erro := s.CriarPre(ctx, req)
		if erro != nil {
			return nil, erro
		}
		return map[string]int{"id": id}, nil
	}
}

func MakeListarTodosPre(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ListarTodosPre()
	}
}

func MakeBuscarPorID(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(int)
		return s.BuscarPorID(id)
	}
}

func MakeAtualizarPre(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		modelo := request.(entity.PreCardapio)
		erro := s.AtualizarPre(modelo.ID, modelo)
		if erro != nil {
			return nil, erro
		}
		return map[string]string{"status": "atualizado"}, nil
	}
}

func MakeDeletarPre(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(int)
		erro := s.DeletarPre(id)
		if erro != nil {
			return nil, erro
		}
		return map[string]string{"status": "Deletado"}, nil
	}
}
