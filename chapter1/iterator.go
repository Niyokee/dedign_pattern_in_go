package chapter1

// 1. iteratorインターフェースを宣言
// せめて、次の要素を取得するメソッドを持たないといけない
// 前の要素を取得したり、現在のいちを取得したり、イテレーションの最後かどうか
// 確認するメソッドがあっても良い。
type Iterator interface {
	// どんな要素を返すかわからないので、返り値はempty interface
	// Java ver.では Objectを返す実装になっている。
	Next() interface{}
	HasNext() bool
}

// 2. コレクションインターフェースを宣言
// 戻り値は、iterator interfaceでなければならない
// いくつかのiteratorを持つ場合はここで宣言する
type Aggregate interface {
	// Iterator Interfaceを満足させる
	// = 中身は問わないから、とにかく Next(), Prev(), HasNext()
	// を持つ
	Iterator() Iterator
}

type Book struct {
	Name string
}

// collection
type BookShelf struct {
	Books []*Book
}

// 4. コレクションインターフェースをコレクションに対して実装する
func (b *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: b}
}

func (b *BookShelf) Add(book *Book) {
	b.Books = append(b.Books, book)
}

// 3. iteratorの具象クラスを、コレクションに対して実装する。
type BookShelfIterator struct {
	BookShelf *BookShelf
	index int
}

func (b *BookShelfIterator) HasNext() bool {
	if b.index >= len(b.BookShelf.Books) {
		return false
	}
	return true
}

func (b *BookShelfIterator) Next() interface{} {
	book := b.BookShelf.Books[b.index]
	b.index++
	return book
}


