package utils

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

func HandelWord(word string) (string, error) {
	word = strings.Trim(word, " ")
	if n := strings.Count(word, " "); n > 0 {
		return "", errors.New("Want only one word, got: " + word)
	}
	return word, nil
}

var words []string

func GetWord() (string, error) {
	file, err := os.Open("./words.txt")
	if err != nil {
		return "", err
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		words = append(words, scan.Text())
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(words))
	if len(words) > 0 {
		return words[i], nil
	} else {
		return "golang", nil
	}
}
