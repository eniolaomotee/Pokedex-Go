package main

import (
	"math/rand"
	"time"

	"github.com/eniolaomotee/Pokedex-Go/internal/pokeapi"
)


func main(){
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	pokeClient := pokeapi.NewClient(10*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		rng: rng,
		pokeMon: make(map[string]pokeapi.CatchPokeMon),
	}

	StartRepl(cfg)
}


