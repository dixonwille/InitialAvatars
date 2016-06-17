package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dixonwille/InitialAvatars/cmd/AvatarServer/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	port string
)

func init() {
	godotenv.Load()
	port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/avatar/{initials}", handlers.Avatar).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
