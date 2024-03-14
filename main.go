package main

import (
	"log"
	"net/http"
	Handler "sangamCyber/Handler"
	Service "sangamCyber/Service"
	domain "sangamCyber/repository"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	// This is setting up a basic HTTP server in Go using the Gorilla Mux router and CORS
	// middleware. Here's a breakdown of what each part is doing:
	mux := mux.NewRouter()
	RepositoryDb := domain.NewRepositoryDb()

	H := Handler.Handlers{Service: Service.NewService(RepositoryDb)}

	mux.HandleFunc("/user/info", H.InsertUserInfo).Methods(http.MethodPost)
	mux.HandleFunc("/user/auth", H.AuthUserInfo).Methods(http.MethodPost)

	router := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Add PUT here
		AllowedHeaders: []string{"*"},
	}).Handler(mux)
	listenAddr := ":8080"

	log.Printf("About to listen on 8080. Go to https://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(listenAddr, router))

}
