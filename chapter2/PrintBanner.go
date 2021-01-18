package chapter2


// ここで<i>Printerを実装する。
// type Bannerを変換する
type PrintBanner struct {
	*Banner
}

// そのまま使っているけど、ここのロジックで変換される
func (p PrintBanner) PrintWeak() {
	p.ShowWithParen()
}

func (p PrintBanner) PrintStrong()  {
	p.ShowWithAster()
}

func NewPrintBanner(s string) *PrintBanner {
	return &PrintBanner{&Banner{string: s}}
}
