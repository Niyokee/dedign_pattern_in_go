package main

import (
	"DesignPattern/chapter1"
	"DesignPattern/chapter2"
	"fmt"
)

func chapter1_main() {
	bs := chapter1.NewBookShelf()
	bs.AppendBook(*chapter1.NewBook("Around the World in 80 days"))
	bs.AppendBook(*chapter1.NewBook("Bible"))
	bs.AppendBook(*chapter1.NewBook("Cinderella"))
	bs.AppendBook(*chapter1.NewBook("Daddy Long Legs*"))

	// bookShelfIteratorではなくIteratorを使ってプログラミングをしようとする姿勢・
	it := bs.Iterator()

	//　実装とは切り離して、数え上げができる。
	//　もし配列を使って本を管理することをやめたとしても、
	// この部分のコードの変更は不要。
	// bookShelfの利用者にとってはいいこと
	for it.HasNext() {
		b := it.Next()
		fmt.Println(b.GetName())
	}
}

func main() {
	s := "Hello"
	p := chapter2.NewPrintBanner(s)
	p.PrintWeak()
	p.PrintStrong()
}