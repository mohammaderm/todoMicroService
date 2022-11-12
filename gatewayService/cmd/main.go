package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mohammaderm/todoMicroService/gatewayService/cmd/app"
)

// @title    Todo API documentation
// @version  0.0.1
// @host     localhost:8080
// @BasePath /

//@securityDefinitions.apikey apiKey
//@in header
//@name Token

func main() {
	servers, shutDown := app.App()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		shutDown()
	}()
	servers.StartServer()
}
