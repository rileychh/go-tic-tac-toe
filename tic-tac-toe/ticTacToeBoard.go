package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type ticTacToeBoard [3][3]cell

type BoardIndex struct {
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

func (b *ticTacToeBoard) GetByIndex(index BoardIndex) (cell, error) {
	row, column := index.row, index.column
	if (row >= len(ticTacToeBoard{}) || column >= len(ticTacToeBoard{}[0])) {
		return 0, errors.New("index out of range")
	}

	return b[row][column], nil
}

func (b *ticTacToeBoard) SetByIndex(index BoardIndex, value cell) error {
	row, column := index.row, index.column
	if (row >= len(ticTacToeBoard{}) || column >= len(ticTacToeBoard{}[0])) {
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

func (b *ticTacToeBoard) GetEmptyCells() (emptyCells []BoardIndex) {
	for row := range b {
		for column := range b[row] {
			if b[row][column] == empty {
				emptyCells = append(emptyCells, BoardIndex{row, column})
			}
		}
	}
	return
}
