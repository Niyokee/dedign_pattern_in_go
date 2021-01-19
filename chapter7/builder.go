package chapter7

// タイトルを一つ含む
// 文字列をいくつか含む
// 箇条書きの項目をいくつか含む

// Builderでは文書を構成するためのメソッドを定める
// Directorクラスがそのメソッドを使って具体的な文書を作る

// Builder
// 文書を作るメソッドを宣言する

type Builder interface {
	makeTitle(title string) string
	makeString(str string) string
	makeItems(items []string) string
	close() string
}

type Director struct {
	Builder Builder
}

func (d *Director) Construct() string {
	result := d.Builder.makeTitle("Greeting")
	result += d.Builder.makeString("朝から昼にかけて")
	result += d.Builder.makeItems([]string{"おはようございます", "こんにちは"})
	result += d.Builder.makeString("よるに")
	result += d.Builder.makeItems([]string{"こんばんは", "おやすみなさい", "さようなら"})
	result += d.Builder.close()
	return result
}

type TextBuilder struct {
}

func (t *TextBuilder) makeTitle(title string) string {
	return "#" + title + "\n"
}

func (t *TextBuilder) makeString(str string) string {
	return "## " + str + "\n"
}

func (t *TextBuilder) makeItems(items []string) string {
	var str string
	for _, item := range items {
		str += item + "\n"
	}
	return str
}

func (t *TextBuilder) close() string {
	return "###"
}
