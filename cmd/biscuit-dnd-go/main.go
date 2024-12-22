package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/k-nox/biscuit-dnd-go/app"
	"github.com/k-nox/biscuit-dnd-go/config"
)

func main() {
	var cfg config.Config
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		panic(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		panic(err)
	}

	app.Run()
}
