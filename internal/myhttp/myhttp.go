package myhttp

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"main/internal/api/jokes"
	"main/internal/config"
	"main/internal/handler"
	"net/http"
	"sync"
	"time"
)

// Server - main HTTP server
func Server(wg *sync.WaitGroup) {
	defer wg.Done()

	apiClient := jokes.NewJokeClient(config.Server.JokeURL)

	h := handler.NewHandler(apiClient)

	serverBind := config.Server.Host + ":" + config.Server.Port
	e := echo.New()
	e.HideBanner = false
	e.Use(middleware.Recover()) // recovers from panics
	e.GET("/health", handler.Health)
	e.GET("/joke", h.Joke)
	log.Print("Start HTTP server on host: " + serverBind)
	s := &http.Server{
		Addr:         serverBind,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}
