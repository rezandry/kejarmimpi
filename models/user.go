package models

//User is struct for database User
type User struct {
	ID              uint   `form:"id" json:"id"`
	Name            string `form:"name" json:"name"`
	Email           string `form:"email" json:"email"`
	Password        string `form:"password" json:"password"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm"`
	Photo           string `form:"photo" json:"photo"`
	Job             string `form:"job" json:"job"`
	Location        string `form:"location" json:"location"`
	Headline        string `form:"headlin" json:"headline"`
	Bio             string `form:"bio" json:"bio"`
	Biography       string `form:"biography" json:"biography"`
	Education       string `form:"education" json:"education"`
	Career          string `form:"career" json:"career"`
	OtherCareer     string `form:"othercareer" json:"othercareer"`
	Organization    string `form:"organization" json:"organization"`
	Facebook        string `form:"facebook" json:"facebook"`
	Linkedin        string `form:"linkedin" json:"linkedin"`
	Twitter         string `form:"twitter" json:"twitter"`
	Instagram       string `form:"instagram" json:"instagram"`
	Token           string `form:"token" json:"token"`
}
