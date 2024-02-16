package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand/v2"
	"net/http"
	"tic-tac-toe/internal/tictactoe"
)

type Page struct {
	Board *tictactoe.Board
}

type Game struct {
	board    tictactoe.Board
	userMark tictactoe.Cell
	funcMap  template.FuncMap
}

func (game *Game) index(writer http.ResponseWriter, _ *http.Request) {
	t, err := template.
		New("index.gohtml").
		Funcs(game.funcMap).
		ParseFiles("web/index.gohtml", "web/board.gohtml")
	if err != nil {
		log.Println("template.ParseFiles:", err)
		writer.WriteHeader(500)
	}

	err = t.Execute(writer, Page{&game.board})
	if err != nil {
		log.Println("t.Execute: ", err)
		writer.WriteHeader(500)
	}
}

func (game *Game) mark(writer http.ResponseWriter, request *http.Request) {
	// Parse the ID of the cell
	cellId := request.Header.Get("HX-Trigger")
	index, err := tictactoe.ParseCoordinate(cellId)
	if err != nil {
		log.Println("/mark: ParseCoordinate:", err)
		writer.WriteHeader(400)
		return
	}

	// Check if cell is in range and empty
	cell, err := game.board.GetByIndex(index)
	if err != nil || cell != tictactoe.Empty {
		writer.WriteHeader(400)
		return
	}

	// Mark and check if the user won
	_ = game.board.SetByIndex(index, game.userMark) // won't error: index was checked in GetByIndex
	if game.checkAndEnd(writer) {
		return
	}

	// Check if the board is full
	emptyCells := game.board.GetEmptyCells()
	if len(emptyCells) == 0 {
		alert(writer, "Game Over. It's a draw.")
		game.reset(writer, request)
		return
	}

	// Set the index and mark for the computer
	index = emptyCells[rand.IntN(len(emptyCells))]
	computerMark := tictactoe.Circle
	if game.userMark == tictactoe.Circle {
		computerMark = tictactoe.Cross
	}

	// Mark and check if the computer won
	_ = game.board.SetByIndex(index, computerMark) // won't error: index is from GetEmptyCells
	if game.checkAndEnd(writer) {
		return
	}

	// Render and write the updated board
	game.writeBoard(writer)
}

func (game *Game) reset(writer http.ResponseWriter, _ *http.Request) {
	game.board = tictactoe.Board{}
	game.writeBoard(writer)
}

func main() {
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"renderCell": func(c tictactoe.Cell) string {
			switch c {
			case tictactoe.Empty:
				return " "
			case tictactoe.Circle:
				return "O"
			case tictactoe.Cross:
				return "X"
			}
			return ""
		},
		"Empty": func() tictactoe.Cell {
			return tictactoe.Empty
		},
	}
	game := Game{
		tictactoe.Board{},
		tictactoe.Circle,
		funcMap,
	}

	http.HandleFunc("/", game.index)
	http.HandleFunc("/mark", game.mark)
	http.HandleFunc("/reset", game.reset)

	addr := "127.0.0.1:9090"
	fmt.Printf("Open http://%s in your brower to play the game!", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
