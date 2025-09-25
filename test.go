package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	rand.Seed(time.Now().UnixNano())
	gameNumber := rand.Intn(100) + 1

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

		// Checking if the type is right and choice number
		if err != nil || num < 1 || num > 3 {
			fmt.Println("Invalid input! Please enter a number between 1 and 3.")
			continue
		}

		diffChoice = num
		break

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

	var guessNumber int
	attempts := 0

	for {
		if chances < 1 {
			fmt.Printf("Too bad! You exhausted all of your chances! The correct number is %d.\n", gameNumber)
			break
		}
		fmt.Print("Enter your guess number: ")
		input, _ := reader.ReadString('\n') // buffering the input
		input = strings.TrimSpace(input)    // removing remaining newline
		num, err := strconv.Atoi(input)

		chances--  // decrease the chances after avery input
		attempts++ // Increase the number of attempts

		// Checking if the type is right
		if err != nil {
			fmt.Println("Invalid input! Please enter a number between 1 and 100.")
			fmt.Printf("You have %d chances remaining.\n", chances)
			continue
		}

		// Checking if the number is within range
		if num < 1 || num > 100 {
			fmt.Println("Out of range! Please enter a number between 1 and 100!")
			fmt.Printf("You have %d chances remaining.\n", chances)
		}

		// Assigning number to main variable
		guessNumber = num

		if guessNumber == gameNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts!\n", attempts)
			break
		} else if guessNumber < gameNumber {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guessNumber)
			if abs(guessNumber-gameNumber) <= 5 {
				fmt.Println("You're almost there! The number is closer than you think!")
			}
			fmt.Printf("You have %d remaining.\n", chances)
			continue
		} else {
			fmt.Printf("Incorrect! The number is lower than %d.\n", guessNumber)

			if abs(guessNumber-gameNumber) <= 5 {
				fmt.Println("You're almost there! The number is closer than you think!")
			}
			fmt.Printf("You have %d chances remaining.\n", chances)
			continue
		}

	}
	fmt.Println("Thank you for playing the game!")
}
