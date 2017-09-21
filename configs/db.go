package configs

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//For connect postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//InitDb is for connecting to database
func InitDb() *gorm.DB {
	var err error
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=kejarmimpi sslmode=disable password=postgresreza")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Success!")
	}
	return db
}
