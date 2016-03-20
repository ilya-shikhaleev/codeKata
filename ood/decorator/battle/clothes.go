package battle

type Clothes struct {
	CharacterDecorator
}

type Panties struct {
	Clothes
}

func (this Panties) Damage() float64 {
	return this.c.Damage() + 2
}

func AddPanties(c Character) *Panties {
	p := &Panties{}
	p.c = c
	return p
}

type Cap struct {
	Clothes
}

func (this Cap) Damage() float64 {
	return this.c.Damage() + 1
}

func AddCap(c Character) *Cap {
	cup := &Cap{}
	cup.c = c
	return cup
}

type Sneakers struct {
	Clothes
}

func (this Sneakers) Damage() float64 {
	return this.c.Damage() + 3
}

func AddSneakers(c Character) *Sneakers {
	s := &Sneakers{}
	s.c = c
	return s
}

type Slippers struct {
	Clothes
}

func (this Slippers) Damage() float64 {
	return this.c.Damage() + 0.5
}

func AddSlippers(c Character) *Slippers {
	s := &Slippers{}
	s.c = c
	return s
}

type Sword struct {
	Clothes
}

func (this Sword) Damage() float64 {
	return this.c.Damage() + 5
}

func AddSword(c Character) *Sword {
	s := &Sword{}
	s.c = c
	return s
}

type Shield struct {
	Clothes
}

func (this Shield) Damage() float64 {
	return this.c.Damage() + 5
}

func AddShield(c Character) *Shield {
	s := &Shield{}
	s.c = c
	return s
}
