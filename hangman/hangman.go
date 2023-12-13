package hangman

import (
	"errors"
	"fmt"
	"strings"
)

type State string

const (
	USED  State = "used"
	FOUND State = "found"
	GOOD  State = "good"
	BAD   State = "bad"
)

type Hangman struct {
	word        string
	lettersUsed []string
	letterFound []string
	State       State
	EndGame     bool
	Life        int
}

func New(word string) *Hangman {
	fmt.Printf("\n * * * HANGMAN * * * \n\n")
	lf := make([]string, len(word))
	return &Hangman{word: word, letterFound: lf, EndGame: false, Life: 6}
}

func (h *Hangman) Display() {
	fmt.Printf("Guessed: ")
	for i := range h.letterFound {
		value := h.letterFound[i]
		if value == "" {
			fmt.Printf("_ ")
		} else {
			fmt.Printf("%s ", value)
		}
	}
	fmt.Printf("\nUsed: ")
	for i := range h.lettersUsed {
		value := h.lettersUsed[i]
		fmt.Printf("%s", value)
	}
	h.printDraw()
}

func (h *Hangman) TryLetter(letter string) error {
	if len(letter) != 1 {
		return errors.New("want only one letter got: " + letter)
	}
	for _, v := range h.lettersUsed {
		if v == letter {
			h.State = USED
			return nil
		}
	}
	h.lettersUsed = append(h.lettersUsed, letter)

	if !strings.Contains(h.word, letter) {
		h.State = BAD
		h.Life--
		return nil
	}
	for i, v := range h.word {
		if v == rune(letter[0]) {
			h.letterFound[i] = letter
		}
	}
	h.isWin()

	h.State = GOOD

	return nil
}
func (h *Hangman) HandleState() bool {
	switch h.State {
	case FOUND:
		fmt.Println("Good guess!")
	case BAD:
		fmt.Println("Bad guess!")
	case USED:
		fmt.Println("Already used!")
		h.State = ""
		return true
	}
	return false
}

func (h *Hangman) isWin() {
	count := 0
	for _, v := range h.letterFound {
		if v == "" {
			break
		}
		count++
	}
	if count == len(h.word) {
		h.EndGame = true
	}
}

func (h *Hangman) printDraw() {
	switch h.Life {
	case 6:
		fmt.Println(level0)
	case 5:
		fmt.Println(level1)
	case 4:
		fmt.Println(level2)
	case 3:
		fmt.Println(level3)
	case 2:
		fmt.Println(level4)
	case 1:
		fmt.Println(level5)
	case 0:
		fmt.Println(level6)
		h.EndGame = true

	}
}

var level0 = ``
var level1 = `
________
`
var level2 = `
  |
  |
  |
__|______
`
var level3 = `
  _____
  |
  |
  |
__|______
`
var level4 = `
  _____
  |   |
  |   o
  |
__|______
`
var level5 = `
  _____
  |   |
  |   o
  |  -|-
__|______
`
var level6 = `
  _____
  |   |
  |   o
  |  -|-
__|__'_'___
`
