package main

import (
	"fmt"

	"github.com/davidheeren/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	locations, err := cfg.client.ListLocations(cfg.nextUrl)
	if err != nil {
		return err
	}

	pageNum, err := pokeapi.LocationsPageNumber(cfg.nextUrl)
	if err != nil {
		return err
	}

	fmt.Printf("----PAGE %d----\n", pageNum)
	for _, l := range locations.Results {
		fmt.Println(l)
	}

	if locations.Next != "" {
		cfg.nextUrl = locations.Next
	}
	if locations.Previous != "" {
		cfg.prevUrl = locations.Previous
	}
	return nil
}

func commandMapB(cfg *config) error {
	locations, err := cfg.client.ListLocations(cfg.prevUrl)
	if err != nil {
		return err
	}

	pageNum, err := pokeapi.LocationsPageNumber(cfg.prevUrl)
	if err != nil {
		return err
	}

	fmt.Printf("----PAGE %d----\n", pageNum)
	for _, l := range locations.Results {
		fmt.Println(l)
	}

	if locations.Next != "" {
		cfg.nextUrl = locations.Next
	}
	if locations.Previous != "" {
		cfg.prevUrl = locations.Previous
	}
	return nil
}
