package app

import (
	"fmt"
	"log"

	"github.com/mohammaderm/todoMicroService/authService/config"
	"github.com/mohammaderm/todoMicroService/authService/internal/repository"
	"github.com/mohammaderm/todoMicroService/authService/internal/usecase"
	"github.com/mohammaderm/todoMicroService/authService/pkg/jwt"
	"github.com/mohammaderm/todoMicroService/authService/pkg/logger"
)

func App() func() {

	// config init
	config, err := config.NewConfig("./config.yaml")
	if err != nil {
		log.Panic(err.Error())
	}

	// logger initi
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
	db, graceShutDown1, err := DBconnection(Logger, &config.Database)
	if err != nil {
		Logger.Info("can not connect to database.", map[string]interface{}{
			"err": err.Error(),
		})
	}
	Logger.Info("database is connected succesfully")

	// app logics
	jwtPkg := jwt.Newjwt(&config.Auth, Logger)
	userRepository := repository.New(Logger, db)
	authUsecase := usecase.New(userRepository, Logger, jwtPkg)

	Logger.Info(fmt.Sprintf("Grpc is runnig up on port %s", config.Grpc.Port))
	err = Server(config.Grpc, authUsecase, Logger)
	if err != nil {
		Logger.Panic("can not connect to server", map[string]interface{}{
			"err":  err.Error(),
			"port": config.Grpc.Port,
		})
	}

	return func() {
		graceShutDown1()
	}

}
