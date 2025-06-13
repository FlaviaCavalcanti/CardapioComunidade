package main

import (
	"aplicacao/repository"
	"aplicacao/service"
	"aplicacao/transport"
	"context"
	"errors"
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {

	var (
		httpAddr = flag.String("http", ":8101", "http listen address") //Porta que o HTTP vai escutar
		ctx      = context.Background()
	)
	//log
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	errChan := make(chan error) //criamos um canal para comunicar os erros

	banco := repository.Dbconect() // Chamando a conexão com o banco
	_ = banco                      //utilizando o banco apenas para nao ficar o warning de variavel não utilizada

	//repository
	db := repository.Dbconect()
	repositorio, erro := repository.NewRepository(ctx, db)
	if erro != nil {
		errChan <- errors.New("finalizado sem error")
		level.Error(logger).Log("exit", erro)
		os.Exit(-1)
	}

	service := service.NewService(ctx, repositorio)

	//Aqui estamos subindo o servidor HTTP em uma goroutine
	go func() {
		handler := transport.NewHttpServer(ctx, service, logger)
		logger.Log("msg", "HTTP", "addr", "8101")
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()
	level.Error(logger).Log("exit", <-errChan)

}
