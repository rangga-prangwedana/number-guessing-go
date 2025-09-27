package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"number-guessing-go/utils"
	"time"
)

type Game struct {
	targetNumber int
	difficulty   Difficulty
	attempts     int
	inputReader  *InputReader
}

func NewGame(reader *bufio.Reader) *Game {
	rand.Seed(time.Now().UnixNano())

	return &Game{
		targetNumber: rand.Intn(100) + 1,
		inputReader:  NewInputReader(reader),
	}
}

func (g *Game) Play() {
	g.showWelcome()
	g.selectDifficulty()
	g.startGameLoop()
	g.showGoodbye()
}

func (g *Game) showWelcome() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
}

func (g *Game) selectDifficulty() {
	fmt.Println("Please select the difficulty level: ")

	for level, diff := range GetAvailableDifficulties() {
		fmt.Printf("%d. %s (%d chances)\n", level, diff.Name, diff.Chances)
	}

	fmt.Print("Enter your choice (1, 2, or 3): ")
	choice := g.inputReader.ReadDifficultyChoice()

	g.difficulty, _ = GetDifficulty(choice)
	fmt.Printf("Great! You have selected the %s difficulty level! You get %d chances!\n", g.difficulty.Name, g.difficulty.Chances)
	fmt.Println("Let's start the game!")
}

func (g *Game) startGameLoop() {
	remainingChances := g.difficulty.Chances

	for remainingChances > 0 {
		fmt.Print("Enter your guess number: ")

		guess, valid := g.inputReader.ReadGuess()
		remainingChances--
		g.attempts++

		if !valid {
			g.showRemainingChances(remainingChances)
			continue
		}

		result := g.evaluateGuess(guess)
		if result == GuessCorrect {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts!\n", g.attempts)
			return
		}

		g.showRemainingChances(remainingChances)
	}

	fmt.Printf("Too bad! You exhausted all of your chances! The correct number is %d.\n", g.targetNumber)
}

type GuessResult int

const (
	GuessCorrect GuessResult = iota
	GuessTooLow
	GuessTooHigh
)

func (g *Game) evaluateGuess(guess int) GuessResult {
	if guess == g.targetNumber {
		return GuessCorrect
	}

	var hint string
	var result GuessResult

	if guess < g.targetNumber {
		hint = fmt.Sprintf("Incorrect! The number is greater than %d.", guess)
		result = GuessTooLow
	} else {
		hint = fmt.Sprintf("Incorrect! The number is lower than %d.", guess)
		result = GuessTooHigh
	}

	fmt.Println(hint)

	if g.isCloseGuess(guess) {
		fmt.Println("You're almost there! The number is closer than you think!")
	}

	return result
}

func (g *Game) isCloseGuess(guess int) bool {
	return utils.Abs(guess-g.targetNumber) <= 5
}

func (g *Game) showRemainingChances(remaining int) {
	if remaining > 0 {
		fmt.Printf("You have %d chances remaining.\n", remaining)
	}
}

func (g *Game) showGoodbye() {
	fmt.Println("Thank you for playing the game!")
}
