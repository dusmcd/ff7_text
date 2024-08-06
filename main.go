package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dusmcd/ff7_text/engine"
)

func main() {
	fmt.Println("Welcome to Final Fantasy 7 Text-based adventure!")
	players, enemies, err := getInitData()
	if err != nil {
		log.Fatal(err)
	}

	currentGameState := engine.CurrentGameState{
		Characters: players,
		Enemies:    enemies,
	}
	engine.PlayGame(0, currentGameState)
	fmt.Println("Thank you for playing!")
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
