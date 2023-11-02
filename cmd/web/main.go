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
	tasks    *postgresql.TasksModel
}

func main() {
	// чем это отличается от иницализации app ниже
	cfg := new(cfg)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "Сетевой адрес HTTP")
	flag.StringVar(&cfg.Dsn, "dsn", "user=weblocalhost dbname=todo-base sslmode=disable", "Название SQL источника данных")

	flag.Parse()
	// Используйте log.New() для создания логгера для записи информационных сообщений. Для этого нужно
	// три параметра: место назначения для записи логов (os.Stdout), строка
	// с префиксом сообщения (INFO или ERROR) и флаги, указывающие, какая
	// дополнительная информация будет добавлена. Обратите внимание, что флаги
	// соединяются с помощью оператора OR |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//DSN - полученный источник данных
	db, err := openDB(cfg.Dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Мы также откладываем вызов db.Close(), чтобы пул соединений был закрыт
	// до выхода из функции main().
	// Подробнее про defer: https://golangs.org/errors#defer
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		tasks:    &postgresql.TasksModel{DB: db},
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

// Функция openDB() обертывает sql.Open() и возвращает пул соединений sql.DB
// для заданной строки подключения (DSN).
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
