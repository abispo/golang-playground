package main

import (
	"golang-playground/gorm/database"
	"golang-playground/gorm/migrations"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	database.InitDB()

	migrations.Migrate()

	defer database.DBCon.Close()
}
