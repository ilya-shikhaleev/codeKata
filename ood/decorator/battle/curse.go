package battle

type Curse struct {
	CharacterDecorator
}

type GenderCurse struct {
	Curse
}

func (this GenderCurse) Damage() float64 {
	if this.c.Gender() == MALE {
		return this.c.Damage()
	}
	return this.c.Damage() / 2
}

func ExecrateGenderCurse(c Character) *GenderCurse {
	curse := &GenderCurse{}
	curse.c = c
	return curse
}

type LowDamageCurse struct {
	Curse
}

func (this LowDamageCurse) Damage() float64 {
	return this.c.Damage() * 2 / 3
}

func ExecrateLowDamageCurse(c Character) *LowDamageCurse {
	curse := &LowDamageCurse{}
	curse.c = c
	return curse
}

type LostLevelCurse struct {
	Curse
}

func (this LostLevelCurse) Damage() float64 {
	return this.c.Damage()
}

func ExecrateLostLevelCurse(c Character) *LostLevelCurse {
	curse := &LostLevelCurse{}
	c.LevelDown()
	curse.c = c
	return curse
}
