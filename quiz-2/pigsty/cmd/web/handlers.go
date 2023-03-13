//Filename: cmd/web/handlers.go

package main

import "net/http"

func (app *application) ViewPig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pig 1\n"))
}

func (app *application) ViewSties(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sty\n"))
}

func (app *application) CheckTemp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sty\n"))
}

func (app *application) CheckHumidity(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sty\n"))
}
