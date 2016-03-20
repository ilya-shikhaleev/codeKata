package battle

type Damager interface {
	Damage() float64
}

type Fighter interface {
	Fight(hero, enemy Damager) bool
}

type Character interface {
	Fighter
	Damager
	LevelUp()
	LevelDown()
	Name() string
	Gender() Gender
}

type Gender uint8

const (
	MALE Gender = iota
	FEMALE
)

const MAX_LEVEL = 10
const MIN_LEVEL = 1

type Hero struct {
	Character
	name   string
	gender Gender
	level  uint8
}

func (this Hero) Fight(hero, ememy Damager) bool {
	return hero.Damage() > ememy.Damage()
}

func (this *Hero) LevelUp() {
	if this.level < MAX_LEVEL {
		this.level++
	}
}

func (this *Hero) LevelDown() {
	if this.level > MIN_LEVEL {
		this.level--
	}
}

func (this Hero) Name() string {
	return this.name
}

func (this Hero) Gender() Gender {
	return this.gender
}

type Ninja struct {
	Hero
}

func (this Ninja) Damage() float64 {
	return float64(this.level) * 1.5
}

func NewNinja(name string, g Gender) *Ninja {
	n := &Ninja{}
	n.name = name
	n.gender = g
	n.level = 1
	return n
}

type Assassin struct {
	Hero
}

func (this Assassin) Damage() float64 {
	return float64(this.level) * 1.3
}

func NewAssassin(name string, g Gender) *Assassin {
	a := &Assassin{}
	a.name = name
	a.gender = g
	a.level = 1
	return a
}

type Wizard struct {
	Hero
}

func (this Wizard) Damage() float64 {
	return float64(this.level) * 2
}

func NewWizard(name string, g Gender) *Wizard {
	a := &Wizard{}
	a.name = name
	a.gender = g
	a.level = 1
	return a
}

type CharacterDecorator struct {
	c Character
}

func (this CharacterDecorator) Fight(hero, monster Damager) bool {
	return this.c.Fight(hero, monster)
}

func (this *CharacterDecorator) LevelUp() {
	this.c.LevelUp()
}

func (this *CharacterDecorator) LevelDown() {
	this.c.LevelDown()
}

func (this CharacterDecorator) Name() string {
	return this.c.Name()
}

func (this CharacterDecorator) Gender() Gender {
	return this.c.Gender()
}
