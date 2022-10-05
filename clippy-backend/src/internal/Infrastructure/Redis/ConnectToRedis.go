package Redis

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

func ConnectToRedis () *rejson.Handler {
	rh := rejson.NewReJSONHandler()
	
	cli := redis.NewClient(&redis.Options{
		Addr:	  os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:		  0,
	})
	
	rh.SetGoRedisClient(cli)

	return rh
}
