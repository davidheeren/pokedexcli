package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	// Do I need a name since this its the key in the map?
	// name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			description: "Displays the previous 20 location areas",
			callback:    commandMapB,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("error: %s\n", err)
				return
			}
			break
		}

		args := cleanInput(scanner.Text())
		if len(args) == 0 {
			continue
		}
		if com, ok := getCommands()[args[0]]; ok {
			err := com.callback()
			if err != nil {
				fmt.Printf("error: %s\n", err)
				return
			}
		} else {
			fmt.Printf("%s is not a valid command. Use the help command\n", args[0])
		}
	}
}
