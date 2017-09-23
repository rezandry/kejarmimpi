package models

//Star is used for artikel
type Star struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Photo   string `json:"photo"`
	Job     string `json:"job"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
