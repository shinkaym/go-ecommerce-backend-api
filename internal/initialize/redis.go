package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/shinkaym/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error: ", zap.Error(err))
	}

	fmt.Println("InitRedis is running")
	global.Rdb = rdb
	redisExample()
}

// redisExample is a sample function that sets and retrieves a key-value pair from Redis
func redisExample() {
	// Set the key "score" with the value 100 in Redis
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Error redis setting:", zap.Error(err)) // Log error if setting the value fails
	}

	// Retrieve the value of the key "score" from Redis
	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error redis setting: ", zap.Error(err)) // Log error if retrieving the value fails
		return
	}
	// Log the retrieved value of "score" using the logger
	global.Logger.Info("value score is::", zap.String("score", value))
}
