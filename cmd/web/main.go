package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nhtron/letsgo/pkg/models/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
)

type appConfig struct {
	addr      string
	staticDir string
	dns       string
}

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	config        *appConfig
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	//fixme: using .env file instead when have chance
	cfg := new(appConfig)
	flag.StringVar(&cfg.addr, "addr", "127.0.0.1:3000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Static dir")
	flag.StringVar(&cfg.dns, "dns", "web:golangCool$@/snippetbox?parseTime=true", "mysql connection string")
	flag.Parse()

	//logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDb(cfg.dns)
	if err != nil {
		errorLog.Fatalln(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatalln(err)
	}

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		config:        cfg,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Println("Starting server on " + cfg.addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDb(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
