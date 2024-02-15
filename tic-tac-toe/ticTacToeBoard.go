package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type ticTacToeBoard [3][3]cell

func (b *ticTacToeBoard) String() string {
	var sb strings.Builder

	sb.WriteString("  A B C \n")
	sb.WriteString(" ╭─┬─┬─╮\n")
	for i, row := range b {
		sb.WriteString(fmt.Sprintf("%d", i+1))
		for _, column := range row {
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
		if i != len(b)-1 {
			sb.WriteString(" ├─┼─┼─┤\n")
		}
	}
	sb.WriteString(" ╰─┴─┴─╯\n")

	return sb.String()
}

// Convert "A1" to (0, 0), "C2" to (2, 1), etc.
func parseCoordinate(coordinate string) (int, int, error) {
	if len(coordinate) != 2 {
		return 0, 0, errors.New("invalid coordinate length")
	}

	column := unicode.ToUpper(rune(coordinate[0]))
	row := rune(coordinate[1])
	if !unicode.IsLetter(column) || !unicode.IsDigit(row) {
		return 0, 0, errors.New("invalid coordinate format")
	}

	x := int(column - 'A')
	y := int(row - '0' - 1)
	if (y >= len(ticTacToeBoard{}) || x >= len(ticTacToeBoard{}[0])) {
		return 0, 0, errors.New("coordinate out of range")
	}

	return y, x, nil
}

func (b *ticTacToeBoard) SetByCoordinate(coordinate string, value cell) error {
	row, column, err := parseCoordinate(coordinate)
	if err != nil {
		return err
	}
	if b[row][column] != empty {
		return errors.New("cell is occupied")
	}

	b[row][column] = value
	return nil
}
