package config

import (
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	Jwt        `mapstructure:"jwt"`
	Neo4j      `mapstructure:"neo4j"`
	Redis      `mapstructure:"redis"`
	MongoDb    `mapstructure:"mongodb"`
	Mysql      `mapstructure:"mysql"`
	Cloudflare `mapstructure:"Cloudflare"`
}
type Jwt struct {
	Key  string `mapstructure:"key"`
	Hour int    `mapstructure:"hour"`
}

type Neo4j struct {
	Url      string `mapstructure:"url"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Redis struct {
	Url      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type MongoDb struct {
	Url      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
}

type Mysql struct {
	Dsn         string `mapstructure:"dsn"`
	MaxOpenConn int    `mapstructure:"MaxOpenConn"`
	MaxIdleConn int    `mapstructure:"MaxIdleConn"`
}

type Cloudflare struct {
	AccountId  string `mapstructure:"AccountId"`
	ApiKey     string `mapstructure:"ApiKey"`
	ApiSecret  string `mapstructure:"ApiSecret"`
	BucketName string `mapstructure:"BucketName"`
	PublicKey  string `mapstructure:"PublicKey"`
}

func Init() {
	viper.SetConfigFile("configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(err)
	}

}
