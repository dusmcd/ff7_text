package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dusmcd/ff7_text/engine"
)

func main() {
	fmt.Println("Welcome to Final Fantasy 7 Text-based adventure!")
	play()
	fmt.Println("Thank you for playing!")
}

func play() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	player := engine.NewCharacter("Cloud")
	enemy := engine.Enemy{
		Name: "Infantry",
		Health: engine.Health{
			CurrentHP: 50,
			MaxHP:     50,
			CurrentMP: 10,
			MaxMP:     10,
		},
	}
	battle := engine.Battle{
		Player: player,
		Enemy:  enemy,
	}
	battle.Fight()
	for {
		fmt.Print("Commands> ")
		playerInput := getPlayerInput(scanner)
		if playerInput == "exit" {
			break
		}

		action := getCommand(playerInput)
		err := action.callback()
		if err != nil {
			fmt.Println("An error occurred")
			log.Println(err.Error())
			continue
		}
	}
}

func getPlayerInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func getInitData() ([]engine.Character, []engine.Enemy, error) {
	data, err := os.ReadFile("init.json")
	if err != nil {
		return []engine.Character{}, []engine.Enemy{}, err
	}

	jsonData := engine.JsonData{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return []engine.Character{}, []engine.Enemy{}, err
	}

	characters := engine.MapJsonToCharacters(jsonData.Characters)
	enemies := engine.MapJsonToEnemies(jsonData.Enemies)

	return characters, enemies, nil
}
