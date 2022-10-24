package main

import (
	"fmt"
	"log"
	"net/http"
	"playlists-server/playlists"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to playlists-server")
}

func main() {
	godotenv.Load(".env")

	router := mux.NewRouter()
	
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/playlists", playlists.GetPlaylists)	

	//TODO:Setup CORS access origin
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}