package main

import (
	"encoding/json"
	"log"
	"net/http"
	"uid/models"
	"uid/repository"
)

var connection = repository.CreateConnection()


func setHeaders(resp http.ResponseWriter) {
	resp.Header().Set("Content-type", "application/json")
}

// func getAllUIDUserServices(resp http.ResponseWriter, req *http.Request) {
// 	setHeaders(resp)
// }

func makeRequest(resp http.ResponseWriter, req *http.Request) {
	setHeaders(resp)
	requestAPIRequest := new(models.RequestAPIRequest)
	err := json.NewDecoder(req.Body).Decode(&requestAPIRequest)

	if err != nil {
		log.Fatal(err)
	} 

	request, err := repository.MakeRequest(requestAPIRequest, connection)

	if err != nil {
		log.Fatal(err)
	} 
	
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(request)

	if err != nil {
		log.Fatal(err)
	} 

	resp.Write(result) 
 }