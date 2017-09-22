package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rezandry/kejarmimpi/models"
)

//Collabs is for collabs
func Collabs(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var collabs models.Collabs
	var user models.User
	var post models.Post

	idPost := c.Params.ByName("id")
	token := c.Request.Header.Get("token")

	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	if err := db.Where("id = ?", idPost).First(&post).Error; err != nil {

	}

	collabs.IDPost = post.ID
	collabs.IDUser = user.ID
	if collabs.Check == false {
		collabs.Check = true
	} else {
		collabs.Check = false
	}
	db.Save(&collabs)
	c.JSON(200, collabs)

}
