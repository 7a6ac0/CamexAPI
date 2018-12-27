package main

import (
	"CamexAPI/auth"
	"CamexAPI/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/get_token", controllers.GetToken).Methods("GET").Queries("imei", "{imei}")
	router.HandleFunc("/api/data/new", controllers.CreateData).Methods("POST")
	router.HandleFunc("/api/datas", controllers.GetDatas).Methods("GET") //  user/2/contacts

	router.Use(auth.JwtAuthentication) //attach JWT auth middleware

	router.NotFoundHandler = auth.NotFoundHandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":" + port, router) //Launch the auth, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
