package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
	THINGS TO NOTE:
	-there should not a space after json:
	for eg: it should be `json:"id"` not `json: "id"`

*/

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var movies []Movie


func getMovies(res http.ResponseWriter, req *http.Request){
	fmt.Println("Getting Movies")
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	fmt.Println("Getting Movie with id ", id)
	for _, movie := range movies{
		if id == movie.ID{
			json.NewEncoder(res).Encode(movie)
			return
		}
	}
}

func createMovie(res http.ResponseWriter, req *http.Request){
	fmt.Println("Creating new movie entry")
	res.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movies)

}

func updateMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	fmt.Println("Updating movie with id ", id)


	for index, item := range movies{
		if item.ID == id{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = id
			movies = append(movies, movie)
			json.NewEncoder(res).Encode(movie)
			return
		}
	}
}

func deleteMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	for index, movie := range movies{
		if id == movie.ID{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}


func main(){
	// main function

	movies = append(movies, Movie{
		ID: "001",
		Isbn: "4232123",
		Title: "Into The Wild",
		Director: &Director{
			ID: "001",
			FirstName: "Sean",
			LastName: "Penn",
		},
	})

	movies = append(movies, Movie{
		ID: "002",
		Isbn: "4232113",
		Title: "Inception",
		Director: &Director{
			ID: "002",
			FirstName: "Christopher",
			LastName: "Nolan",
		},
	})

	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	port := ":8000"
	fmt.Println("Starting Server @", port )
	log.Fatal(http.ListenAndServe(port, router))

}