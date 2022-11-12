package app

import (
	"fmt"
	"log"

	"github.com/mohammaderm/todoMicroService/todoService/config"
	"github.com/mohammaderm/todoMicroService/todoService/internal/repository"
	"github.com/mohammaderm/todoMicroService/todoService/internal/usecase"
	"github.com/mohammaderm/todoMicroService/todoService/pkg/logger"
)

func App() func() {
	// config init
	config, err := config.NewConfig("./config.yaml")
	if err != nil {
		log.Panic(err.Error())
	}

	// logger init
	Logger, err := logger.New(&logger.Logconfig{
		Path:         config.Logger.Internal_Path,
		Pattern:      config.Logger.Filename_Pattern,
		MaxAge:       config.Logger.Max_Age,
		RotationTime: config.Logger.Rotation_Time,
		RotationSize: config.Logger.Max_Size,
		Mode:         config.Logger.Mode,
	})
	if err != nil {
		log.Panic(err.Error())
	}

	// db connection
	db, graceShutDown1, err := DBconnection(Logger, config.Database)
	if err != nil {
		Logger.Panic("can not connect to db", map[string]interface{}{
			"err": err.Error(),
		})
	}
	Logger.Info("db connected succesfully")

	// redis connection
	redis, err := ConnectRedis(&config.Redis, Logger)
	if err != nil {
		Logger.Panic("can not connect to redis(cache)", map[string]interface{}{
			"err": err.Error(),
		})
	}
	Logger.Info("redis connected succesfully")
	// --------------------------------------------

	repo := repository.New(db)
	cache := usecase.NewCache(redis)
	useCase := usecase.New(repo, cache)

	// grpc server
	Logger.Info(fmt.Sprintf("Grpc is runnig up on port %s", config.Grpc.Port))
	err = Server(config.Grpc, useCase)
	if err != nil {
		Logger.Panic("can not connect to server", map[string]interface{}{
			"err":  err.Error(),
			"port": config.Grpc.Port,
		})
	}
	return graceShutDown1
}
