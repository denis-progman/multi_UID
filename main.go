package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations() {
	db, err := sql.Open("mysql", "gouser:gopass@tcp(mysql-go-db:3306)/godb?multiStatements=true")
	if err != nil {
		log.Fatal("Can not connect: ", err)
	}

    driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Can not make instance: ", err)
	}
    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "godb", 
        driver,
    )
	if err != nil {
		log.Fatal("Can not make migrations: ", err)
	}
    m.Up()
    // m.Down()
}

func main() {
	runMigrations()
	router := mux.NewRouter()
	port := ":80"
	router.HandleFunc("/", func(resp http.ResponseWriter, rec *http.Request) {
		fmt.Fprintln(resp, "App is run")
	})

	log.Println("Server is listening port: ", port)
	log.Fatalln(http.ListenAndServe(port, router))
} 
