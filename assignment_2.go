package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Basic input functions
	reader := bufio.NewReader(os.Stdin)

	// Prompt user for the first word
	fmt.Print("Enter the first word: ")
	word1, _ := reader.ReadString('\n')
	word1 = strings.TrimSpace(word1)

	// Prompt user for the second word
	fmt.Print("Enter the second word: ")
	word2, _ := reader.ReadString('\n')
	word2 = strings.TrimSpace(word2)

	// Implement word compare/character check logic
	if len(word1) != len(word2) {
		fmt.Println("The words are of different lengths and cannot be compared character by character.")
	} else {
		fmt.Println("Character comparison results:")
		for i := 0; i < len(word1); i++ {
			if word1[i] == word2[i] {
				fmt.Printf("Character %d: %c == %c\n", i+1, word1[i], word2[i])
			} else {
				fmt.Printf("Character %d: %c != %c\n", i+1, word1[i], word2[i])
			}
		}
	}
}
