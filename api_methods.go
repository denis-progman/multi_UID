package main

import (
	"encoding/json"
	"net/http"
)

func setHeaders(resp http.ResponseWriter) {
	resp.Header().Set("Content-type", "application/json")
}

func getAllUIDUserServices(resp http.ResponseWriter, req *http.Request) {
	setHeaders(resp http.ResponseWriter)
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshaling the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)  
}

func makeRequest(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshaling the request"}`))
		return
	} 
	post.Id = len(posts) + 1
	posts = append(posts, post)
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(posts)
	resp.Write(result) 
 }