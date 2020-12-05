package main

import (
	"github.com/nhtron/letsgo/pkg/utils"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir(app.config.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", utils.NoDirListing(fileServer)))
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
