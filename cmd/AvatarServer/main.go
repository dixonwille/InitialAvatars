package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dixonwille/InitialAvatars/cmd/AvatarServer/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	osPort string
	port   string
)

func init() {
	godotenv.Load()
	osPort = os.Getenv("AVATAR_PORT")
	if osPort == "" {
		osPort = "80"
	}
	flag.StringVar(&port, "l", osPort, "Port to listen on")
}

func main() {
	flag.Parse()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/avatar/{initials}", handlers.Avatar).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
