package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct{
	name string
	description string
	callback func() error
}

func main(){

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanned := scanner.Scan()
		if !scanned{
			return 
		}
		word := scanner.Text()
		text := strings.ToLower(word)
		firstWord := strings.Fields(text)

		cmd := firstWord[0]

		if command, ok := getCommands()[cmd]; ok{
			if err := command.callback(); err != nil{
				fmt.Println("error", err)
			}
			continue
		}

		fmt.Println("Unknown command")
		continue
		
	}

}


func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}

