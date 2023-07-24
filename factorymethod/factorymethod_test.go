package factorymethod_test

import (
	"design-pattern/factorymethod"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAk47(t *testing.T) {
	ak47, _ := factorymethod.GetGun("ak47")
	name := "ak47 888"
	power := 888
	ak47.SetName(name)
	ak47.SetPower(power)

	assert.Equal(t, name, ak47.GetName())
	assert.Equal(t, power, ak47.GetPower())
	assert.Equal(t, "this is "+ak47.GetName(), ak47.Show())
}

func TestMusket(t *testing.T) {
	musket, _ := factorymethod.GetGun("musket")
	name := "musket 666"
	power := 666
	musket.SetPower(power)

	musket.SetName(name)

	assert.Equal(t, name, musket.GetName())
	assert.Equal(t, power, musket.GetPower())
	assert.Equal(t, "this is "+musket.GetName(), musket.Show())
}

func TestWrongGunType(t *testing.T) {
	_, err := factorymethod.GetGun("we")
	assert.EqualError(t, err, "Wrong gun type passed")
}
func Example() {
	ak47, _ := factorymethod.GetGun("ak47")
	musket, _ := factorymethod.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)

	// Output:
	// Gun: AK47 gun
	// Power: 4
	// Gun: Musket gun
	// Power: 1
}

func printDetails(g factorymethod.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
