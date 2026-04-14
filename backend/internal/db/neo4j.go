package db

import (
	"backend/internal/config"
	"context"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

var Neo4j = config.Conf.Neo4j
var Neo4jDb neo4j.Driver

func Neo4jInit() {

	Neo4j.Url = config.Conf.Neo4j.Url
	Neo4j.User = config.Conf.Neo4j.User
	Neo4j.Password = config.Conf.Neo4j.Password

	driver, err := neo4j.NewDriver(Neo4j.Url, neo4j.BasicAuth(Neo4j.User, Neo4j.Password, ""))
	Neo4jDb = driver
	if err != nil {
		panic(err)
	}
	err = driver.VerifyConnectivity(context.Background())
	if err != nil {
		panic(err)
	}
}
