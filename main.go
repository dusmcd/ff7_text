package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to Final Fantasy 7 Text-based adventure!")
	play()
	fmt.Println("Thank you for playing!")
}

func play() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
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
