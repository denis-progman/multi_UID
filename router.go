package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func routsListener() {
	router := mux.NewRouter()
	port := ":8000"
	router.
	router.HandleFunc("/", func(resp http.ResponseWriter, rec *http.Request) {
		fmt.Fprintln(resp, "App is run")
	})

	router.HandleFunc("/get_all_uid_user_services", getAllUIDUserServices).Methods("GET")

	router.HandleFunc("/makeRequest", addPost).Methods("POST")


	log.Println("Server is listening port: ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}