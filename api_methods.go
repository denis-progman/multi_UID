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
	requestAPIRequest := new(models.RequestAPIRequest)
	log.Println(json.NewDecoder(req.Body).Decode(&requestAPIRequest))

	if request, err := repository.MakeRequest(requestAPIRequest, connection); err == nil {
		setHeaders(resp)
		resp.WriteHeader(http.StatusOK)
		if result, err := json.Marshal(request); err == nil {
			resp.Write(result) 
		}
		log.Println(err)
	} 
	// log.Println(err)
 }