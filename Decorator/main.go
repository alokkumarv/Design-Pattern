package main

import "fmt"

type Coffiee interface {
	Cost() int
	Description() string
}

// Base Object

type Espresso struct{}

func (e *Espresso) Cost() int {
	return 50
}

func (e *Espresso) Description() string {
	return "This is espresso"
}

// Milk decorator on top of escpresso

type Milk struct {
	Coffiee Coffiee
}

func (m *Milk) Cost() int {
	return m.Coffiee.Cost() + 100
}

func (m *Milk) Description() string {
	return m.Coffiee.Description() + "with MILK"
}

// Suger Decorator on top of Milk

type Sugar struct {
	Coffiee Coffiee
}

func (s *Sugar) Cost() int {
	return s.Coffiee.Cost() + 200
}

func (s *Sugar) Description() string {
	return s.Coffiee.Description() + "with MILK "
}

func main() {

	// Create a base object
	coffee := &Espresso{}
	fmt.Println(coffee.Description(), ":", coffee.Cost())
	// Create a decorator Milk on top of expresso
	coffeeWithMilk := &Milk{Coffiee: coffee}
	fmt.Println(coffeeWithMilk.Description(), ":", coffeeWithMilk.Cost())

	// Create a decorator Sugar on top of Milk
	coffeeWithMilkandSugar := &Sugar{Coffiee: coffeeWithMilk}
	fmt.Println(coffeeWithMilkandSugar.Description(), ":", coffeeWithMilkandSugar.Cost())

}
