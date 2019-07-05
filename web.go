package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html" , "./templates/footer.html", "./templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}



func main() {
	fmt.Println("listening default port (:3002)")
	var a int
	http.HandleFunc("/", IndexHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	fmt.Println("Server Started!")
	http.ListenAndServe(":3002", nil)
	fmt.Println("Server Stopped!")
	fmt.Scan(&a)
}