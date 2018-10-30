package main

import (
	"cashgone/auth"
	"cashgone/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contact/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/contacts", controllers.GetContacts).Methods("GET") //  user/2/contacts
	router.HandleFunc("/api/contact/{id:[0-9]+}", controllers.GetContact).Methods("GET")

	githubRouter := router.PathPrefix("/github").Subrouter()
	githubRouter.HandleFunc("/search/repos", controllers.SearchRepos).Methods("GET").Queries("q", "{query}")

	router.Use(auth.JwtAuthentication) //attach JWT auth middleware

	router.NotFoundHandler = auth.NotFoundHandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" //localhost
	}

	err := http.ListenAndServe(":" + port, router) //Launch the auth, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
