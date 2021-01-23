package main

import (
	"DesignPattern/chapter1"
	"DesignPattern/chapter2"
	"DesignPattern/chapter3"
	"DesignPattern/chapter5"
	"DesignPattern/chapter7"
	"fmt"
)


func main_() {
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

func main() {
	bs := &chapter1.BookShelf{}
	b1 := &chapter1.Book{"Around the world in 10 days"}
	b2 := &chapter1.Book{"Bible"}
	b3 := &chapter1.Book{"Daddy-Long-Legs"}
	bs.Add(b1)
	bs.Add(b2)
	bs.Add(b3)
	it := bs.Iterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println(book)
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
