package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for key, com := range getCommands() {
		fmt.Printf("%s: ", key)
		fmt.Println(com.description)
	}
	return nil
}
