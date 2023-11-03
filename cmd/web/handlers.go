package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pilrugen.com/todorugen/pkg/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	t, err := app.tasks.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Tasks: t}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Sever Error", 500)
	}

	w.Write([]byte("Home page"))
}

func (app *application) pageTask(w http.ResponseWriter, r *http.Request) {
	t, err := app.tasks.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, task := range t {
		fmt.Fprintf(w, "%v\n", task)
	}
	w.Write([]byte("Tasks page"))
}

func (app *application) showTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	t, err := app.tasks.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := &templateData{
		Task:   t,
		NextID: t.ID + 1,
		PrevID: t.ID - 1,
	}

	if data.PrevID < 1 {
		data.PrevID = 1
	}

	MaxID, err := app.tasks.MaxID()
	if err != nil {
		app.serverError(w, err)
		return
	}

	if data.NextID > MaxID {
		data.NextID = MaxID
	}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}

	//fmt.Fprintf(w, "%v", t)

	fmt.Fprintf(w, "Tasks with ID %d showing...", id)
}

func (app *application) createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Creating a new task..."))
}
