package builder_test

import (
	"design-pattern/builder"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormal(t *testing.T) {
	normalBuilder := builder.GetBuilder("normal")
	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	assert.Equal(t, builder.House{WindowType: "Wooden Window", DoorType: "Wooden Door", Floor: 2}, normalHouse)
}

func TestIgloo(t *testing.T) {
	iglooBuilder := builder.GetBuilder("igloo")
	director := builder.NewDirector(iglooBuilder)
	iglooHouse := director.BuildHouse()

	assert.Equal(t, builder.House{WindowType: "Snow Window", DoorType: "Snow Door", Floor: 1}, iglooHouse)
}

func TestNoBuilder(t *testing.T) {
	noBuilder := builder.GetBuilder("no")
	assert.Equal(t, nil, noBuilder)
}

func Example() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

	// Output:
	// Normal House Door Type: Wooden Door
	// Normal House Window Type: Wooden Window
	// Normal House Num Floor: 2
	//
	// Igloo House Door Type: Snow Door
	// Igloo House Window Type: Snow Window
	// Igloo House Num Floor: 1
}
