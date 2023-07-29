package main

import (
	"encoding/json"
	"io"
	"log"
	"uid/models"

	"github.com/prevoty/pdo"
)



func createConnection() *pdo.MySQL {
	connection, err := pdo.NewMySQL("gouser:gopass@tcp(mysql-go-db:3306)/godb")
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

func getUIDByUserID(userID int, connection *pdo.MySQL) {
	uid := new(models.Uid)
	connection.Find(uid, "WHERE `id` = ?", userID)
}

func makeRequest(request_body io.Reader, connection *pdo.MySQL) (int64, string) {
	request := new(models.Request)
	err := json.NewDecoder(request_body).Decode(&request)

	if err != nil {
		return 0, `{"error": "Error unmarshaling the request"}`
	}

	id, err := connection.Create(&request)
	if err != nil {
		log.Fatal(err)
	}
	return id, ""
}
