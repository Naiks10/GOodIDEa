package main

import (
	"net/http"
	"html/template"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/downloads.html" , "./templates/footer.html", "./templates/header.html")
	if err != nil {}
	if r.URL.Path != "/download" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	t.ExecuteTemplate(w, "download", nil)
}

func GSHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/getstart.html" , "./templates/footer.html", "./templates/header.html")
	if err != nil {}
	if r.URL.Path != "/getstart" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	t.ExecuteTemplate(w, "getstart", nil)
}

func ExtensionsHandler (w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/marketplace.html" , "./templates/footer.html", "./templates/header.html")
	if err != nil {}
	if r.URL.Path != "/market" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	t.ExecuteTemplate(w, "market", nil)
}

func DevsHandler (w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/devs.html" , "./templates/footer.html", "./templates/header.html")
	if err != nil {}
	if r.URL.Path != "/devs" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	t.ExecuteTemplate(w, "devs", nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        t, err := template.ParseFiles("./templates/404.html",  "./templates/header.html")
		if err != nil {}
		t.ExecuteTemplate(w, "404", nil)
    }
}