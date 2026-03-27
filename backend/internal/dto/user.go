package dto

import "mime/multipart"

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

type AvatarReq struct {
	Avatar *multipart.FileHeader `form:"avatar"`
}
type AvatarRes struct {
	AvatarUrl string `json:"avatarurl"`
}
