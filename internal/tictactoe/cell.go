package tictactoe

type Cell int

// What iota is: https://stackoverflow.com/a/14426447/11631322
const (
	Empty Cell = iota
	Circle
	Cross
)

func (c Cell) String() string {
	switch c {
	case Empty:
		return "empty"
	case Circle:
		return "circle"
	case Cross:
		return "cross"
	}
	return ""
}

func (c Cell) Display() string {
	switch c {
	case Empty:
		return " "
	case Circle:
		return "O"
	case Cross:
		return "X"
	}
	return ""
}
