package main

import (
	"fmt"
	"net/http"
)

//? is used to launch the server in local
func main() {

	http.Handle("/", http.FileServer(http.Dir("./view/html")))

	fmt.Println("[INFO] - Starting the server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("[ERROR] - Server could not start properly.\n ", err)
	}
}
