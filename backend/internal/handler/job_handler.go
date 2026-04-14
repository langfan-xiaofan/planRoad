package handler

import (
	"backend/internal/dao"
	"backend/internal/dto"
	"backend/pkg/res"

	"github.com/gin-gonic/gin"
)

func GetJobHandler(c *gin.Context) {
	Res, err := dao.GetJob()
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	res.Success(c, Res)
	return
}

func GetJobPictureHandler(c *gin.Context) {
	var req dto.GetPositionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	Res, err := dao.GetJobPicture(req.Position)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	res.Success(c, Res)
	return
}

func GetPositionSkillsHandler(c *gin.Context) {
	var req dto.GetPositionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	Res, err := dao.GetPositionSkills(req.Position)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	res.Success(c, Res)

}
func CreatUserPictureHandler(c *gin.Context) {
	var req dto.UserPictureReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	id := c.GetUint("id")
	req.UserId = id
	err = dao.CreatUserPicture(req)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	res.Success(c, nil)
	return
}
func GetUpPathHandler(c *gin.Context) {
	var req dto.GetPositionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	Res, err := dao.GetUpPath(req.Position)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	res.Success(c, Res)
	return
}

func GetJobRelationshipHandler(c *gin.Context) {
	var req dto.GetPositionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	Res, err := dao.GetJobRelationship(req)

	if err != nil {
		res.Fail(c, 400, nil, err.Error())
		return
	}
	res.Success(c, Res)
	return
}
