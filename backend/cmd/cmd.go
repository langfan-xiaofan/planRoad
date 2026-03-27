package cmd

import (
	"backend/internal/cloudflare"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/router"
	"backend/pkg/agent"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	fileParser := agent.OpenAI()
	agent := agent.NewAgent()
	err := godotenv.Load()
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
	router.InitRouter(r, agent, fileParser, db.MysqlDatabase, db.RedisDatabase)
	err = cloudflare.CloudFlareInit()
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
=======
	router.InitRouter(r, agent)
>>>>>>> 817bb7a2540f5637395e2a2e1b06aeef325acf68
	err = cloudflare.CloudFlareInit()
	if err != nil {
		panic(err)
	}
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
