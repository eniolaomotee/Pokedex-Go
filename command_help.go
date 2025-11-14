package main

import "fmt"

func commandHelp(cfg *config,args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _,cmd := range getCommands(){
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}


func commandCatch(cfg *config, args []string) error{
	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon name to catch")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil{
		return err
	}

	// Chance
	exp := float64(pokemon.BaseExperience)
	exp_min, exp_max := 50.0, 300.0
	
	t :=  (exp - exp_min) / (exp_max - exp_min)

	if t < 0 { t = 0}
	if t > 1 {t = 1}

	// Chance Bounds
	p_max, p_min := 0.85, 0.15
	p := 	p_max + t * (p_min - p_max) 
	
	if cfg.rng.Float64() < p{
		cfg.pokeMon[pokemon.Name] = pokemon
		fmt.Printf("You caught a %s\n",pokemon.Name)
	}else{
		fmt.Printf("%s escaped !\n",pokemon.Name)
	}

	return nil
}