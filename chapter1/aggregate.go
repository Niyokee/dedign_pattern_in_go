package chapter1

// 集合体を表すインターフェース
type Aggregate interface {
	Iterator() Iterator
}
