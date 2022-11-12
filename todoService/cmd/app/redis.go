package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/mohammaderm/todoMicroService/todoService/config"
	"github.com/mohammaderm/todoMicroService/todoService/pkg/logger"
)

func ConnectRedis(config *config.Redis, logger logger.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Server + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, err
	}
	logger.Info(pong)

	return client, nil
}
