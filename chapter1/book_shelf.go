package chapter1

type BookShelf struct {
	books []Book
	last  int
}

func NewBookShelf() *BookShelf {
	bs := new(BookShelf)
	return bs
}

func (b BookShelf) GetBookAt(i int) *Book {
	return &b.books[i]
}

func (b *BookShelf) AppendBook(book Book) {
	b.books = append(b.books, book)
	b.last++
}

func (b BookShelf) GetLen() int {
	return b.last
}

func (b BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(b)
}
