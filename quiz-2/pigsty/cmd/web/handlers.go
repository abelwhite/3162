//Filename: cmd/web/handlers.go

package main

import "net/http"

func (app *application) Viewpig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
}
