package app

import (
	"context"
	"log"
	"main/internal/config"
	"main/internal/httpserver"
	"os"
	"os/signal"
	"time"
)

// Init - like init()
func Init() {
	log.Print("load configs...")
	config.Init()
	log.Print("load configs done")
}

// Main - like main() func
func Main() {
	ctx := context.Background() // root context

	go httpserver.Worker(ctx)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit // wait for quit here...
	log.Printf("Gracefull HTTP server shutdown max %d sec...", config.Http.GracefulShutdownTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Http.GracefulShutdownTimeout)*time.Second)
	defer func() {
		log.Print("HTTP server stopped")
		cancel()
	}()

	if err := httpserver.EchoServer.Shutdown(ctx); err != nil {
		// e.Logger.Fatal(err)
		log.Print("HTTP server shutdown error: ", err)
	} else {
		log.Print("HTTP server shutdown ok")
		// закрываем другие сервисы...
	}

}
