package main

import (
	"fmt"	
)



func commandMap(cfg *config, args []string) error{
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil{
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, location := range locationResp.Results{
		fmt.Println(location.Name)
	}

	return nil
}


func commandMapB(cfg *config, args []string) error{
	if cfg.Previous == nil{
		return fmt.Errorf("no previous page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil{
		return err
	}
	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, location := range locationResp.Results{
		fmt.Println(location.Name)
	}

	return nil

	
}



func commandExplore(cfg *config, args []string) error{
	if len(args) < 1{
		return fmt.Errorf("please provide a location area name")
	}

	locationName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil{
		return err
	}

	fmt.Printf("Location Area: %s\n", locationArea.Name)
	fmt.Println("Pokemon Encounters:")
	for _, encounter := range locationArea.PokemonEncounters{
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
	
}