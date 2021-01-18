package factory

/*
product: 抽象メソッドuseのみ定義されている抽クラス
factory: メソッドcreatを実装している抽象クラス
IDCard: メソッドuseを実装している
IdCardFactory: メソッドcreateProduct, registerProductを実装しているクラス
 */

type product interface {
	use()
}

type factorize interface {
	createProduct(owner string) string
	registerProduct(p *product)
}
type factory struct {
	factorize
}

func (f factory) NewProduct(owner string) *product {
	p := f.createProduct(owner)
	f.registerProduct(p)
	return p
}