package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Handler struct {
	Template string
}

func (m Handler) Start(template string) {
	m.Template = template
	http.HandleFunc("/"+m.Template, m.Execute);
}

func (m Handler) Execute(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/"+m.Template+".html" , "./templates/footer.html", "./templates/header.html")
	if err != nil {
		fmt.Println(err)
	}
	if r.URL.Path != "/" + m.Template {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	t.ExecuteTemplate(w, m.Template, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        t, err := template.ParseFiles("./templates/404.html",  "./templates/header.html", "./templates/footer.html"))
		if err != nil {}
		t.ExecuteTemplate(w, "404", nil)
    }
}