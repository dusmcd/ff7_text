package engine

import (
	"fmt"
)

type StoryBlock struct {
	Moment    int
	Narration string
	Dialogue  []map[string]string
	Room      string
}

func PrintStory(story []StoryBlock) {
	for _, block := range story {
		fmt.Println(block.Narration)
		for _, line := range block.Dialogue {
			for character := range line {
				fmt.Printf("%s: %s\n", character, line[character])
			}
		}
	}
}

func FindGameMoment(story []StoryBlock, gameMoment int) []StoryBlock {
	startingIndex := 0
	for i, block := range story {
		if block.Moment == gameMoment {
			startingIndex = i
		}
	}
	return story[startingIndex:]
}
