// Filename: cmd/web/data.go
package main

import (
	"github.com/abelwhite/3162/test-1/internal/models"
)

type templateData struct {
	Pig            []*models.Pig
	Room           *models.Room
	ErrorsFromForm map[string]string
	// FormData       url.Values
	Flash string //flash is the key
}
