package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"uid/db"
)

func RoutsListener() {
	router := mux.NewRouter()
	port := ":80"
	
	router.HandleFunc("/", func(resp http.ResponseWriter, rec *http.Request) {
		fmt.Fprintln(resp, "App is run")
	})

	router.HandleFunc("/migrations", func(resp http.ResponseWriter, rec *http.Request) {
		fmt.Fprintln(resp, db.RunMigrations(rec.URL.Query().Get("make")))
	})

	router.HandleFunc("/makeRequest", makeRequest).Methods("POST")




	log.Println("Server is listening port: ", port)
	go log.Fatalln(http.ListenAndServe(port, router))
}