package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	//"math/rand"
	"net/http"
	//"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lsatname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "movies_name_1", Director: &Director{Firstname: "director_firstname", Lastname: "director_lastname"}})
	movies = append(movies, Movie{ID: "2", Isbn: "67890", Title: "movies_name_2", Director: &Director{Firstname: "director_firstname", Lastname: "director_lastname"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	//r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	//r.HandleFunc("/movies", createMovies).Methods("POST")
	//r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
