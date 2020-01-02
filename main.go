package main

import (
	"fmt"
	// adapter "github.com/sdstolworthy/design_patterns/adapter"
	decorator "github.com/sdstolworthy/design_patterns/decorator"
	// strategy "github.com/sdstolworthy/design_patterns/strategy"
	factory "github.com/sdstolworthy/design_patterns/factory"
)

func main() {
	fmt.Println("starting")
	// strategy.RunStrategyDuck()
	// strategy.RunStrategyPhoto()
	// adapter.RunAdapterDrone()
	// decorator.RunDecoratorCoffee()
	decorator.RunPizzaDemo()
}
