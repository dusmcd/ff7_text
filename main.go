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
	players, enemies, err := getInitData()
	if err != nil {
		log.Fatal(err)
	}
	ticker := engine.InitiateRandomEncounters(5, scanner)
	for {
		select {
		case <-ticker.C:
			engine.StopRandomEncounters(ticker)
			battle := engine.NewBattle(players[0], enemies[0])
			battle.Fight(scanner)
			engine.RestartRandomEncounters(ticker, 5)
		default:
			fmt.Print("Commands> ")
			playerInput := getPlayerInput(scanner)
			if playerInput == "exit" {
				return
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
