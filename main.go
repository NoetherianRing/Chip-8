package main

import (
	"github.com/NoetherianRing/Chip-8/app"
	"github.com/NoetherianRing/Chip-8/config"
	"github.com/faiface/pixel/pixelgl"
	"gopkg.in/yaml.v2"
	"os"
)

func run() {

	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	var cfg config.Config
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	myApp, err := app.NewApp(cfg)
	if err != nil {
		panic(err)
	}
	myApp.Run()
}
func main() {
	pixelgl.Run(run)
}
