package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Index struct {
	Artists      []Artist
	Locations    []Location
	ConcertDates []ConcertDate
	Relations    []Relation
}

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Location struct {
	Id           int
	Locations    []string
	ConcertDates []ConcertDate
}

type ConcertDate struct {
	Id           int
	ConcertDates []string
}

type Relation struct {
	Id             int
	Dateslocations [][]string
}

//? this function go get the artists API with is URL and drop it into the terminal
func getArtistsApi() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	//*We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//*convert the body into the the artist struct
	var index Index
	json.Unmarshal(body, &index)

	for i, p := range index.Artists {
		fmt.Println("Artists",(i+1), ":", p.Name)
      fmt.Println("image:", p.Image)
		fmt.Println("-------------------------")
      fmt.Println("Members:", p.Members)
      fmt.Println("Creation Date:", p.CreationDate)
      fmt.Println("First Album:", p.FirstAlbum)
      fmt.Println("Locations:", p.Locations)
      fmt.Println("Concert Dates:", p.ConcertDates)
      fmt.Println("Relations:", p.Relations)
      fmt.Println()
	}
}
