package main

import (
	"log"
	"net/http"

	"github.com/govnocods/JWT-Authorization/app"
	"github.com/govnocods/JWT-Authorization/internal/db"
)

func main() {

	database := &db.DataBase{}
	database.Connect()

	app := app.NewApp(database)

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
