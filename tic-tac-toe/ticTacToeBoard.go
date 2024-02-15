package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type ticTacToeBoard [3][3]cell

type boardIndex struct {
	row    int
	column int
}

func (b *ticTacToeBoard) String() string {
	var sb strings.Builder

	sb.WriteString("  A B C \n")
	sb.WriteString(" ╭─┬─┬─╮\n")
	for rowIndex := range b {
		sb.WriteString(fmt.Sprintf("%d", rowIndex+1))
		for _, column := range b[rowIndex] {
			sb.WriteString("│")
			switch column {
			case empty:
				sb.WriteString(" ")
			case circle:
				sb.WriteString("O")
			case cross:
				sb.WriteString("X")
			}
		}
		sb.WriteString("│\n")
		if rowIndex != len(b)-1 {
			sb.WriteString(" ├─┼─┼─┤\n")
		}
	}
	sb.WriteString(" ╰─┴─┴─╯\n")

	return sb.String()
}

// Convert "A1" to (0, 0), "C2" to (2, 1), etc.
func parseCoordinate(coordinate string) (boardIndex, error) {
	if len(coordinate) != 2 {
		return boardIndex{}, errors.New("invalid coordinate length")
	}

	letterPart := unicode.ToUpper(rune(coordinate[0]))
	numberPart := rune(coordinate[1])
	if !unicode.IsLetter(letterPart) || !unicode.IsDigit(numberPart) {
		return boardIndex{}, errors.New("invalid coordinate format")
	}

	column := int(letterPart - 'A')
	row := int(numberPart - '0' - 1)
	if (row >= len(ticTacToeBoard{}) || column >= len(ticTacToeBoard{}[0])) {
		return boardIndex{}, errors.New("coordinate out of range")
	}

	return boardIndex{row, column}, nil
}

func (b *ticTacToeBoard) SetByCoordinate(coordinate string, value cell) error {
	index, err := parseCoordinate(coordinate)
	if err != nil {
		return err
	}

	row, column := index.row, index.column
	if b[row][column] != empty {
		return errors.New("cell is occupied")
	}

	b[row][column] = value
	return nil
}
