package game

import (
	"bufio"
	"strconv"
	"strings"
)

type InputReader struct {
	reader *bufio.Reader
}

func NewInputReader(reader *bufio.Reader) *InputReader {
	return &InputReader{reader: reader}
}

func (ir *InputReader) ReadInt() (int, error) {
	input, err := ir.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

func (ir *InputReader) ReadDifficultyChoice() int {
	for {
		num, err := ir.ReadInt()
		if err != nil || num < 1 || num > 3 {
			println("Invalid input! Please enter a nubmer between 1 and 3.")
			print("Enter your choice (1, 2, or 3): ")
			continue
		}
		return num
	}
}

func (ir *InputReader) ReadGuess() (int, bool) {
	num, err := ir.ReadInt()

	if err != nil {
		println("Invalid input! Please enter a number between 1 and 100!")
		return 0, false
	}

	if num < 1 || num > 100 {
		println("Out of range! Please enter a number between 1 and 100!")
		return 0, false
	}

	return num, true
}
