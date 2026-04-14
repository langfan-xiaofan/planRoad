package model

type User struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Account  string `json:"account"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RemuseId string `json:"resume_id"`
}
