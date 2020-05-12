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
	"time"
)

var EchoServer *echo.Echo

// ServerContext - custom context
type ServerContext struct {
	echo.Context
}

// Worker - main HTTP server
func Worker(ctx context.Context) {
	_ = ctx

	apiClient := jokes.NewJokeClient(config.JokeApp.JokeURL)

	h := handler.NewHandler(apiClient)

	serverBind := config.Http.Host + ":" + config.Http.Port

	EchoServer = echo.New()
	EchoServer.HideBanner = false

	EchoServer.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ServerContext{c}
			return next(cc)
		}
	})

	EchoServer.Static("/", "assets")
	EchoServer.Use(middleware.Recover()) // recovers from panics

	EchoServer.GET("/health", handler.Health)
	EchoServer.GET("/joke", h.Joke)

	log.Print("Start HTTP server on addr: " + serverBind)
	s := &http.Server{
		Addr:         serverBind,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	EchoServer.Logger.Fatal(EchoServer.StartServer(s))
}
