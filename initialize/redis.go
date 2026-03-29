package initialize

import (
	"Diggpher/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func ConnRedis() {
	global.Log.Info("Connecting to Redis")
	client := redis.NewClient(&redis.Options{
		Addr:     global.CONFIG.Redis.Addr,
		Password: global.CONFIG.Redis.Password,
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		global.Log.Fatal("Failed to connect to Redis", zap.Error(err))
	}
	global.Redis = client
	global.Log.Info("Redis connected successfully")
}
