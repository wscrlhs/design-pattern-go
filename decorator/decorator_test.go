package decorator_test

import (
	"design-pattern/decorator"
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestDecorator(t *testing.T) {
	pizza := &decorator.VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	assert.Equal(t, 32, pizzaWithCheeseAndTomato.GetPrice())
}

func Example() {

	pizza := &decorator.VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())

	// // Output:
	// Price of veggeMania with tomato and cheese topping is 32
}
