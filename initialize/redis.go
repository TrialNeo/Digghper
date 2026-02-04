package initialize

import (
	"Diggpher/global"
	"context"
	"github.com/go-redis/redis/v8"
)

func ConnRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     global.CONFIG.Redis.Addr,
		Password: global.CONFIG.Redis.Password,
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	global.Redis = client
}
