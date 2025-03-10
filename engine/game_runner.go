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
	player := gameState.Characters[0]
	enemy := gameState.Enemies[0]

	for {
		select {
		case <-ticker.C:
			startRandomBattle(ticker, &player, &enemy, scanner)
		default:
			continueStory(gameMoment)
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
			}

		}
	}

}

func getPlayerInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func startRandomBattle(ticker *time.Ticker, player *Character, enemy *Enemy, scanner *bufio.Scanner) {
	StopRandomEncounters(ticker)
	battle := NewBattle(player, enemy)
	battle.Fight(scanner)
	RestartRandomEncounters(ticker, 5)
}

func continueStory(gameMoment int) {
	story := []StoryBlock{
		{
			Moment: 0,
			Narration: "A train speeds through a large metropolis. It screeches to a halt at a large energy plant\n" +
				"A young mercenary leaps off the top of the train while his companions wait for him",
			Dialogue: []map[string]string{
				{
					"Companion 1": "What's your name",
				},
				{
					"Mercenary": "I don't care about learning your name",
				},
			},
		},
		{
			Moment:    1,
			Narration: "The mercenary encounters two armed soldiers. He engages them in battle",
			Dialogue: []map[string]string{
				{
					"Barret": "Ex-Shinra, huh? I don't trust you",
				},
				{
					"Mercenary": "It doesn't matter what you think",
				},
			},
		},
	}

	currentStory := FindGameMoment(story, gameMoment)
	PrintStory(currentStory)
}
