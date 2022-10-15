package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"peli1", 1234, "Desconocido"},
	Movie{"peli1", 1234, "Desconocido"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "holamundo con GOLANG con router")
}
func MovieList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)

}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]
	fmt.Fprintf(w, "Se va a mostrar la pelicula numero %s", movieId)
}
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movieData Movie
	err := decoder.Decode(&movieData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	log.Println(movieData)
	movies = append(movies, movieData)
}
