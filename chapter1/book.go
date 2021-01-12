package chapter1

type Book struct {
	name  string
}

func (b Book) GetName() string {
	return b.name
}

func NewBook(name string) *Book {
	b := new(Book)
	b.name = name
	return b
}
