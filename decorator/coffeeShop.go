package decorator

import "fmt"

type Beverage struct {
	Cost        float32
	Description string
}

func (b Beverage) PrintDescription() {
	fmt.Println("printing description")
	fmt.Println(b.Description, b.Cost)
}

func WithWhip(b Beverage) Beverage {
	b.Cost += 0.1
	b.Description += ", with whip"
	return b
}

func RunDecoratorCoffee() {
	beverage := Beverage{Cost: 0.0, Description: "Dark Roast"}
	beverage.PrintDescription()
	beverage = WithWhip(beverage)
	beverage.PrintDescription()
}
