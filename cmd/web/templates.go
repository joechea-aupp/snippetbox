package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/joechea-aupp/snippetbox/internal/models"
)

type templateData struct {
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	CurrentYear     int
	Form            any
	Flash           string
	IsAuthenticated bool
}

// this function basically load template for each page into a memeory
// so next time when page is call, it wont load data from disk again
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// use filepath.Glob() to get slice of all filepath within the match pattern *
	// like: [ui/html/pages/home.tmpl.html ui/html/pages/show.tmpl.html ui/html/pages/create.tmpl.html]
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// extract the file name (like home.tmpl.html) from the full filepath
		// and assigned it to a variable name
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
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
