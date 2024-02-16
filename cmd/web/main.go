package main

import (
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

func renderCell(c tictactoe.Cell) string {
	switch c {
	case tictactoe.Empty:
		return " "
	case tictactoe.Circle:
		return "O"
	case tictactoe.Cross:
		return "X"
	}
	return ""
}

func (game *Game) writeBoard(writer http.ResponseWriter) {
	t, err := template.
		New("board.gohtml").
		Funcs(game.funcMap).
		ParseFiles("web/board.gohtml")
	if err != nil {
		log.Println("template.ParseFiles:", err)
		writer.WriteHeader(500)
		return
	}

	err = t.Execute(writer, Page{&game.board})
	if err != nil {
		log.Println("t.Execute:", err)
		writer.WriteHeader(500)
	}
}

func (game *Game) index(writer http.ResponseWriter, request *http.Request) {
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

	_ = game.board.SetByIndex(index, game.userMark)

	emptyCells := game.board.GetEmptyCells()
	if len(emptyCells) != 0 {
		index := emptyCells[rand.IntN(len(emptyCells))]
		computerMark := tictactoe.Circle
		if game.userMark == tictactoe.Circle {
			computerMark = tictactoe.Cross
		}
		_ = game.board.SetByIndex(index, computerMark)
	}

	// Render and write the updated board
	game.writeBoard(writer)
}

func (game *Game) reset(writer http.ResponseWriter, request *http.Request) {
	game.board = tictactoe.Board{}
	game.writeBoard(writer)
}

func main() {
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"renderCell": renderCell,
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

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
