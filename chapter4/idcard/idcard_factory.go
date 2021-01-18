package idcard

import (
	"DesignPattern/chapter4/factory"
)

type Factory struct {
	*factory.Factory
	owners []string
}

// NewIDCardFactory func for initializing IDCardFactory
func NewIDCardFactory() *Factory {
	idCardFactory := &Factory{
		Factory: &factory.Factory{},
	}
	idCardFactory.Factory = idCardFactory
	return idCardFactory
}
func (i *Factory) CreateProduct(owner string) factory.Product {
	return NewIDCard(owner)
}

func (i *Factory) RegisterProduct(ic *factory.Product) {
	i.owners = append(i.owners, ic.GetOwner())
}