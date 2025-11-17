package main

import (
	//"bufio"
	"fmt"
	//"os"
	"strings"
	"math/rand"
	"github.com/eniolaomotee/Pokedex-Go/internal/pokeapi"
	"github.com/mattn/go-tty"
)


const prompt = "Pokedex > "




type config struct {
	Next *string
	Previous *string
	pokeapiClient *pokeapi.Client
	cmd string
	args []string
	rng *rand.Rand
	pokeMon map[string]pokeapi.CatchPokeMon
}



type cliCommand struct{
	name string
	description string
	callback func(*config, []string) error
}



func StartRepl(cfg *config){
	// scanner := bufio.NewScanner(os.Stdin)

	// define a channel to receive from keypress function
	lines := make(chan string)

	go keyPress(lines)

	for word := range lines {
		// fmt.Print("Pokedex > ")
		// scanned := scanner.Scan()
		// if !scanned{
		// 	return 
		// }
		// word := scanner.Text()
		text := CleanInput(word)
		if len(text) == 0 {
			continue
		}

		cmd := text[0]
		cfg.cmd = cmd
		if len(text) > 1{
			cfg.args = text[1:]
		}else{
			cfg.args = []string{}
		}
		

		if command, ok := getCommands()[cmd]; ok{
			if err := command.callback(cfg, cfg.args); err != nil{
				fmt.Println("error", err)
			}
			continue
		}

		fmt.Println("Unknown command")
		continue
		
	}

}



func CleanInput(text string) []string{
	cleanText := strings.ToLower(text)
	words := strings.Fields(cleanText)
	return words
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
		"map":{
			name: "map",
			description: "Location areas of Pokemon map",
			callback: commandMap,
		},
		"mapb":{
			name :"mapb",
			description: "Previous page of location areas of Pokemon map",
			callback: commandMapB,
		},
		"explore":{
			name: "explore",
			description: "explore the location area of a Pokemon",
			callback: commandExplore,
		},
		"catch":{
			name :"catch",
			description: "catch a Pokemon in a location area",
			callback: commandCatch,
		},
		"inspect":{
			name: "inspect",
			description: "inspect a pokemon",
			callback: commandInspect,
		},
		"pokedex":{
			name: "pokedex",
			description: "List all the Pokemon user caught",
			callback: commandPokedex,
		},
	}
}



func keyPress(out chan <- string){

	const prompt = "Pokedex >" 


	tty, err := tty.Open()
	if err != nil{
		return 
	}
	defer tty.Close()

	history := []string{"help","exit","map","mapb","explore","catch","inspect","pokedex"}
	HIndex := len(history)
	
	var line string
	for {
		r, err := tty.ReadRune()
		if err != nil{
			return
		}

		switch r {

		case 3: // Ctrl + C
			fmt.Println("\n bye")
			return


		case 13:  // Enter
			fmt.Println()
			if line != "" {
				history = append(history, line)
				HIndex = len(history)
				out <- line
			}
			line = ""
			fmt.Print(prompt)
		
		case 127:  // Backspace
			if len(line) > 0{
				line = line[:len(line)-1]
				redraw(line)
			}
		
		
		case 27: // Esc prefix - possible arrow key
			// read the two runes : '[' and code
			r2, err2 := tty.ReadRune()
			if err2 != nil{
				continue
			}
			r3, err3 := tty.ReadRune()
			if err3 != nil{
				continue
			}


			if r2 == '['{
				switch r3 {
				case 'A': // Up arrow
					if HIndex > 0 {
						HIndex--
						line = history[HIndex]
						redraw(line)
					}
				case 'B': // Down arrow
					if HIndex < len(history) -1 {
						HIndex++
						line = history[HIndex]
					}else{
						HIndex = len(history)
						line = ""
					}
					redraw(line)
				}
			}
		
		default: 
			// printable char
			if r >= 32 { // ignore control chars
				line += string(r)
				redraw(line)

			}
			
		}

	}
}

func redraw(line string){
	// move to start, clear line,
	fmt.Print("\r\033[K")
    fmt.Print(prompt)
    fmt.Print(line)

}