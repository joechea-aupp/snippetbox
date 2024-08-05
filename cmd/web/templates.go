package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/joechea-aupp/snippetbox/internal/models"
	"github.com/joechea-aupp/snippetbox/ui"
)

type templateData struct {
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	CurrentYear     int
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

// this function basically load template for each page into a memeory
// so next time when page is call, it wont load data from disk again
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// use filepath.Glob() to get slice of all filepath within the match pattern *
	// like: [ui/html/pages/home.tmpl.html ui/html/pages/show.tmpl.html ui/html/pages/create.tmpl.html]
	pages, err := fs.Glob(ui.Files, "html/page/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// extract the file name (like home.tmpl.html) from the full filepath
		// and assigned it to a variable name
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		// add the template set to the map, using the name of page
		// like home.tmpl.htmk as the key
		cache[name] = ts
	}
	return cache, nil
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}
