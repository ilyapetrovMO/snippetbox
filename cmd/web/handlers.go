package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ilyapetrovMO/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	l, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}

	for _, s := range l {
		fmt.Fprintf(w, "%v\n", s)
	}

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
	}
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "<h1>%s | %d</h1><p>%s</p><div>Created: %s</div><div>Expires: %s</div>",
		s.Title, s.ID, s.Content, s.Created, s.Expires)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "My Dummy data"
	content := "This is dummy content"
	expires := "1"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
