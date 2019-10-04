package main

import (
	"golang-playground/gorm/database"
	"golang-playground/gorm/migrations"
)

func main() {
	database.InitDB()

	migrations.Migrate()

	defer database.DBCon.Close()
}
