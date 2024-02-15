package main

import (
	"fmt"
	"strings"
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
