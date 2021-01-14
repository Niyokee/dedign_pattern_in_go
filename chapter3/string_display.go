package chapter3

import "fmt"

type StringDisplay struct {
	*AbstractDisplay
	string string
	width  int
}

func (s *StringDisplay) open()  {
	s.printLine()
}

func (s *StringDisplay) print()  {
	fmt.Printf("%s%s%s\n", "|", s.string, "|")
}

func (s *StringDisplay) close()  {
	s.printLine()
}

func (s *StringDisplay) printLine()  {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func NewStringDisplay(string string) *StringDisplay {
	stringDisplay := &StringDisplay{
		AbstractDisplay: &AbstractDisplay{},
		string:             string,
		width:           len(string),
	}
	stringDisplay.AbstractDisplay.printer = stringDisplay
	return stringDisplay
}