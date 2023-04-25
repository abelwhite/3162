// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))
	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave)
	//we wrap

	router.Handler(http.MethodGet, "/login", dynamicMiddleware.ThenFunc(app.login))
	router.Handler(http.MethodGet, "/signup", dynamicMiddleware.ThenFunc(app.signup))
	router.Handler(http.MethodGet, "/dashboard", dynamicMiddleware.ThenFunc(app.dashboard))
	router.Handler(http.MethodGet, "/setting", dynamicMiddleware.ThenFunc(app.settings))
	router.Handler(http.MethodGet, "/profile", dynamicMiddleware.ThenFunc(app.profile))

	router.Handler(http.MethodGet, "/pig/create", dynamicMiddleware.ThenFunc(app.pigCreateShow))
	router.Handler(http.MethodPost, "/pig/create", dynamicMiddleware.ThenFunc(app.pigCreateSubmit))
	router.Handler(http.MethodGet, "/pig/show", dynamicMiddleware.ThenFunc(app.pigShow))
	router.Handler(http.MethodGet, "/pig/delete", dynamicMiddleware.ThenFunc(app.pigDelete))
	router.Handler(http.MethodGet, "/pig/update", dynamicMiddleware.ThenFunc(app.pigUpdate))
	router.Handler(http.MethodPost, "/pig/update", dynamicMiddleware.ThenFunc(app.pigUpdateQuery))

	router.Handler(http.MethodGet, "/room/create", dynamicMiddleware.ThenFunc(app.roomCreateShow))
	router.Handler(http.MethodGet, "/room/show", dynamicMiddleware.ThenFunc(app.roomShow))

	router.Handler(http.MethodGet, "/pigsty/create", dynamicMiddleware.ThenFunc(app.pigstyCreateShow))
	router.Handler(http.MethodGet, "/pigsty/show", dynamicMiddleware.ThenFunc(app.pigstyShow))

	//returns to the router to our middleware interesting in between go server and mux
	//Client -> Goserver ->Middleware -> Router -> Handler

	//tidy up the middle wear
	standardMiddleware := alice.New(app.RecoverPanicMiddleware, app.logRequestMiddleware, securityHeadersMiddleware)

	return standardMiddleware.Then(router)
}
