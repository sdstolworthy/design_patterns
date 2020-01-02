package strategy

import "fmt"

type FlyBehavior interface {
	Fly()
}

type QuackBehavior interface {
	Quack()
}

type LoudQuack struct{}

func (LoudQuack) Quack() {
	fmt.Println("QUACKKKKK")
}

type SwooshFly struct{}

func (SwooshFly) Fly() {
	fmt.Println("Swoosh!")
}

type DiveFly struct{}

func (DiveFly) Fly() {
	fmt.Println("Diving down, captain!")
}

type Duck struct {
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
}

func (d *Duck) Quack() {
	d.QuackBehavior.Quack()
}

func (d *Duck) Fly() {
	d.FlyBehavior.Fly()
}
