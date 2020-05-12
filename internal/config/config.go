package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type http struct {
	Port                    string `yaml:"port" env:"PORT" env-description:"HTTP port"`
	Host                    string `yaml:"host" env:"HOST" env-default:"0.0.0.0" env-description:"HTTP host"`
	GracefulShutdownTimeout int    `yaml:"GracefulShutdownTimeout" env:"GRACEFUL_SHUTDOWN_TIMEOUT" env-default:"30" env-description:"HTTP server graceful shutdown period"`
}

type jokeapp struct {
	JokeURL string `yaml:"joke-url" env:"JOKE_URL" env-description:"Joke URL"`
}

var Http http
var JokeApp jokeapp

// Init - loads configs
func Init() {
	if _, err := os.Stat("config.yml"); err == nil {
		{
			err := cleanenv.ReadConfig("config.yml", &Http)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("loaded config Http:%+v", Http)
		}
		{
			err := cleanenv.ReadConfig("config.yml", &Http)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("loaded config JokeApp:%+v", JokeApp)
		}
	} else {
		log.Print(err)
		log.Print("config.yml not found. hope everything is in env?")
	}
}
