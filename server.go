package main

import (
	"github.com/jinzhu/gorm"
	"ervinismu/km-gingonic/controllers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gin-gonic/gin"
)

func main() {

	var err error
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=postgres sslmode=disable password=newpassword")
	defer db.Close()
	db.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal("database error")
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controllers.Register)
	}
	r.Run()
}