package battle

type Clothes struct {
	PlayerDecorator
}

type Panties struct {
	Clothes
}

func (this Panties) Damage() float64 {
	return this.p.Damage() + 2
}

func AddPanties(p Character) *Panties {
	c := &Panties{}
	c.p = p
	return c
}

type Cap struct {
	Clothes
}

func (this Cap) Damage() float64 {
	return this.p.Damage() + 1
}

func AddCap(p Character) *Cap {
	c := &Cap{}
	c.p = p
	return c
}

type Sneakers struct {
	Clothes
}

func (this Sneakers) Damage() float64 {
	return this.p.Damage() + 3
}

func AddSneakers(p Character) *Sneakers {
	c := &Sneakers{}
	c.p = p
	return c
}

type Slippers struct {
	Clothes
}

func (this Slippers) Damage() float64 {
	return this.p.Damage() + 0.5
}

func AddSlippers(p Character) *Slippers {
	c := &Slippers{}
	c.p = p
	return c
}

type Sword struct {
	Clothes
}

func (this Sword) Damage() float64 {
	return this.p.Damage() + 5
}

func AddSword(p Character) *Sword {
	c := &Sword{}
	c.p = p
	return c
}

type Shield struct {
	Clothes
}

func (this Shield) Damage() float64 {
	return this.p.Damage() + 5
}

func AddShield(p Character) *Shield {
	c := &Shield{}
	c.p = p
	return c
}
