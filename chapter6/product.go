package chapter6

import "fmt"

//  複製を可能にするためのもの
//  何を「使う」のかは実装に任されている
//　このインターフェースやManagerに具体的な名前が入っていないことに注意！
//  ソース中にクラスの名前を書くと、そのクラスとの密接な関係ができてしまう
type Product interface {
	use(s string)
	createClone() *Product
}

// Productインターフェースを利用して、インスタンスの複製を行う

type Manager struct {
	// インスタンスの名前とインスタンスの対応関係をmapで表現
	showCase map[string]*Product
}


// showCaseにいれるものは、なんでもいい。
// ただし、productインターフェースを満たすもの
func (m *Manager) register(name string, proto *Product) {
	m.showCase[name] = proto
}

func (m *Manager) create(name string) *Product {
	p := m.showCase[name]
	return p.createClone()
}

type MessageBox struct {
	decochar rune
}

func (m *MessageBox) use(s string)  {
	len := len(s)
	for i := 0; i < len + 4; i++ {
		fmt.Print(string(m.decochar))
	}
	fmt.Println("")
	fmt.Println(string(m.decochar), " ", s, " ", string(m.decochar))
	fmt.Println("")
}

func (m *MessageBox) createClone() *Product {
	var p *Product
	return p
}