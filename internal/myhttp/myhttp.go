package myhttp

import (
	"github.com/ilyakaznacheev/cleanenv"
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

	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	serverBind := cfg.Host + ":" + cfg.Port
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
