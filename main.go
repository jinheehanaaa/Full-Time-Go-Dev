package main

import "fmt"

type Position struct {
	x, y int
}

func (p *Position) Move(val int) {
	fmt.Println("The position is moved by:", val)
}

type Player struct {
	Position
}

type Color int

// fmt.Stringer
func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "BLUE"
	case ColorBlack:
		return "BLACK"
	case ColorYellow:
		return "YELLOW"
	case ColorPink:
		return "PINK"
	default:
		panic("invalid color given")
	}
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	p := Player{}
	p.Move(1000)
	fmt.Println("the color is:", ColorBlack)
}
