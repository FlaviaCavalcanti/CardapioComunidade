package endpoint

import (
	"aplicacao/entity"
	"aplicacao/service"
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

func MakeCriarCardapio(s service.Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Cardapio)
		id, erro := s.CriarCardapio(ctx, req)
		if erro != nil {
			return nil, erro
		}
		return map[string]int{"id": id}, nil
	}
}

func MakeListarCardapios(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.ListarTudoCardapio()
	}
}

func MakeAtualizarCardapioE(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		modelo := request.(entity.Cardapio)
		erro := s.AtualizarCardapioS(modelo.ID, modelo)
		if erro != nil {
			return nil, erro
		}
		return map[string]string{"Status": "Atualizado"}, nil
	}
}

func MakeDeletarCardapio(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(int)
		erro := s.DeletarCardapioS(id)
		if erro != nil {
			return nil, erro
		}
		return map[string]string{"status": "Deletado"}, nil
	}
}

func MakeBuscarCardapioPorId(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(int)
		return s.BuscarIDCardapio(id)
	}
}

func MakeJuntandoCardQtd(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data := request.(string)
		return s.BuscarCardapioEQtd(data)

	}
}
