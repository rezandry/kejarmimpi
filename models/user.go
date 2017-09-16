package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID              uint   `json:"user_id"`
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Photo           string `json:"profile_photo"`
	Work            string `json:"work"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}