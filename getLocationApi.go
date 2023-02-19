package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Location struct {
Id int
Locations []string
ConcertDates string
}
//! ON A UN PB CHEF PAR ICI
//? this function reads the location API from a local file and prints the relevant information
func handleLocation(w http.ResponseWriter, r *http.Request) {
	locations, err := getLocation("./data/location.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//* Parse the HTML template
	tmpl, err := template.ParseFiles("./view/html/location.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//* Executes the template by passing the location data
	err = tmpl.Execute(w, locations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//* Get the location from the API
func getLocation(filePath string) ([]Location, error) {
	//* Opening the local JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier : %w", err)
	}
	defer file.Close()

	//* Reading the contents of the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier : %w", err)
	}

	//* Decoding JSON content in an Location slice
	var locations []Location
	err = json.Unmarshal(content, &locations)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du décodage du JSON : %w", err)
	}

	return locations, nil
}