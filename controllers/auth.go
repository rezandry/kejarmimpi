package controllers

import (
	"ervinismu/BE-kejarmimpi/code/models"
	"regexp"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Struct for error response
type Response struct {
	Code    string `json:"error_code"`
	Message string `json:"message"`
}

func Register(c *gin.Context) {

	response := Response{}
	// db := InitDb()
	// defer db.Close()
	user := models.User{}
	c.BindJSON(&user)
	// Check
	if user.Email == string("") {
		response.Code = "401"
		response.Message = "Email can not be empty"
		c.JSON(400, response)
		return
	}
	if user.Username == string("") {
		response.Code = "401"
		response.Message = "Username can not be empty"
		c.JSON(400, response)
		return
	}
	if user.Password+user.PasswordConfirm == string("") {
		response.Code = "401"
		response.Message = "Password can not be empty"
		c.JSON(400, response)
		return
	}
	// Check Valid email
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegexp.MatchString(user.Email) {
		response.Code = "401"
		response.Message = "invalid format email!"
		c.JSON(400, response)
		return
	}
	// Check if password not same
	if user.Password != user.PasswordConfirm {
		response.Code = "401"
		response.Message = "Password not same!"
		c.JSON(400, response)
		return
	}
	// Check email where email exist
	if err := db.Where("email= ?", user.Email).First(&user).Error; err != nil {
		password := []byte(user.Password + user.PasswordConfirm)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			response.Code = "401"
			response.Message = "Failed to encrypt password!"
			c.JSON(400, response)
		}
		user.Password = string(hashedPassword)
		user.PasswordConfirm = string("")
		// Insert into database after encrypt
		db.Create(&user)
		// Response
		response.Code = "200"
		response.Message = "Register Succes!"
		c.JSON(400, response)
	} else {
		response.Code = "401"
		response.Message = "Email has been used!"
		c.JSON(400, response)
		return
	}
}
