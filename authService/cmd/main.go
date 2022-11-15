package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mohammaderm/todoMicroService/authService/cmd/app"
)

func main() {
	time.Sleep(time.Duration(20) * time.Second)
	shutDown := app.App()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGINT)

	go func() {
		<-c
		shutDown()
	}()
}
