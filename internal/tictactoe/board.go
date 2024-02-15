package tictactoe

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Board [3][3]Cell

type BoardIndex struct {
	row    int
	column int
}

func (b *Board) String() string {
	var sb strings.Builder

	sb.WriteString("  A B C \n")
	sb.WriteString(" ╭─┬─┬─╮\n")
	for row := range b {
		sb.WriteString(fmt.Sprintf("%d", row+1))
		for column := range b[row] {
			sb.WriteString("│")
			switch b[row][column] {
			case Empty:
				sb.WriteString(" ")
			case Circle:
				sb.WriteString("O")
			case Cross:
				sb.WriteString("X")
			}
		}
		sb.WriteString("│\n")
		if row != len(b)-1 {
			sb.WriteString(" ├─┼─┼─┤\n")
		}
	}
	sb.WriteString(" ╰─┴─┴─╯\n")

	return sb.String()
}

func (b *Board) GetByIndex(index BoardIndex) (Cell, error) {
	row, column := index.row, index.column
	if (row >= len(Board{}) || column >= len(Board{}[0])) {
		return 0, errors.New("index out of range")
	}

	return b[row][column], nil
}

func (b *Board) SetByIndex(index BoardIndex, value Cell) error {
	row, column := index.row, index.column
	if (row >= len(Board{}) || column >= len(Board{}[0])) {
		return errors.New("index out of range")
	}

	b[row][column] = value
	return nil
}

// ParseCoordinate Convert "A1" to (0, 0), "C2" to (2, 1), etc.
func ParseCoordinate(coordinate string) (BoardIndex, error) {
	if len(coordinate) != 2 {
		return BoardIndex{}, errors.New("invalid coordinate length")
	}

	letterPart := unicode.ToUpper(rune(coordinate[0]))
	numberPart := rune(coordinate[1])
	if !unicode.IsLetter(letterPart) || !unicode.IsDigit(numberPart) {
		return BoardIndex{}, errors.New("invalid coordinate format")
	}

	column := int(letterPart - 'A')
	row := int(numberPart - '0' - 1)
	return BoardIndex{row, column}, nil
}

func (b *Board) GetEmptyCells() (emptyCells []BoardIndex) {
	for row := range b {
		for column := range b[row] {
			if b[row][column] == Empty {
				emptyCells = append(emptyCells, BoardIndex{row, column})
			}
		}
	}
	return
}

func (b *Board) CheckWin() Cell {
	// Check rows
	for i := 0; i < 3; i++ {
		if b[i][0] != Empty && b[i][0] == b[i][1] && b[i][0] == b[i][2] {
			return b[i][0]
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if b[0][i] != Empty && b[0][i] == b[1][i] && b[0][i] == b[2][i] {
			return b[0][i]
		}
	}

	// Check diagonals
	if b[0][0] != Empty && b[0][0] == b[1][1] && b[0][0] == b[2][2] {
		return b[0][0]
	}
	if b[0][2] != Empty && b[0][2] == b[1][1] && b[0][2] == b[2][0] {
		return b[0][2]
	}

	// No winner
	return Empty
}
