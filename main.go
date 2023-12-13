package main

import (
	"fmt"
	"hangman/hangman"
	"hangman/utils"
	"log"
)

func main() {
	word, err := utils.GetWord()
	if err != nil {
		log.Fatal(err)
	}
	game := hangman.New(word)

	var entry string
	for !game.EndGame {
		//check if entry is already used
		if game.HandleState() {
			continue
		}
		game.Display()

		fmt.Println("What is your letter ?")
		_, err := fmt.Scanln(&entry)
		if err != nil {
			fmt.Println(err)
			game.State = ""
			continue
		}
		err = game.TryLetter(entry)
		if err != nil {
			fmt.Println(err)
			game.State = ""
			continue
		}
	}

	// End of game
	fmt.Println("The word was: " + word)
	if game.Life < 1 {
		fmt.Println(" * * * GAME OVER * * * ")
	} else {
		fmt.Println(" * * * YOU WIN * * * ")
	}

}
