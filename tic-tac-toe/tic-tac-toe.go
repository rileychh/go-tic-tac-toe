package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func scanlnUntilSuccess(a ...any) (n int) {
	for attempt := 0; attempt < 5; attempt++ {
		var err error
		n, err = fmt.Scanln(a...)
		if err == nil {
			break
		}
		fmt.Printf("Looks like there's an problem getting your input: %v. Please try again.", err)
	}
	return n
}

func main() {
	fmt.Println("Welcome to the tic-tac-toe game!")

	var userMark, computerMark cell
chooseMark:
	for {
		fmt.Print("Do you want to play as circle (O) or cross (X)? [O/X] ")

		var userInput string
		scanlnUntilSuccess(&userInput)
		userInput = strings.ToUpper(userInput)

		switch userInput {
		case "O":
			userMark = circle
			computerMark = cross
			break chooseMark
		case "X":
			userMark = cross
			computerMark = circle
			break chooseMark
		default:
			fmt.Println("You entered something I can't understand. Please try again.")
		}
	}

	fmt.Printf("You'll be playing as %v, and the computer will play as %v.\n", userMark, computerMark)

	var board ticTacToeBoard
	var isUserTurn = rand.IntN(2) == 0

		fmt.Print(&board)
	for {
		if isUserTurn {
			fmt.Print("It's your turn. Enter a coordinate: ")
		} else {
			fmt.Println("It's the computer's turn.")
		}

		isUserTurn = !isUserTurn
		fmt.Print(&board)
	}
}
