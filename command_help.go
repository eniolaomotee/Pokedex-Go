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

	fmt.Printf("Throwing a Pokeball at %s ...", pokemonName)


	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil{
		return err
	}
	fmt.Printf("You caught a %s!\n", pokemon.Name)
	return nil
}