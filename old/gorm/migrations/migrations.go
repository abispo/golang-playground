package migrations

import (
	"golang-playground/gorm/database"
	"golang-playground/gorm/models"
)

func Migrate() {
	database.DBCon.AutoMigrate(models.Product{})

	//database.DBCon.Model(&models.Product{}).DropColumn("info")
}
