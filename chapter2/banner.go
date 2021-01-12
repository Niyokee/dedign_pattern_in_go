package chapter2

import "fmt"

type Banner struct {
	string string
}

func (b Banner) ShowWithParen()  {
	fmt.Printf("(%v)", b.string)
}

func (b Banner) ShowWithAster()  {
	fmt.Printf("*%v*", b.string)
}

func NewBanner(s string) *Banner {
	b := new(Banner)
	b.string = s
	return b
}

