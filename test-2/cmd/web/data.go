// Filename: cmd/web/data.go
package main

import (
	"net/url"

	"github.com/abelwhite/pigstydash/internal/models"
)

type templateData struct {
	Pig             []*models.Pig
	Room            []*models.Room
	Pigsty          []*models.Pigsty
	ErrorsFromForm  map[string]string
	FormData        url.Values
	Flash           string //flash is the key
	CSRFToken       string
	IsAuthenticated bool
}
