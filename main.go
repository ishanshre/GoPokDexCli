package main

import (
	"github.com/ishanshre/GoPokDexCli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAddress     *string
	previousLocationAddress *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startREPL(&cfg)
}
