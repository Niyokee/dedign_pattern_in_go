package factory

type Factorize interface {
	CreateProduct(owner string) *Product
	RegisterProduct(p *Product)
}

type Factory struct {
	factory Factorize
}

func (f Factory) NewProduct(owner string) *Product {
	p := f.factory.CreateProduct(owner)
	f.factory.RegisterProduct(p)
	return p
}