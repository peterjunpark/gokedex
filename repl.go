package main

import (
	"bufio"
	"fmt"
	"os"
)

type cmd struct {
	command     string
	callback    func(*config, ...string) error
	description string
}

func getCmds(c *config) map[string]cmd {
	return map[string]cmd{
		"help": {
			command:     "help",
			callback:    doHelpCmd,
			description: "List available commands and their usage",
		},
		"exit": {
			command:     "exit",
			callback:    doExitCmd,
			description: "Quit the Gokedex",
		},
		"map": {
			command:     "map",
			callback:    doMapCmd,
			description: "Show the next location area on the map",
		},
		"mapb": {
			command:     "mapb",
			callback:    doMapBCmd,
			description: "Show the previous location area on the map",
		},
		"explore": {
			command:     "explore {location}",
			callback:    doExploreCmd,
			description: "Explore a location",
		},
		"battle": {
			command:     "battle {pokemon}",
			callback:    doBattleCmd,
			description: "Battle and (try to) catch a Pokemon",
		},
		"inspect": {
			command:     "inspect {pokemon}",
			callback:    doInspectCmd,
			description: "Inspect a Pokemon in your Gokedex",
		},
		"dex": {
			command:     "dex",
			callback:    doDexCmd,
			description: "List Pokemon in your Gokedex",
		},
	}
}

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Gokedex • ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		cmdName := cleaned[0]
		args := []string{}
		if len(cleaned) >= 1 {
			args = cleaned[1:]
		}

		cmd, ok := getCmds(c)[cmdName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(c, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
