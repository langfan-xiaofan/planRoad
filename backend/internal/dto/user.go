package dto

type RegisterReq struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
