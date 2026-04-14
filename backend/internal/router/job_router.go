package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initJobRouter(g *gin.Engine) {
	jobGroup := g.Group("/job")
	jobGroup.Use(middleware.JwtMiddleware())
	{
		jobGroup.GET("/getjob", handler.GetJobHandler)
		jobGroup.POST("/getjobpicture", handler.GetJobPictureHandler)
		jobGroup.POST("/getpositionskills", handler.GetPositionSkillsHandler)
		//jobGroup.POST("/creatuserpicture", handler.CreatUserPictureHandler)
	}
	jobGroup.POST("/creatuserpicture", handler.CreatUserPictureHandler)
}
