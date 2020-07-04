package main

import (
	"log"
	"logindemo/config"
	"logindemo/controler"
	"logindemo/db"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	db.Connect()

	handleRequests()

}

func handleRequests() {

	r := mux.NewRouter()
	r.HandleFunc("/api/register", controler.SignUpUser).Methods("POST")
	r.HandleFunc("/api/login", controler.LoginUser).Methods("POST")
	r.HandleFunc("/api/users", controler.AllUsers).Methods("GET")

	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)

	log.Println("Starting server...")
	config, err := config.GetConfig()
	if err != nil {
		println(err.Error())
	} else {
		log.Panic(http.ListenAndServe(config.Production.PORT, serveMux))
	}

}
