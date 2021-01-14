package chapter3

import "fmt"

// AbstractDisplayの実装
type CharDisplay struct {
	*AbstractDisplay
	Char rune // rune to be displayed
}

func (c *CharDisplay) open()  {
	fmt.Print("<<")
}

func (c *CharDisplay) print()  {
	fmt.Print(string(c.Char))
}

func (c *CharDisplay) close()  {
	fmt.Println(">>")
}
func NewCharDisplay(r rune) *CharDisplay {
	charDisplay := &CharDisplay{
		AbstractDisplay: &AbstractDisplay{},
		Char: r,
	}
	charDisplay.printer = charDisplay
	return charDisplay
}
