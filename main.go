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



func main() {
	fmt.Println("listening default port (:3002)")
	port := os.Getenv("PORT")
	var a int
	http.HandleFunc("/", IndexHandler)
	m := Handler {}
	m.Start("downloads")
	//http.HandleFunc("/download", DownloadHandler)
	http.HandleFunc("/getstart", GSHandler)
	http.HandleFunc("/market", ExtensionsHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	fmt.Println("Server Started!" + port)
	http.ListenAndServe(":" + port, nil)
	fmt.Println("Server Stopped!")
	fmt.Scan(&a)
}