package controllers

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/rezandry/kejarmimpi/models"
	//For connect postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//InitDb is for connecting to database
func InitDb() *gorm.DB {
	var err error
	dbhost := os.Getenv("DATABASE_URL")
	if dbhost == "" {
		dbhost = "host=127.0.0.1 user=postgres dbname=kejarmimpi sslmode=disable password=postgresreza"
	}
	db, err := gorm.Open("postgres", dbhost)
	db.AutoMigrate(&models.Post, &models.User)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Success!")
	}
	return db
}
