package cmd

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	err = db.Neo4jInit()
	if err != nil {
		panic(err)
	}
	err = db.Neo4jInit()
	if err != nil {
		panic(err)
	}
	err = db.MysqlInit()
	if err != nil {
		panic(err)
	}
	err = db.RedisInit()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	router.InitRouter(r)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
