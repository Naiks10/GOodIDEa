package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html" , "./templates/footer.html", "./templates/header.html")
	if err != nil {}
	if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	t.ExecuteTemplate(w, "index", nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        t, err := template.ParseFiles("./templates/404.html",  "./templates/header.html")
		if err != nil {}
		t.ExecuteTemplate(w, "404", nil)
    }
}

func main() {
	fmt.Println("listening default port (:3002)")
	port := os.Getenv("PORT")
	var a int
	http.HandleFunc("/", IndexHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	fmt.Println("Server Started!")
	http.ListenAndServe(":" + port, nil)
	fmt.Println("Server Stopped!")
	fmt.Scan(&a)
}