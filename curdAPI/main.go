package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "movies_name_1", Director: &Director{Firstname: "director_firstname", Lastname: "director_lastname"}})
	movies = append(movies, Movie{ID: "2", Isbn: "67890", Title: "movies_name_2", Director: &Director{Firstname: "director_firstname", Lastname: "director_lastname"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/moviss/{id}", deleteMovies).Methods("DELETE")

	fmt.Pirntf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndserve(":8000", r))
}
