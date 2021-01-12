package chapter1

/*
Iterator pattern
for i := 0; i < n; i++ {}を抽象化して考える。
要素を順番にスキャンしていくインターフェース（API)を定める
*/

type Iterator interface {
	// check if next element exists
	HasNext() bool
	// return next element and make a next step(i++)
	Next() *Book
}
