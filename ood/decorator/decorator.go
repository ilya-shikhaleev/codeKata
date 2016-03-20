package main

import (
	"fmt"
	"github.com/ilya-shikhaleev/codeKata/ood/decorator/cafe"
)

func displayBeverage(beverage cafe.Beverage) {
	fmt.Printf("Your %s costs %v\n", beverage.GetDescription(), beverage.GetCost())
}

func main() {
	coffee := cafe.NewCoffee()
	displayBeverage(coffee)

	cappuccino := cafe.NewCappuccino()
	cinnamonCappuccino := cafe.NewCinnamon(cappuccino)
	displayBeverage(cinnamonCappuccino)

	displayBeverage(cafe.NewLemon(cafe.NewCinnamon(cafe.NewTea()), 3))
	displayBeverage(cafe.NewIceCube(cafe.NewMilkshake(), 3, cafe.DRY_ICE_CUBE))
}
