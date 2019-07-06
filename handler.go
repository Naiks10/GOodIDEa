package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Handler struct {
	Template string
}

func (m *Handler) SetTemplate(template string) string {
	m.Template = template
	return m.Template
}

func (m Handler) Start(template string) {
	http.HandleFunc("/"+m.SetTemplate(template), m.Execute)
}

func (m Handler) Execute(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/"+m.Template+".html", "./templates/footer.html", "./templates/header.html")
	if err != nil {
		fmt.Println(err)
	}
	if r.URL.Path != "/"+m.Template {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	t.ExecuteTemplate(w, m.Template, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("./templates/404.html", "./templates/header.html", "./templates/footer.html")
		if err != nil {
			fmt.Println(err)
		}
		t.ExecuteTemplate(w, "404", nil)
	}
}
