package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandInspect(cfg *config, args []string)error{
	if len(args) < 1 {
		return fmt.Errorf("please provide provide a pokemon to inspect")
	}

	pokemon := args[0]

	pokeMonCaught, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil{
		return fmt.Errorf("error occured while getting caught pokemon %s", err)
	}

	// Handle the  logic here
	pokeCaught, ok := cfg.pokeMon[pokeMonCaught.Name]

	if !ok {
		fmt.Println("you have not caught that pokemon")
	}else{
		fmt.Printf("Name : %s\n", pokeCaught.Name)
		fmt.Printf("Height : %v\n", pokeCaught.Height)
		fmt.Printf("Weight : %v\n", pokeMonCaught.Weight)
		fmt.Println("Stats")
		for _,value := range pokeMonCaught.Stats{
			fmt.Printf(" -%s: %v\n", value.Stat.Name, value.BaseStat)
		}
		fmt.Println("Types")
		for _,v := range pokeCaught.Types{
			fmt.Println(" -",v.Type.Name)
		}
	}

	return nil

}	