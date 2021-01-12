package chapter1

// 具体的な反復氏の役: Concrete Iterator
type BookShelfIterator struct {
	bookShelf BookShelf
	index     int
}

func (b *BookShelfIterator) BookShelfIterator(bs BookShelf) {
	b.bookShelf = bs
	b.index = 0
}

func NewBookShelfIterator(bs BookShelf) *BookShelfIterator {
	bsi := new(BookShelfIterator)
	bsi.bookShelf = bs
	return bsi
}

func (b *BookShelfIterator) HasNext() bool {
	if b.index < b.bookShelf.GetLen() {
		return true
	}
	return false
}

func (b *BookShelfIterator) Next() *Book {
	if b.HasNext() {
		book := b.bookShelf.GetBookAt(b.index)
		b.index++
		return book
	}
	return nil
}
