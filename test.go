package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	var diffChoice int

	for {
		fmt.Print("Enter your choice (1, 2, or 3): ")
		input, _ := reader.ReadString('\n') // buffering the input
		input = strings.TrimSpace(input)    // removing remaining newline
		num, err := strconv.Atoi(input)

		// Checking if the type is right
		if err != nil {
			fmt.Println("Invalid input! Please enter a number between 1 and 3.")
			continue
		}

		// Checking choice number
		if num >= 1 && num <= 3 {
			diffChoice = num
			break
		}
	}

	var difficulty string
	var chances int

	switch diffChoice {
	case 1:
		difficulty = "Easy"
		chances = 10
	case 2:
		difficulty = "Medium"
		chances = 5
	default:
		difficulty = "Hard"
		chances = 3
	}

	fmt.Printf("Great! You have selected the %s difficulty level! You get %d chances!\n", difficulty, chances)
	fmt.Println("Let's start the game!")
}
