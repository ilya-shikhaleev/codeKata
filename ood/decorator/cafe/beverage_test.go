package cafe

import "testing"

type Wanted struct {
	cost        float64
	description string
}

type Case struct {
	in   Beverage
	want Wanted
}

func TestBeverage(t *testing.T) {
	cases := []Case{
		{
			NewCoffee(),
			Wanted{
				60.0,
				"Coffee",
			},
		},
		{
			NewLemon(NewCinnamon(NewTea()), 3),
			Wanted{
				80.0,
				"Tea, cinnamon, 3 lemon(s)",
			},
		},
		{
			NewIceCube(NewMilkshake(), 3, DRY_ICE_CUBE),
			Wanted{
				125.0,
				"Milkshake, 3 dry ice cube(s)",
			},
		},
		{
			NewCinnamon(NewCappuccino()),
			Wanted{
				100.0,
				"Cappuccino, cinnamon",
			},
		},
		{
			NewIceCube(NewLatte(), 2, WATER_ICE_CUBE),
			Wanted{
				100.0,
				"Latte, 2 water ice cube(s)",
			},
		},
	}
	for _, c := range cases {
		if c.in.GetCost() != c.want.cost || c.in.GetDescription() != c.want.description {
			t.Errorf("Beverage %s costs %v, expected %s costs %v", c.in.GetDescription(), c.in.GetCost(), c.want.description, c.want.cost)
		}
	}

}
