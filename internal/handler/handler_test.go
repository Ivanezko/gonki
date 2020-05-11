package handler_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"main/internal/api"
	"main/internal/api/mocks"
	"main/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Joke(t *testing.T) {
	tests := []struct {
		name     string
		joke     *api.JokeResponse
		err      error
		codeWant int
		bodyWant string
	}{
		{
			name:     "simple test",
			joke:     &api.JokeResponse{Joke: "test joke"},
			err:      nil,
			codeWant: http.StatusOK,
			bodyWant: "test joke",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := mocks.Client{}
			apiMock.On("GetJoke").Return(tt.joke, tt.err)
			h := handler.NewHandler(&apiMock)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/joke")

			/*
				req, _ := http.NewRequest("GET", "/joke", nil)
				rr := httptest.NewRecorder()
				handler := http.HandleFunc(h.Joke)
				handler.
				h.Joke(req.Context())
			*/
			if assert.NoError(t, h.Joke(c)) {
				assert.Equal(t, tt.codeWant, rec.Code)
				assert.Equal(t, tt.bodyWant, rec.Body.String())
			}
		})
	}
}
