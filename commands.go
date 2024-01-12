package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func doHelpCmd(c *config, _ ...string) error {
	fmt.Println("Welcome to the Gokedex!")
	fmt.Println("\nAvailable commands:")
	for _, cmd := range getCmds(c) {
		fmt.Printf(" - %s: %s\n", cmd.command, cmd.description)
	}

	return nil
}

func doExitCmd(_ *config, _ ...string) error {
	os.Exit(0)
	return nil
}

func doMapCmd(c *config, _ ...string) error {
	res, err := c.pokeapiClient.ListLocations(c.nextLocationsUrl)
	if err != nil {
		return err
	}

	c.nextLocationsUrl = res.Next
	c.prevLocationsUrl = res.Previous

	fmt.Println("Locations:")

	for _, loc := range res.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	return nil
}

func doMapBCmd(c *config, _ ...string) error {
	if c.prevLocationsUrl == nil {
		return errors.New("you're already on the first page")
	}
	res, err := c.pokeapiClient.ListLocations(c.prevLocationsUrl)
	if err != nil {
		return err
	}

	c.nextLocationsUrl = res.Next
	c.prevLocationsUrl = res.Previous

	fmt.Println("Locations:")

	for _, loc := range res.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	return nil
}

func doExploreCmd(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a location to explore")
	}
	locationName := args[0]

	res, err := c.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s\n", res.Name)

	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func doBattleCmd(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a Pokemon to battle")
	}
	pokemon := args[0]

	res, err := c.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	const threshold = 40
	randNum := rand.Intn(res.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemon)
	}

	c.caughtPokemon[res.Name] = res

	fmt.Printf("Caught %s\n", res.Name)

	return nil
}

func doInspectCmd(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a Pokemon to inspect")
	}
	pokemon := args[0]

	p, ok := c.caughtPokemon[pokemon]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("%s:\n", p.Name)
	fmt.Printf(" - height: %v\n", p.Height)
	fmt.Printf(" - weight: %v\n", p.Weight)
	for _, stat := range p.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println(" - types:")
	for _, t := range p.Types {
		fmt.Printf("   - %s\n", t.Type.Name)
	}
	return nil
}

func doDexCmd(c *config, _ ...string) error {
	if len(c.caughtPokemon) < 1 {
		return errors.New("gokedex is empty")
	}
	fmt.Println("Gokedex:")
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
