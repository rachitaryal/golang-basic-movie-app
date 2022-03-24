package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title`
	Director *Director `json: "director"`
}

type Director struct {
	ID string `json: "id"`
	FirstName string `json: "firstName"`
	LastName string `json: "lastName"`
}

func getMovies(res http.ResponseWriter, req *http.Request){}

func getMovie(res http.ResponseWriter, req *http.Request){}

func createMovie(res http.ResponseWriter, req *http.Request){}

func updateMovie(res http.ResponseWriter, req *http.Request){}

func deleteMovie(res http.ResponseWriter, req *http.Request){}

var movies []Movie

func main(){
	// main function

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