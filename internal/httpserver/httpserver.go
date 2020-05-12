package httpserver

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"main/internal/api/jokes"
	"main/internal/config"
	"main/internal/handler"
	"net/http"
	"sync"
	"time"
)

// ServerContext - custom context
type ServerContext struct {
	echo.Context
}

// Worker - main HTTP server
func Worker(ctx context.Context, wg *sync.WaitGroup) {
	_ = ctx
	defer wg.Done()

	apiClient := jokes.NewJokeClient(config.Server.JokeURL)

	h := handler.NewHandler(apiClient)

	serverBind := config.Server.Host + ":" + config.Server.Port

	e := echo.New()
	e.HideBanner = false

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ServerContext{c}
			return next(cc)
		}
	})

	e.Static("/", "assets")
	e.Use(middleware.Recover()) // recovers from panics

	e.GET("/health", handler.Health)
	e.GET("/joke", h.Joke)

	log.Print("Start HTTP server on addr: " + serverBind)
	s := &http.Server{
		Addr:         serverBind,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}
