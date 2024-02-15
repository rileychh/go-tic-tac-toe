# Go Tic-Tac-Toe

This project is a command-line based Tic-Tac-Toe game implemented in Go. It features both user and computer players with
a random turn order.

## Getting Started

1. Clone the repository

```bash
git clone git@github.com:rileychh/go-tic-tac-toe.git
```

2. Navigate to the project directory

```bash
cd go-tic-tac-toe
```

3. Run the game

```bash
go run .
```

## Game Rules

The game is played on a grid that's 3 squares by 3 squares. At the start of the game, you can choose to play as either
X or O. The computer will play as the other symbol. Players take turns putting their marks in empty squares. The first
player to get 3 of her marks in a row (horizontally, vertically, or diagonally) is the winner. When all 9 squares are
full, the game ends with a tie.

## Code Structure

- `tic_tac_toe_board.go`: This file contains the implementation of the Tic-Tac-Toe board, including the methods to get 
  and set the board cells, check for a win, and get the empty cells.
- `cell.go`: This file defines the cell type and its possible values (empty, circle, cross).
