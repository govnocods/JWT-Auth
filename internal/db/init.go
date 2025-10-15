package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type DataBase struct {
	db *sql.DB
}

func (d *DataBase) Connect() *DataBase {
	var err error
	d.db, err = sql.Open("sqlite", "file:database.db?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatal(err)
	}

	if err = d.db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successful connection to DataBase")
	}

	return &DataBase{db: d.db}
}
