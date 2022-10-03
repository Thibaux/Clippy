package Redis

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis () *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:	  os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:		  0,
	})
	
	return rdb
}
