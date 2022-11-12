package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mohammaderm/authService/cmd/app"
)

func main() {
	shutDown := app.App()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGINT)

	go func() {
		<-c
		shutDown()
	}()
}
