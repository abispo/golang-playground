package database

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DBCon is the gorm.DB reference
var DBCon *gorm.DB

// InitDB called on the main function
func InitDB() {
	var err error

	DBCon, err = gorm.Open("mysql", viper.Get("CONNECTION_STRING"))

	if err != nil {
		panic(err)
	}
}
