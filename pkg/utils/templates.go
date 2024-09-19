package utils

import (
	"fmt"
	"html/template"
	"io"
)

func RenderTemplate(w io.Writer, name string, data any) error {
	path := fmt.Sprintf("templates/%s.html", name)
	templ, err := template.ParseFiles(
		path,
		"templates/blocks/header.html",
		"templates/blocks/footer.html",
	)
	if err != nil {
		return err
	}
	return templ.Execute(w, data)
}
