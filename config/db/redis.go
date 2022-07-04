package dbs

import (
	"fmt"
	"github.com/MKA-Nigeria/mkanmedia-go/config"
	"net/url"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

//ConnectRedis returns a connection to posgress instance
func ConnectRedis() *redis.Client {
	redisUrl := config.Env.REDIS_URL
	password := ""
	resolvedUrl := redisUrl
	if !strings.Contains(redisUrl, "localhost") {
		parsedURL, _ := url.Parse(redisUrl)
		password, _ = parsedURL.User.Password()
		resolvedUrl = parsedURL.Host
	}
	client := redis.NewClient(&redis.Options{
		Addr:        resolvedUrl,
		Password:    password,
		DialTimeout: time.Second * 20,
		DB:          0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}
