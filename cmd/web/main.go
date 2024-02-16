package main

import (
	"html/template"
	"log"
	"net/http"
	"tic-tac-toe/internal/tictactoe"
)

type Page struct {
	Board *tictactoe.Board
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

func main() {
	board := tictactoe.Board{}
	userMark := tictactoe.Circle

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"renderCell": renderCell,
		"Empty": func() tictactoe.Cell {
			return tictactoe.Empty
		},
	}
	page := Page{&board}

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

	http.HandleFunc("/mark", func(writer http.ResponseWriter, request *http.Request) {
		// Parse the ID of the cell
		cellId := request.Header.Get("HX-Target")
		index, err := tictactoe.ParseCoordinate(cellId)
		if err != nil {
			log.Println("/mark: ParseCoordinate: ", err)
			writer.WriteHeader(403)
			return
		}

		// Check if cell is in range and empty
		cell, err := board.GetByIndex(index)
		if err != nil || cell != tictactoe.Empty {
			writer.WriteHeader(403)
			return
		}

		_ = board.SetByIndex(index, userMark)

		// Write the rendered cell
		_, err = writer.Write([]byte(renderCell(userMark)))
		if err != nil {
			log.Println("/mark: writer.Write: ", err)
			writer.WriteHeader(500)
			return
		}
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
