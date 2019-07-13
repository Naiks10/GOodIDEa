package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html", "./templates/footer.html", "./templates/header.html")
	if err != nil {
		fmt.Println("IndexHandler status : Error => ", err)
	} else {
		fmt.Println("IndexHandler status : OK")
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	fmt.Println("listening default port")

	port := os.Getenv("PORT")
	var a int
	defer func() {
		fmt.Println("Server Stopped!")
		fmt.Scan(&a)
	}()
	http.HandleFunc("/", IndexHandler)
	m := Handler{}
	m.Start("downloads")
	m.Start("getstart")
	m.Start("marketplace")
	m.Start("devs")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	//fmt.Println("Server Started!" + port)
	http.ListenAndServe(":"+port, nil)
}
