package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to the tic-tac-toe game!")

	for {
		fmt.Print("Do you want to play as circle (O) or cross (X)? [O/X] ")

		var userInput string
		_, err := fmt.Scanln(&userInput)
		userInput = strings.ToUpper(userInput)
		if userInput == "O" || userInput == "X" {
			break
		} else if err != nil {
			fmt.Printf("Looks like there's an problem getting your input: %v.", err)
		} else {
			fmt.Print("You entered something I can't understand.")
		}
		fmt.Println(" Please try again.")
	}

	var board board
	for i, row := range board {
		for j := range row {
			if (i*len(row)+j)%2 == 0 {
				board[i][j] = circle
			} else {
				board[i][j] = cross
			}
		}
	}
	fmt.Println(board)
}
