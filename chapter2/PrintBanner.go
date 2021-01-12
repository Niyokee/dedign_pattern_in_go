package chapter2

type PrintBanner struct {
	Banner
}

func (p PrintBanner) PrintWeak() {
	p.ShowWithParen()
}

func (p PrintBanner) PrintStrong()  {
	p.ShowWithAster()
}

func NewPrintBanner(s string) *PrintBanner {
	p := new(PrintBanner)
	p.string = s
	return p
}
