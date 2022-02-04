package templating

import (
	"embed"
	"html/template"
	"io"
	reading_io "tdd/reading-io"
)

var (
	//postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
	postTemplates embed.FS
)

func Render(w io.Writer, p reading_io.Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
