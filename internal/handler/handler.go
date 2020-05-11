package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"main/internal/api"
	"net/http"
)

// Handler - common handler
type Handler struct {
	jokeClient api.Client
}

// NewHandler - create new handler
func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient: jokeClient,
	}
}

// Joke - get joke by api
func (h *Handler) Joke(c echo.Context) error {
	joke, err := h.jokeClient.GetJoke()
	if err != nil {
		log.Print(err)
		return c.HTML(http.StatusInternalServerError, "oops")
	}

	return c.HTML(http.StatusOK, joke.Joke)
}

// Health - get OK if server is running
func Health(c echo.Context) error {
	log.Print("health requested")
	return c.HTML(http.StatusOK, "OK")
}
