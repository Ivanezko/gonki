package config

// Server
type server struct {
	Port string `yaml:"port" env:"PORT"`
	Host string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
}

var Server server
