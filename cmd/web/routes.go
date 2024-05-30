package main

import (
	"net/http"
	"websocket/internals/handlers"

	"github.com/bmizerany/pat"
)

func routes()http.Handler{
	mux:=pat.New()

	mux.Get("/",http.HandlerFunc(handlers.Home))
	mux.Get("/ws",http.HandlerFunc(handlers.WsEndpoint))


	return mux
}