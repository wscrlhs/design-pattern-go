// Package abstractfactory 抽象工厂模式
// 抽象工厂是一种创建型设计模式， 它能创建一系列相关的对象， 而无需指定其具体类。
// 包含系列中所有产品构造方法的接口，这些方法必须返回抽象产品类型。
// 抽象工厂，具体工厂，抽象产品，具体产品
package abstractfactory

import "fmt"

// ISportsFactory 抽象工厂接口
// 抽象工厂接口声明了一组能返回不同抽象产品的方法。这些产品属于同一个系列
// 且在高层主题或概念上具有相关性。同系列的产品通常能相互搭配使用。系列产
// 品可有多个变体，但不同变体的产品不能搭配使用。
type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

// Adidas 具体工厂
// 具体工厂可生成属于同一变体的系列产品。工厂会确保其创建的产品能相互搭配
// 使用。具体工厂方法签名会返回一个抽象产品，但在方法内部则会对具体产品进
// 行实例化。
type Adidas struct {
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

// Nike 具体工厂
// 每个具体工厂中都会包含一个相应的产品变体。
type Nike struct {
}

func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}

// IShoe 抽象产品
// 系列产品中的特定产品必须有一个基础接口。所有产品变体都必须实现这个接口。
type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetSize() int {
	return s.size
}

// AdidasShoe 具体产品
// 具体产品由相应的具体工厂创建。
type AdidasShoe struct {
	Shoe
}

// NikeShoe 具体产品
type NikeShoe struct {
	Shoe
}

// IShirt 抽象产品
// 这是另一个产品的基础接口。所有产品都可以互动，但是只有相同具体变体的产
// 品之间才能够正确地进行交互。
type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func (s *Shirt) GetSize() int {
	return s.size
}

// AdidasShirt 具体产品
type AdidasShirt struct {
	Shirt
}

// NikeShirt 具体产品
type NikeShirt struct {
	Shirt
}
