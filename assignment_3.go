package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// WordGame struct to hold the word and guessed letters
type WordGame struct {
	Word           string
	GuessedLetters []rune
}

// Method to check if a letter is in the word
func (wg *WordGame) CheckLetter(letter rune) bool {
	for _, l := range wg.Word {
		if l == letter {
			return true
		}
	}
	return false
}

// Method to display the current state of guessed letters
func (wg *WordGame) DisplayWord() string {
	display := ""
	for _, letter := range wg.Word {
		if contains(wg.GuessedLetters, letter) {
			display += string(letter) + " "
		} else {
			display += "_ "
		}
	}
	return display
}

// Helper function to check if a slice contains a rune
func contains(slice []rune, item rune) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Function to fetch a random word from the API
func fetchRandomWord() (string, error) {
	response, err := http.Get("https://random-word-api.herokuapp.com/word?number=1")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var words []string
	err = json.Unmarshal(body, &words)
	if err != nil {
		return "", err
	}

	if len(words) > 0 {
		return words[0], nil
	}

	return "", fmt.Errorf("no words received from API")
}

func main() {
	// Fetch a random word
	word, err := fetchRandomWord()
	if err != nil {
		fmt.Println("Error fetching random word:", err)
		return
	}

	word = strings.ToLower(word)
	fmt.Println("A random word has been fetched. Start guessing!")

	// Initialize the WordGame struct
	game := WordGame{
		Word:           word,
		GuessedLetters: []rune{},
	}

	reader := bufio.NewReader(os.Stdin)

	// Game loop
	for {
		fmt.Println("Current word:", game.DisplayWord())
		fmt.Print("Guess a letter: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		guess := rune(input[0])

		if contains(game.GuessedLetters, guess) {
			fmt.Println("You already guessed that letter.")
			continue
		}

		game.GuessedLetters = append(game.GuessedLetters, guess)

		if game.CheckLetter(guess) {
			fmt.Println("Good guess!")
		} else {
			fmt.Println("Sorry, that letter is not in the word.")
		}

		if game.DisplayWord() == strings.Join(strings.Split(word, ""), " ") {
			fmt.Println("Congratulations! You guessed the word:", word)
			break
		}
	}
}
