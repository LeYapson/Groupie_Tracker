package main

import (
	/*"fmt"
	"html/template"
	"net/http"
	"path"
	"time"*/
)

//? is used to launch the server in local
func main() {
	getArtistsApi()

	/*http.Handle("/", http.FileServer(http.Dir("./view/html")))

	fmt.Println("[INFO] - Starting the server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("[ERROR] - Server could not start properly.\n ", err)
	}
}

//* this function allow a URL path to a html file.
func indexHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := "index.html"
	if req.URL.Path == "/" {
		name = "index.html"
	} else {
		name = path.Base(req.URL.Path)
	}

	data := struct {
		Time time.Time
	}{
		Time: time.Now(),
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}*/
}
