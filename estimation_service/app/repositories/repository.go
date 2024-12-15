package repositories

import (
	"estimation_service/config"
	"github.com/redis/go-redis/v9"
)

var (
	Connection *redis.Client
)

func InitDBConnection() {
	ConnectionInit()
}

func connectionToRedis(Host, dbPassword string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

func ConnectionInit() {
	dbHost := config.GetConfig().Database.Redis.Host
	dbPassword := config.GetConfig().Database.Redis.Password

	db := connectionToRedis(dbHost, dbPassword)
	Connection = db
}
