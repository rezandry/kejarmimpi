package models

//Post is struct for save data post
type Post struct {
	ID         uint   `form:"id" json:"id"`
	Content    string `form:"content" json:"content"`
	Date       string `form:"date" json:"date"`
	Photo      string `form:"photo" json:"photo"`
	IDUser     uint   `form:"idUser" json:"idUser"`
	User       User   `form:"User" json:"User"`
	IDCategory uint   `form:"idCategory" json:"idCategory"`
}
