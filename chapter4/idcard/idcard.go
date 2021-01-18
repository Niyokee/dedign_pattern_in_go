package idcard

import (
	"fmt"
)

type IDCard struct {
	owner string
	//factory.Product
}

func (i IDCard) Use()  {
	fmt.Println(i.owner, "のカードを使います")
}

func (i IDCard) GetOwner() string {
	return i.owner
}

func NewIDCard(owner string) *IDCard {
	fmt.Println(owner + "のカードを作ります")
	i := &IDCard{owner: owner}
	// interfaceの実装
	return i
}

