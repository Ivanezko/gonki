package handler

import (
	"github.com/labstack/echo"
	"log"
	"main/internal/api"
	"net/http"
)

type Handler struct {
	jokeClient api.Client
}

func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient: jokeClient,
	}
}

func (h *Handler) Joke(c echo.Context) error {
	joke, err := h.jokeClient.GetJoke()
	if err != nil {
		log.Print(err)
		return c.HTML(http.StatusInternalServerError, "oops")
	}

	return c.HTML(http.StatusOK, joke.Joke)
}

func Health(c echo.Context) error {
	log.Print("health requested")
	return c.HTML(http.StatusOK, "OK")
}
