package main

import (
	"DesignPattern/chapter1"
	"DesignPattern/chapter2"
	"DesignPattern/chapter3"
	"DesignPattern/chapter5"
	"DesignPattern/chapter7"
	"fmt"
)


func main() {
	t := chapter7.TextBuilder{}
	fmt.Printf("%T", t)
	d := chapter7.Director{&chapter7.TextBuilder{}}
	fmt.Println(d)
	fmt.Println(d.Construct())
}
func chapter5_main() {
	obj1 := chapter5.GetInstance()
	obj2 := chapter5.GetInstance()
	fmt.Println(obj1 == obj2)
}

func chapter3_main() {
	// Hを持ったCharDisplayのインスタンスを一個作る
	d1 := chapter3.NewCharDisplay('H')
	// Hello Worldを持ったStringDisplayのインスタンスを一個作る
	d2 := chapter3.NewStringDisplay("Hello World")
	d3 := chapter3.NewStringDisplay("こんにちは")

	d1.Display()
	d2.Display()
	d3.Display()

}

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

func chapter2_main() {
	s := "Hello"
	// ここでは<i>Printerを使ってプログラミングをしている
	// →実際に呼び出しているのはPrintBannerだがPrintの実装を持つ
	// ここでPrintBannerの存在は完全に隠されている。
	p := chapter2.NewPrintBanner(s)
	p.PrintWeak()
	p.PrintStrong()
}
