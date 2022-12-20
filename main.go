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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getallmovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}
func getonemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, value := range movies {
		if value.ID == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}

	}
	json.NewEncoder(w).Encode("No movies available for given id.")

}
func deleteonemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	json.NewEncoder(w).Encode("Please provide valid id you want to delete")

}
func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break

		}
	}
	//movies=append(movies[:index],movies[index+1:]...)
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = params["id"]
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}
func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is Empty cannot create movie")
		return
	}
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(50000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to the page"))
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "12", Isbn: "1233", Title: "javascript", Director: &Director{Firstname: "ujjwal", Lastname: "silwal"}})
	movies = append(movies, Movie{ID: "13", Isbn: "1244", Title: "rust", Director: &Director{Firstname: "anish", Lastname: "manandhar"}})

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/movies", getallmovie).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteonemovie).Methods("DELETE")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", getonemovie).Methods("GET")
	//http.ListenAndServe(":8000", nil)
	fmt.Println("Listening to the port")
	log.Fatal(http.ListenAndServe(":2000", r))

}
