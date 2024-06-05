package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for index , item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for index , item := range movies{
		if item.ID	== params["id"]{
		movies = append(movies[:index],movies[index+1:]...)
		var movie Movie
		_= json.NewDecoder(r.Body).Decode(&movie)
		movie.ID = params["id"]
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movie)
		return
		}
	}
}

func main(){
	r := mux.NewRouter()
	movies = append(movies, Movie{ID:"1",Isbn: "438777",Title: "ANIME",Director: &Director{Firstname: "Hentai",Lastname: "DOE"}})
	movies = append(movies, Movie{ID: "2",Isbn: "45444",Title: "Biopic",Director: &Director{Firstname: "Steve",Lastname: "Jordan"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at PORT 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}