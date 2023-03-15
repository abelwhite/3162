//Filename: cmd/web/handlers.go

package main

import "net/http"

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login\n"))
}

func (app *application) Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign up Here\n"))
}

func (app *application) Dashboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dashboard\n"))
}

func (app *application) Viewpigs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pigs\n"))
}

func (app *application) Viewsties(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sty\n"))
}

func (app *application) Checktemp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Its really hot\n"))
}

func (app *application) Checkhumidity(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Very Humid :)\n"))
}

func (app *application) Waterbin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Water bin 1: FUll\n"))
	w.Write([]byte("Water bin 2: EMPTY\n"))
}

func (app *application) Feedbin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Feed bin 1: FUll\n"))
	w.Write([]byte("Feed bin 2: EMPTY\n"))
}

func (app *application) Settings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Settings page. Change Setting Comming Soon.\n"))
}
