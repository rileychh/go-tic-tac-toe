package main

import (
	"html/template"
	"log"
	"net/http"
	"tic-tac-toe/internal/tictactoe"
)

type Page struct {
	Board tictactoe.Board
}

func main() {
	board := tictactoe.Board{}
	for row := range board {
		for col := range board[row] {
			if (row*len(board)+col)%2 == 0 {
				board[row][col] = tictactoe.Circle
			} else {
				board[row][col] = tictactoe.Cross
			}
		}
	}
	board[1][1] = tictactoe.Empty

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"cell": func(c tictactoe.Cell) string {
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
	page := Page{board}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.
			New("index.gohtml").
			Funcs(funcMap).
			ParseFiles("web/index.gohtml")
		if err != nil {
			log.Println("template.ParseFiles: ", err)
			writer.WriteHeader(500)
		}

		err = t.Execute(writer, page)
		if err != nil {
			log.Println("t.Execute: ", err)
			writer.WriteHeader(500)
		}
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
