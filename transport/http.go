package transport

import (
	"aplicacao/endpoint"
	"aplicacao/entity"
	"aplicacao/service"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir qualquer origem. Em produção, você deve restringir isso para o domínio do seu frontend.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		// Se for uma requisição OPTIONS (pre-flight), apenas envie os cabeçalhos e saia
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Chamar o próximo manipulador na cadeia
		next.ServeHTTP(w, r)
	})
}

func NewHttpServer(ctx context.Context, s service.Service, logger kitlog.Logger) *mux.Router {

	CriarPreCardapio := kithttp.NewServer(
		endpoint.MakeCriarPreEndpoint(s, logger),
		decodeCriarPreRequest,
		encodeResponse,
	)

	ListarTodosPreCardapio := kithttp.NewServer(
		endpoint.MakeListarTodosPre(s),
		decodeVaziaRequest,
		encodeResponse,
	)

	BuscarPreCardapioID := kithttp.NewServer(
		endpoint.MakeBuscarPorID(s),
		decodeIDRequest,
		encodeResponse,
	)

	AtualizarPreCardapio := kithttp.NewServer(
		endpoint.MakeAtualizarPre(s),
		decoceAtualizarPreRequest,
		encodeResponse,
	)

	DeletarPre := kithttp.NewServer(
		endpoint.MakeDeletarPre(s),
		decodeIDRequest,
		encodeResponse,
	)

	CriarCardapio := kithttp.NewServer(
		endpoint.MakeCriarCardapio(s, logger),
		decodeCriarCardapioRequest,
		encodeResponse,
	)

	ListarTodosCardapios := kithttp.NewServer(
		endpoint.MakeListarCardapios(s),
		decodeVaziaRequest,
		encodeResponse,
	)
	AtualizarCardapio := kithttp.NewServer(
		endpoint.MakeAtualizarCardapioE(s),
		decoceAtualizarCardapioRequest,
		encodeResponse,
	)

	DeletarCardapio := kithttp.NewServer(
		endpoint.MakeDeletarCardapio(s),
		decodeIDRequest,
		encodeResponse,
	)

	BuscarCardapioID := kithttp.NewServer(
		endpoint.MakeBuscarCardapioPorId(s),
		decodeIDRequest,
		encodeResponse,
	)

	CriarQuantidade := kithttp.NewServer(
		endpoint.MakeCriarQtdAlmoco(s, logger),
		decodeCriarQtdRequest,
		encodeResponse,
	)

	BuscarQuantidades := kithttp.NewServer(
		endpoint.MakeListarQuantidades(s),
		decodeVaziaRequest,
		encodeResponse,
	)

	AtualizarQuantidades := kithttp.NewServer(
		endpoint.MakeAtualizarQuantidades(s),
		decoceAtualizarQuantidadesRequest,
		encodeResponse,
	)

	DeletarQuantidade := kithttp.NewServer(
		endpoint.MakeDeletarQuantidade(s),
		decodeIDRequest,
		encodeResponse,
	)

	BuscarCardapioQtd := kithttp.NewServer(
		endpoint.MakeJuntandoCardQtd(s),
		decodeCardapioDataRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Use(corsMiddleware)

	//Rotas Pré-cardapio
	r.Methods(http.MethodPost, http.MethodOptions).Path("/pre-cardapio").Handler(CriarPreCardapio)
	r.Methods(http.MethodGet, http.MethodOptions).Path("/pre-cardapios").Handler(ListarTodosPreCardapio)
	r.Methods(http.MethodGet, http.MethodOptions).Path("/pre-cardapio/{id}").Handler(BuscarPreCardapioID)
	r.Methods(http.MethodPut, http.MethodOptions).Path("/pre-cardapio/{id}").Handler(AtualizarPreCardapio)
	r.Methods(http.MethodDelete, http.MethodOptions).Path("/pre-cardapio/{id}").Handler(DeletarPre)

	//Rotas Cardapio
	r.Methods(http.MethodPost, http.MethodOptions).Path("/cardapio").Handler(CriarCardapio)
	r.Methods(http.MethodGet, http.MethodOptions).Path("/cardapios").Handler(ListarTodosCardapios)
	r.Methods(http.MethodGet, http.MethodOptions).Path("/cardapio/{id}").Handler(BuscarCardapioID)
	r.Methods(http.MethodGet, http.MethodOptions).Path("/cardapio-data/{data}").Handler(BuscarCardapioQtd) //Busca tudo relacionado ao cardápio por data 
	r.Methods(http.MethodPut, http.MethodOptions).Path("/cardapio/{id}").Handler(AtualizarCardapio)
	r.Methods(http.MethodDelete, http.MethodOptions).Path("/cardapio/{id}").Handler(DeletarCardapio)

	//Rotas Quantidade de Almoços
	r.Methods(http.MethodPost, http.MethodOptions).Path("/quantidade-almoco").Handler(CriarQuantidade)
	r.Methods(http.MethodGet, http.MethodOptions).Path("/quantidade-almocos").Handler(BuscarQuantidades)
	r.Methods(http.MethodPut, http.MethodOptions).Path("/quantidade-almoco/{id}").Handler(AtualizarQuantidades)
	r.Methods(http.MethodDelete, http.MethodOptions).Path("/quantidade-almoco/{id}").Handler(DeletarQuantidade)

	return r
}

// Função decode para criar um Pré-Cardápio
func decodeCriarPreRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req entity.PreCardapio
	erro := json.NewDecoder(r.Body).Decode(&req)
	return req, erro
}

// Função decode para criar um Cardápio
func decodeCriarCardapioRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req entity.Cardapio
	erro := json.NewDecoder(r.Body).Decode(&req)
	return req, erro
}
func decodeCriarQtdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req entity.QuantidadeAlmoco
	erro := json.NewDecoder(r.Body).Decode(&req)
	return req, erro
}

// Função decode vazia - Utilizada para listar todas as tarefas criadas -- GET ALL
func decodeVaziaRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

// Função decode para quando é necessário o buscar com ID - GET ID
func decodeIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, erro := strconv.Atoi(vars["id"])
	if erro != nil {
		return nil, erro
	}
	return id, nil
}

// Função decode para atualizar o Pré-Cardápio
func decoceAtualizarPreRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	var modelo entity.PreCardapio
	if err := json.NewDecoder(r.Body).Decode(&modelo); err != nil {
		return nil, err
	}

	modelo.ID = id
	return modelo, nil
}

// Função decode para atualizar o Cardápio
func decoceAtualizarCardapioRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	var modelo entity.Cardapio
	if err := json.NewDecoder(r.Body).Decode(&modelo); err != nil {
		return nil, err
	}

	modelo.ID = id
	return modelo, nil
}

// Função decode para atualizar a quantidade de almoços
func decoceAtualizarQuantidadesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	var modelo entity.QuantidadeAlmoco
	if err := json.NewDecoder(r.Body).Decode(&modelo); err != nil {
		return nil, err
	}

	modelo.ID = id
	return modelo, nil
}

func decodeCardapioDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	data := vars["data"]
	return data, nil
}

// ENCODE
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
