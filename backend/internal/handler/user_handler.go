package handler

import (
	"backend/internal/dao"
	"backend/internal/dto"
	"backend/internal/service"
	"backend/pkg/res"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		res.Fail(c, 400, nil, "参数错误")
		return
	}
	userService := service.NewUserService()
	err := userService.Register(req)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	res.Success(c, nil)
	return
}

func LoginHandler(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Fail(c, 400, nil, "参数错误")
		return
	}
	userService := service.NewUserService()
	token, err := userService.Login(req)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	res.Success(c, token)
	return
}

func UploadAvatarHandler(c *gin.Context) {
	var req dto.AvatarReq
	var avatarRes dto.AvatarRes
	if err := c.ShouldBind(&req); err != nil {
		res.Fail(c, 400, nil, "参数错误")
		return
	}
	var err error
	userService := service.NewUserService()
	id := c.GetUint("userid")
	avatarRes.AvatarUrl, err = userService.UploadAvatar(req, id)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}

	res.Success(c, avatarRes.AvatarUrl)
	return
}

func GetUserPictureHandler(c *gin.Context) {
	id := c.GetUint("id")
	//Res, err := dao.GetUserPicture(id)
	Res, err := dao.GetUserPictureByUserId(id)
	if err != nil {
		res.Fail(c, 404, nil, err.Error())
		return
	}
	res.Success(c, Res)
	return
}
