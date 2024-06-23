package main

import "fmt"

func main() {
	// Variable Declaration and Standard Output Print Functions
	var a int = 10
	var b string = "Hello"
	var c float64 = 3.14
	var d bool = true

	e := 20
	f := "World"

	fmt.Println(a) // Output: 10
	fmt.Println(b) // Output: Hello
	fmt.Println(c) // Output: 3.14
	fmt.Println(d) // Output: true
	fmt.Println(e) // Output: 20
	fmt.Println(f) // Output: World

	fmt.Printf("a: %d, b: %s, c: %f, d: %t, e: %d, f: %s\n", a, b, c, d, e, f)

	// Basic Condition Statements and Data Types
	age := 21

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior.")
	}

	var x int = 5
	var y float64 = 7.5
	var z string = "Go Language"
	var w bool = true

	fmt.Printf("x: %d, y: %f, z: %s, w: %t\n", x, y, z, w)

	// Basic Looping Structure
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	k := 0
	for {
		fmt.Println(k)
		k++
		if k >= 5 {
			break
		}
	}

	// Initialize Entries Map
	entries := make(map[string]int)
	entries["Alice"] = 25
	entries["Bob"] = 30
	entries["Charlie"] = 35

	fmt.Println(entries)

	// Show Placeholder
	name := "Alice"
	agePlaceholder := 25
	height := 5.5

	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, agePlaceholder, height)

	// Basic Conditions/Evaluations
	num := 15

	if num%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	switch num {
	case 10:
		fmt.Println("The number is ten.")
	case 15:
		fmt.Println("The number is fifteen.")
	default:
		fmt.Println("The number is neither ten nor fifteen.")
	}

	// Print Guessed Letters
	word := "Golang"
	guessedLetters := []rune{'G', 'a', 'o'}

	for _, letter := range word {
		if contains(guessedLetters, letter) {
			fmt.Printf("%c ", letter)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
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
