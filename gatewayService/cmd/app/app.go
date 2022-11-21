package app

import (
	"fmt"

	"log"
	"net/http"
	"sync"

	"github.com/mohammaderm/todoMicroService/gatewayService/config"
	"github.com/mohammaderm/todoMicroService/gatewayService/internal/delivery"
	"github.com/mohammaderm/todoMicroService/gatewayService/pkg/logger"
	"github.com/mohammaderm/todoMicroService/gatewayService/pkg/monitoring"
)

func App() (*serverDep, func()) {

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
	})
	if err != nil {
		log.Panic(err.Error())
	}

	// -----------------------------

	// app handlers
	authHandler := delivery.NewAuthHandler(Logger, &config.Service)
	todoHandler := delivery.NewTodoHandler(Logger, &config.Service)
	categoryHandler := delivery.NewCategoryHandler(Logger, &config.Service)

	// servers init
	metrics := monitoring.New(config.Metrics.Port, Logger, config.Server)
	router := delivery.RouterProvider(&delivery.RouteProvider{
		AuthHandler:     authHandler,
		TodoHandler:     todoHandler,
		CategoryHandler: categoryHandler,
		Monitoring:      metrics,
		Cfg:             &config.Token,
	})

	server, graceShutDown := delivery.ServerProvider(Logger, config.Server, router)
	serverDep := serverDepProvider(Logger, server, metrics)

	return serverDep, func() {
		metrics.Shotdown()
		graceShutDown()
	}
}

type serverDep struct {
	server  *http.Server
	logger  logger.Logger
	metrics monitoring.MetricsCallectors
}

func (s serverDep) StartServer() {

	s.logger.Info(fmt.Sprintf("server is runnig in %s ...", s.server.Addr))
	s.logger.Info("metrics server is runnig ...")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := s.server.ListenAndServe()
		if err != nil {
			s.logger.Warning("can not runnig server", map[string]interface{}{
				"err":  err.Error(),
				"Addr": s.server.Addr,
			})
		}
	}()

	go func() {
		defer wg.Done()
		err := s.metrics.Start()
		if err != nil {
			s.logger.Warning("can not runnig metrics server", map[string]interface{}{
				"err": err.Error(),
			})
		}
	}()
	wg.Wait()
}

func serverDepProvider(logger logger.Logger, httpServer *http.Server, metricCollector monitoring.MetricsCallectors) *serverDep {
	return &serverDep{
		server:  httpServer,
		logger:  logger,
		metrics: metricCollector,
	}
}
