package engine

import (
	"bufio"
	"fmt"
	"time"
)

type Battle struct {
	Player Character
	Enemy  Enemy
	Status string
}

type Game struct {
	BattleState bool
}

func (battle Battle) Fight(scanner *bufio.Scanner) {
	fmt.Println("You are fighting", battle.Enemy.Name)
	for {
		if battle.Status == "over" {
			// show result of battle, e.g., player won, player defeated, or player escaped
			fmt.Println("Battle over")
			break
		}
		fmt.Println("Choose an action:")
		fmt.Println("\tAttack")
		fmt.Print("Input action: ")
		scanner.Scan()
		battle.Player.Attack(&battle.Enemy)
		fmt.Printf("You attacked %s for %d HP\n", battle.Enemy.Name, 100)
		battle.Enemy.Attack(&battle.Player)
		fmt.Printf("%s attacked you for %d HP\n", battle.Enemy.Name, 10)
		if battle.Enemy.CurrentHP <= 0 || battle.Player.CurrentHP <= 0 {
			battle.Status = "over"
		}
	}
}

func NewBattle(player Character, enemy Enemy) Battle {
	return Battle{
		Player: player,
		Enemy:  enemy,
	}
}

func InitiateRandomEncounters(interval int, scanner *bufio.Scanner) *time.Ticker {
	// this will run concurrently with the main game logic
	// for now, I will set this on a timer that will go off at a specified interval
	// i.e., the player will encounter an enemy every 100 seconds (or whatever the interval is)
	ch := time.NewTicker(time.Duration(interval) * time.Second)
	return ch
}

func StopRandomEncounters(ticker *time.Ticker) {
	// this will stop the timer initiated by the the InitiateRandomEncounters function
	ticker.Stop()
}

func RestartRandomEncounters(ticker *time.Ticker, interval int) {
	ticker.Reset(time.Duration(interval) * time.Second)
}
