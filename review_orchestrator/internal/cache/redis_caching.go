package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectRedisCache() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer rdb.Close()

	status, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("Redis connection was refused")
		return
	}
	fmt.Println(status)
}
