package main

import (
	"github.com/nhtron/letsgo/pkg/models"
	"html/template"
	"path/filepath"
	"time"
)

type TemplateData struct {
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	CurrentYear int
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	fileNames, err := filepath.Glob(filepath.Join(dir, "*.page.gohtml"))
	if err != nil {
		return nil, err
	}
	for _, fileName := range fileNames {
		name := filepath.Base(fileName)
		ts, err := template.New(name).Funcs(functions).ParseFiles(fileName)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.gohtml"))
		if err != nil {
			return nil, err
		}
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
