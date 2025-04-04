package main

import (
	"database/sql"

	"github.com/goIdioms/store/database"
)

func initDatabase() (*sql.DB, error) {
	return database.Connect()
}
