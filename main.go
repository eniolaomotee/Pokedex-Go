package main

import (
	"time"

	"github.com/eniolaomotee/Pokedex-Go/internal/pokeapi"
)


func main(){
	pokeClient := pokeapi.NewClient(10*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	StartRepl(cfg)
}


