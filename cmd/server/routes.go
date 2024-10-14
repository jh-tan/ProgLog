package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/jh-tan/proglog/internal/server"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	server := server.NewHTTPServer()
	router.HandlerFunc(http.MethodGet, "/", server.HandleConsume)
	router.HandlerFunc(http.MethodPost, "/", server.HandleProduce)
	return router
}
