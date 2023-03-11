package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Create a multiplexer
	router := httprouter.New()
	// Create a file server
	// fileServer := http.FileServer(http.Dir("./static/"))

	// router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/v1/viewpig", app.Viewpig)
	
	// Create our server
	return router

}
