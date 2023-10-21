package main

import (
	"fmt"
	"log"
	"nativeApi/db"
	"nativeApi/handler"
	"nativeApi/models"
	"net/http"
	"os"
)

func main() {

	log.Print("The is Server Running on localhost port 3000")

	//initialize the database
	db.Moviedb["001"] = models.Movie{ID: "001", Title: "A Space Odssey", Description: "Science fiction"}
	db.Moviedb["002"] = models.Movie{ID: "002", Title: "Citizen Kane", Description: "Drama"}
	db.Moviedb["003"] = models.Movie{ID: "003", Title: "Raiders of the Lost Ark ", Description: "Action and adventure"}
	db.Moviedb["004"] = models.Movie{ID: "004", Title: "The General", Description: "Comedy"}

	// route goes here

	// test route
	http.HandleFunc("/", handler.TestHandler)

	//get movies
	http.HandleFunc("movies", handler.GetMovies)

	//get single movie
	http.HandleFunc("/movie", handler.GetMovie)

	//Add movie
	http.HandleFunc("/movie/add", handler.AddMovie)

	// delete movie
	http.HandleFunc("/movie/delete", handler.DeleteMovie)

	err := http.ListenAndServe(":3000", nil)
	//print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
