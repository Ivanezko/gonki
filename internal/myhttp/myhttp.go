package myhttp

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"main/internal/config"
	"net/http"
	"sync"
	"time"
)

func httpHealth(c echo.Context) error {
	log.Print("health requested")
	return c.HTML(http.StatusOK, "OK")
}

// Server - main HTTP server
func Server(wg *sync.WaitGroup) {
	defer wg.Done()

	serverBind := config.Server.Host + ":" + config.Server.Port
	e := echo.New()
	e.HideBanner = false
	e.Use(middleware.Recover()) // recovers from panics
	e.GET("/health", httpHealth)
	log.Print("Start HTTP server on host: " + serverBind)
	s := &http.Server{
		Addr:         serverBind,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}
