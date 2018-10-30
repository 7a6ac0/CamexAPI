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
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/contacts", controllers.GetContacts).Methods("GET") //  user/2/contacts

	router.Use(auth.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = auth.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the auth, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
