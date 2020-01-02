package decorator

import (
	"fmt"
)

type Pizza struct {
	description string
	cost        float32
}

type PizzaDecorator func(Pizza) Pizza

func (p *Pizza) PrintPizza() {
	fmt.Println(p.description, p.cost)
}

func WithThinCrust(p Pizza) Pizza {
	p.description += "Thin Crust Pizza"
	return p
}

func WithCheese(p Pizza) Pizza {
	p.description += ", cheese"
	p.cost += 1.5
	return p
}

func WithOlives(p Pizza) Pizza {
	p.description += ", olives"
	p.cost += 0.5
	return p
}

func RunPizzaDemo() {
	p := WithOlives(WithCheese(WithThinCrust(*new(Pizza))))
	p.PrintPizza()

}
