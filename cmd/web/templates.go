package main

import (
	"html/template"
	"path/filepath"

	"github.com/joechea-aupp/snippetbox/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
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

		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		// add the template set to the map, using the name of page
		// like home.tmpl.htmk as the key
		cache[name] = ts
	}
	return cache, nil
}
