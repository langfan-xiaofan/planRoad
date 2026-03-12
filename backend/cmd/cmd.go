package cmd

import (
	"backend/internal/router"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
