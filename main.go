package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// ANSI color codes for better visualization
const (
	bgRed    = "\033[41;1;37m" // Red background
	bgGreen  = "\033[42;1;37m" // Green background
	bgYellow = "\033[43;1;37m" // Yellow background
	Reset    = " \033[0m "     // Reset color
)

// Read words from a file and return them as a slice of strings
func readWordsFromFile() ([]string, error) {
	file, err := os.Open("./Wordle.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	return words, nil
}

// Pick a random word from a slice of strings
func pickRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

// Check if a character is present in a string
func contains(str string, substr byte) bool {
	for i := range str {
		if str[i] == substr {
			return true
		}
	}
	return false
}

// Compare the target word with the player's guess and provide feedback
func compareWords(word, guess string) string {
	var result string
	for i := 0; i < len(word); i++ {
		if i < len(guess) && word[i] == guess[i] {
			result += bgGreen + " " + string(word[i]) + Reset // Correct letter in correct position
		} else if contains(word, guess[i]) {
			result += bgYellow + " " + string(guess[i]) + Reset // Correct letter in wrong position
		} else {
			result += bgRed + " " + string(guess[i]) + Reset // Incorrect letter
		}
	}
	return result
}

// Display game status, show score, and ask if the player wants to play again or exit
func status(win bool, word string, playerscore int) bool {
	if !win {
		fmt.Printf("\n\n"+bgRed+"Sorry, you lose! The word was %s"+Reset, word)
	} else {
		fmt.Println(bgGreen + "You win!" + Reset)
	}
	fmt.Print("\nDo you want to play again? (y/n): ")
	var choice string
	fmt.Scanln(&choice)
	if choice == "y" {
		return true
	} else {
		fmt.Println("Thanks for playing Wordle!")
		return false
	}
}

func main() {
	// Read words from a file
	words, err := readWordsFromFile()
	if err != nil {
		fmt.Println("Error reading words file:", err)
		return
	}

	// Game introduction
	fmt.Println("\nWelcome to Wordle!")
	fmt.Println("Try to guess the 5-letter word.")
	fmt.Println("You have 6 attempts.")

	playerscore := 0
	// Game loop
	for {
		win := false
		score := 70
		// Print current player score
		fmt.Printf("\n\nYour score: %d\n\n", playerscore)

		// Pick a random word for the player to guess
		word := pickRandomWord(words)
		attempts := 0

		// Player guessing loop
		for attempts < 6 {
			fmt.Printf("Attempt %d: ", attempts+1)
			score -= 10
			var input string
			fmt.Scanln(&input)
			guess := strings.TrimSpace(strings.ToLower(input))

			// Validate guess
			if len(guess) != 5 {
				fmt.Println("Please enter a 5-letter word.")
				continue
			}
			attempts++

			// Check if guess is correct
			if word == guess {
				playerscore += score
				win = true
				break
			}

			// Provide feedback on the guess
			result := compareWords(word, guess)
			fmt.Println(result)
		}

		// Game over, display result and ask for replay
		if attempts == 6 || win {
			if !status(win, word ,playerscore) {
				break
			}
		}
	}
}
