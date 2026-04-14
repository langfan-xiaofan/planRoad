package cmd

import (
	"backend/internal/cloudflare"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/middleware"
	"backend/internal/router"
	"backend/pkg/agent"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	agent.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	config.Init()
	db.Neo4jInit()
	db.MysqlInit()
	db.RedisInit()
	db.MongoDbInit()
	r := gin.Default()
	router.InitRouter(r, agent.Agent, agent.FileParser, db.MysqlDatabase, db.RedisDatabase, agent.ChatModel)
	r.Use(middleware.CrosMiddleware())
	cloudflare.CloudFlareInit()
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
