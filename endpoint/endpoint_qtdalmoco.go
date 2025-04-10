package endpoint

import (
	"aplicacao/entity"
	"aplicacao/service"
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

func MakeCriarQtdAlmoco(s service.Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.QuantidadeAlmoco)
		id, erro := s.CriarQtdS(ctx, req)
		if erro != nil {
			return nil, erro
		}
		return map[string]int{"id": id}, nil
	}
}

func MakeListarQuantidades(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.BuscarQuantidadesS()
	}
}

func MakeAtualizarQuantidades(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		modelo := request.(entity.QuantidadeAlmoco)
		erro := s.AtualizarQuantidadeS(modelo.ID, modelo)
		if erro != nil {
			return nil, erro
		}
		return map[string]string{"Status": "Atualizado"}, nil
	}
}

func MakeDeletarQuantidade(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(int)
		erro := s.DeletarQuantidadeS(id)
		if erro != nil {
			return nil, erro
		}
		return map[string]string{"Status": "Deletado"}, nil
	}
}
