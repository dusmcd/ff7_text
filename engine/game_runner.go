package engine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type CurrentGameState struct {
	Characters []Character
	Enemies    []Enemy
}

func PlayGame(gameMoment int, gameState CurrentGameState) {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	ticker := InitiateRandomEncounters(5, scanner)
	for {
		select {
		case <-ticker.C:
			player := gameState.Characters[0]
			enemy := gameState.Enemies[0]
			startRandomBattle(ticker, player, enemy, scanner)
		default:
			fmt.Print("Commands> ")
			playerInput := getPlayerInput(scanner)
			if playerInput == "exit" {
				return
			}
			continueStory(gameMoment, playerInput)
		}
	}

}

func getPlayerInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func startRandomBattle(ticker *time.Ticker, player Character, enemy Enemy, scanner *bufio.Scanner) {
	StopRandomEncounters(ticker)
	battle := NewBattle(player, enemy)
	battle.Fight(scanner)
	RestartRandomEncounters(ticker, 5)
}

func continueStory(gameMoment int, playerInput string) {

	action := getCommand(playerInput)
	err := action.callback()
	if err != nil {
		fmt.Println("An error occurred")
		log.Println(err.Error())
	}
}
