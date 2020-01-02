package factory

import (
	"fmt"
)

type Pizza struct{}

type PizzaStore interface {
	CreatePizza() Pizza
}

type NYPizzaStore struct{}

func (p *NYPizzaStore) CreatePizza() {
	fmt.Println("NY Pizza")
}

type ChicagoPizzaStore struct{}

func (p *ChicagoPizzaStore) CreatePizza() {
	fmt.Println("Chicago Pizza")
}
