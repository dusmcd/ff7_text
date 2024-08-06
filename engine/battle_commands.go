package engine

import "fmt"

type BattleCommand struct {
	Name        string
	Description string
	Callback    func(*Character, *Enemy) error
}

func getBattleCommands() map[string]BattleCommand {
	return map[string]BattleCommand{
		"Attack": {
			Name:        "Attack",
			Description: "physical attack with weapon",
			Callback:    attackCallback,
		},
	}
}

func attackCallback(character *Character, target *Enemy) error {
	character.Attack(target)
	fmt.Printf("%s attacked %s for %d HP\n", character.Name, target.Name, 100)
	return nil
}

func invalidCallback(character *Character, target *Enemy) error {
	fmt.Println("Invalid command")
	return nil
}

func getBattleCommand(playerInput string) BattleCommand {
	command, found := getBattleCommands()[playerInput]
	if !found {
		return BattleCommand{
			Name:        "Invalid",
			Description: "Invalid command",
			Callback:    invalidCallback,
		}
	}
	return command

}
