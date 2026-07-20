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
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for key, com := range getCommands() {
		fmt.Printf("%s: ", key)
		fmt.Println(com.description)
	}
	return nil
}
