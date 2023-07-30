package db

import (
	"database/sql"
	"log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(action string) string {
	db, err := sql.Open("mysql", "gouser:gopass@tcp(mysql-go-db:3306)/godb?multiStatements=true")
	if err != nil {
		log.Fatal("Can not connect: ", err)
	}

    driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Can not make instance: ", err)
	}
    m, err := migrate.NewWithDatabaseInstance(
        "file://db/migrations",
        "godb", 
        driver,
    )
	if err != nil {
		log.Fatal("Can not make migrations: ", err)
	}

	switch action {
		case "up":
			err = m.Up()
		case "down": 
			err = m.Down()
		default:
			return "Migrations: undefined command"
	}

	if err != nil {
		log.Fatal("Can not make migrations: ", err)
	}
	return "Migrations is :" +  action
}