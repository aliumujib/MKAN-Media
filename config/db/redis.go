package dbs

import (
	"fmt"
	"github.com/MKA-Nigeria/mkanmedia-go/config"
	"time"

	"github.com/go-redis/redis"
)

//ConnectRedis returns a connection to posgress instance
func ConnectRedis() *redis.Client {
	url := config.Env.REDIS_URL
	client := redis.NewClient(&redis.Options{
		Addr:        url,
		Password:    "",
		DialTimeout: time.Second * 20,
		DB:          0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}
