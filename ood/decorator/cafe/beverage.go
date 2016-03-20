package cafe

import "fmt"

type Beverage interface {
	GetDescription() string
	GetCost() float64
}

type BeverageDescription struct {
	description string
}

func (this BeverageDescription) GetDescription() string {
	return this.description
}

type Coffee struct {
	BeverageDescription
}

func (this Coffee) GetCost() float64 {
	return 60.0
}

func NewCoffee() *Coffee {
	coffe := &Coffee{}
	coffe.description = "Coffee"
	return coffe
}

type Cappuccino struct {
	Coffee
}

func (this Cappuccino) GetCost() float64 {
	return 80.0
}

func NewCappuccino() *Cappuccino {
	cappuccino := &Cappuccino{}
	cappuccino.description = "Cappuccino"
	return cappuccino
}

type Latte struct {
	Coffee
}

func (this Latte) GetCost() float64 {
	return 90.0
}

func NewLatte() *Latte {
	latte := &Latte{}
	latte.description = "Latte"
	return latte
}

type Tea struct {
	BeverageDescription
}

func (this Tea) GetCost() float64 {
	return 30.0
}

func NewTea() *Tea {
	tea := &Tea{}
	tea.description = "Tea"
	return tea
}

type Milkshake struct {
	BeverageDescription
}

func (this Milkshake) GetCost() float64 {
	return 80.0
}

func NewMilkshake() *Milkshake {
	milkshake := &Milkshake{}
	milkshake.description = "Milkshake"
	return milkshake
}

//Decorators
type Cinnamon struct {
	beverage Beverage
}

func (this Cinnamon) GetCost() float64 {
	return this.beverage.GetCost() + 20.0
}

func (this Cinnamon) GetDescription() string {
	return fmt.Sprintf("%v, cinnamon", this.beverage.GetDescription())
}

func NewCinnamon(beverage Beverage) *Cinnamon {
	cinnamon := &Cinnamon{beverage}
	return cinnamon
}

type Lemon struct {
	beverage Beverage
	count    uint
}

func (this Lemon) GetCost() float64 {
	lemonCost := 10.0
	return this.beverage.GetCost() + lemonCost*float64(this.count)
}

func (this Lemon) GetDescription() string {
	return fmt.Sprintf("%v, %v lemon(s)", this.beverage.GetDescription(), this.count)
}

func NewLemon(beverage Beverage, count uint) *Lemon {
	lemon := &Lemon{beverage, count}
	return lemon
}

type IceCubeType int

const (
	WATER_ICE_CUBE IceCubeType = iota
	DRY_ICE_CUBE
)

type IceCube struct {
	beverage Beverage
	count    uint
	iceType  IceCubeType
}

func (this IceCube) GetCost() float64 {
	var cost float64
	switch this.iceType {
	case WATER_ICE_CUBE:
		cost = 5.0
	case DRY_ICE_CUBE:
		cost = 15.0
	}
	return this.beverage.GetCost() + cost*float64(this.count)
}

func (this IceCube) GetDescription() string {
	var iceType string
	switch this.iceType {
	case WATER_ICE_CUBE:
		iceType = "water"
	case DRY_ICE_CUBE:
		iceType = "dry"
	}
	return fmt.Sprintf("%v, %v %s ice cube(s)", this.beverage.GetDescription(), this.count, iceType)
}

func NewIceCube(beverage Beverage, count uint, iceType IceCubeType) *IceCube {
	iceCube := &IceCube{beverage, count, iceType}
	return iceCube
}
