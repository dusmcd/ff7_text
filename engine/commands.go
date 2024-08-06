package engine

import (
	"fmt"
)

type command struct {
	name        string
	description string
	callback    func() error
}

func getAllCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "displays a list of commands available to the player",
			callback:    helpCommandCallback,
		},
	}
}

func helpCommandCallback() error {
	fmt.Println("Here are the available commands:")
	commands := getAllCommands()
	for command := range commands {
		fmt.Printf("  %s\t%s\n", commands[command].name, commands[command].description)
	}
	fmt.Println("  exit\texits the game")
	return nil
}

func getCommand(playerInput string) command {
	action, found := getAllCommands()[playerInput]
	if !found {
		return command{
			name:        "invalid",
			description: "invalid command",
			callback:    invalidCommandCallback,
		}
	}

	return action
}

func invalidCommandCallback() error {
	fmt.Println("Invalid command. Type 'help' for valid commands")
	return nil
}
