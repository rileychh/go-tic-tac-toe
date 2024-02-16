package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"tic-tac-toe/internal/tictactoe"
)

func alert(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(`<script>alert("` + message + `")</script>`))
	if err != nil {
		log.Println("alert: Write:", err)
	}
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

func (game *Game) checkAndEnd(writer http.ResponseWriter) bool {
	if winner := game.board.CheckWin(); winner != tictactoe.Empty {
		script := fmt.Sprintf("Game Over. The winner is %v.", winner)
		alert(writer, script)
		game.reset(writer, nil)
		return true
	}
	return false
}
