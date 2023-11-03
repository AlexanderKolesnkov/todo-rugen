package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"pilrugen.com/todorugen/pkg/models/postgresql"

	_ "github.com/lib/pq"
)

type cfg struct {
	Addr string
	Dsn  string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	tasks    *postgresql.TaskModel
}

func main() {
	// чем это отличается от иницализации app ниже
	cfg := new(cfg)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "Сетевой адрес HTTP")
	flag.StringVar(&cfg.Dsn, "dsn", "user=weblocalhost dbname=todo-base sslmode=disable", "Название SQL источника данных")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//DSN - полученный источник данных
	db, err := openDB(cfg.Dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		tasks:    &postgresql.TaskModel{DB: db},
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск веб-сервера на localhost%s", cfg.Addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
