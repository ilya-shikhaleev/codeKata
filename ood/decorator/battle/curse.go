package battle

type Curse struct {
	PlayerDecorator
}

type GenderCurse struct {
	Curse
}

func (this GenderCurse) Damage() float64 {
	if this.p.Gender() == MALE {
		return this.p.Damage()
	}
	return this.p.Damage() / 2
}

func ExecrateGenderCurse(p Character) *GenderCurse {
	c := &GenderCurse{}
	c.p = p
	return c
}

type LowDamageCurse struct {
	Curse
}

func (this LowDamageCurse) Damage() float64 {
	return this.p.Damage() * 2 / 3
}

func ExecrateLowDamageCurse(p Character) *LowDamageCurse {
	c := &LowDamageCurse{}
	c.p = p
	return c
}

type LostLevelCurse struct {
	Curse
}

func (this LostLevelCurse) Damage() float64 {
	return this.p.Damage()
}

func ExecrateLostLevelCurse(p Character) *LostLevelCurse {
	c := &LostLevelCurse{}
	p.LevelDown()
	c.p = p
	return c
}
