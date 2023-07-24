package abstractfactory_test

import (
	"design-pattern/abstractfactory"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdidas(t *testing.T) {
	adidasFactory, _ := abstractfactory.GetSportsFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	logo := "adidas shoe"
	size := 10
	adidasShoe.SetLogo(logo)
	adidasShoe.SetSize(size)

	assert.Equal(t, logo, adidasShoe.GetLogo())
	assert.Equal(t, size, adidasShoe.GetSize())

	logo = "adidas shirt"
	size = 20
	adidasShirt.SetLogo(logo)
	adidasShirt.SetSize(size)
	assert.Equal(t, logo, adidasShirt.GetLogo())
	assert.Equal(t, size, adidasShirt.GetSize())
}

func TestNike(t *testing.T) {
	NikeFactory, _ := abstractfactory.GetSportsFactory("nike")
	NikeShoe := NikeFactory.MakeShoe()
	NikeShirt := NikeFactory.MakeShirt()

	logo := "Nike shoe"
	size := 15
	NikeShoe.SetLogo(logo)
	NikeShoe.SetSize(size)

	assert.Equal(t, logo, NikeShoe.GetLogo())
	assert.Equal(t, size, NikeShoe.GetSize())

	logo = "Nike shirt"
	size = 25
	NikeShirt.SetLogo(logo)
	NikeShirt.SetSize(size)
	assert.Equal(t, logo, NikeShirt.GetLogo())
	assert.Equal(t, size, NikeShirt.GetSize())
}

func TestWrongBrandType(t *testing.T) {
	_, err := abstractfactory.GetSportsFactory("wrong")
	assert.EqualError(t, err, "Wrong brand type passed")
}

func Example() {
	adidasFactory, _ := abstractfactory.GetSportsFactory("adidas")
	nikeFactory, _ := abstractfactory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	// Output:
	// Logo: nike
	// Size: 14
	// Logo: nike
	// Size: 14
	// Logo: adidas
	// Size: 14
	// Logo: adidas
	// Size: 14
}

func printShoeDetails(s abstractfactory.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s abstractfactory.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
