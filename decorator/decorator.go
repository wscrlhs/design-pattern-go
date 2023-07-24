package decorator

// IPizza 零件接口
type IPizza interface {
	GetPrice() int
}

// VeggeMania 具体零件
type VeggeMania struct {
}

func (p *VeggeMania) GetPrice() int {
	return 15
}

// TomatoTopping 具体装饰
type TomatoTopping struct {
	Pizza IPizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 7
}

// CheeseTopping 具体装饰
type CheeseTopping struct {
	Pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
