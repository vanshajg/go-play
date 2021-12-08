package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port string `env:"PORT" envDefault:":8080"`
}

func GetConfig(files ...string) (*Configuration, error) {
	err := godotenv.Load(files...)

	if err != nil {
		log.Printf("No .env file could be found %q\n", files)
	}

	cfg := Configuration{}

	err = env.Parse(&cfg)

	if err != nil {
		log.Printf("no .env file present in %q", files)
		return nil, err
	}
	return &cfg, nil
}
