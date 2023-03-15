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

	router.HandlerFunc(http.MethodGet, "/login", app.Login)
	router.HandlerFunc(http.MethodGet, "/signup", app.Signup)
	router.HandlerFunc(http.MethodGet, "/viewpig", app.Viewpigs)
	router.HandlerFunc(http.MethodGet, "/viewsty", app.Viewsties)
	router.HandlerFunc(http.MethodGet, "/dashboard", app.Dashboard)
	router.HandlerFunc(http.MethodGet, "/checktemp", app.Checktemp)
	router.HandlerFunc(http.MethodGet, "/checkhumidity", app.Checkhumidity)
	router.HandlerFunc(http.MethodGet, "/waterbin", app.Waterbin)
	router.HandlerFunc(http.MethodGet, "/feedbin", app.Feedbin)
	router.HandlerFunc(http.MethodGet, "/settings", app.Settings)

	// Create our server
	return router

}
