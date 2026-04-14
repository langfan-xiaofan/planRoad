package cmd

import (
	"backend/internal/cloudflare"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/middleware"
	"backend/internal/router"
	"backend/pkg/agent"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	agent.Init()
	if err != nil {
		panic(err)
	}
	err = config.Init()
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
	router.InitRouter(r, agent.Agent, agent.FileParser, db.MysqlDatabase, db.RedisDatabase, agent.ChatModel)
	err = cloudflare.CloudFlareInit()
	if err != nil {
		panic(err)
	}
	err = cloudflare.CloudFlareInit()
	if err != nil {
		panic(err)
	}
	r.Use(middleware.CrosMiddleware())
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
