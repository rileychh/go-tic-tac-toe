package main

type cell int

// What iota is: https://stackoverflow.com/a/14426447/11631322
const (
	empty cell = iota
	circle
	cross
)

func (c cell) String() string {
	switch c {
	case empty:
		return "empty"
	case circle:
		return "circle"
	case cross:
		return "cross"
	}
	return ""
}
