package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/task", app.pageTask)
	mux.HandleFunc("/task/show", app.showTask)
	mux.HandleFunc("/task/create", app.createTask)

	return mux
}
