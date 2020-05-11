package api

// Client interacts with 3rd party api
type Client interface {
	GetJoke() (*JokeResponse, error)
}
