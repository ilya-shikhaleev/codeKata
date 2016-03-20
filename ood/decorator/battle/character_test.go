package battle

import (
	"testing"
)

func TestPlayerFight(t *testing.T) {
	var cases []CaseFight

	n := NewNinja("Ninja", FEMALE)
	n.LevelUp()

	a := NewAssassin("Assassin", MALE)
	a.LevelUp()
	a.LevelUp()
	cases = append(cases, CaseFight{[2]Character{n, a}, false})

	for _, c := range cases {
		p1 := c.in[0]
		p2 := c.in[1]
		fight := p1.Fight(p1, p2)
		if fight != c.want {
			t.Errorf("Fight between %v(dmg=%.1f) and %v(dmg=%.1f) ended with %v, expected %v",
				p1.Name(), p1.Damage(), p2.Name(), p2.Damage(), fight, c.want)
		}
	}
}

func TestPlayerDamage(t *testing.T) {
	var cases []CasePlayer

	n := NewNinja("Ninja", MALE)
	n.LevelDown()
	cases = append(cases, CasePlayer{n, WantedPlayer{1.5, "Ninja", MALE}})

	n2 := NewNinja("Ninja2", FEMALE)
	n2.LevelUp()
	cases = append(cases, CasePlayer{n2, WantedPlayer{3, "Ninja2", FEMALE}})

	a := NewAssassin("Assassin", MALE)
	a.level = 10
	a.LevelUp()
	cases = append(cases, CasePlayer{a, WantedPlayer{13, "Assassin", MALE}})

	w := NewWizard("Wizard", FEMALE)
	w.level = 5
	w.LevelDown()
	cases = append(cases, CasePlayer{w, WantedPlayer{8, "Wizard", FEMALE}})

	for _, c := range cases {
		if c.in.Damage() != c.want.damage {
			t.Errorf("Hero %s has damage %v, expected %v", c.in.Name(), c.in.Damage(), c.want.damage)
		}
		if c.in.Name() != c.want.name {
			t.Errorf("Hero %s has name %v, expected %v", c.in.Name(), c.in.Name(), c.want.name)
		}
		if c.in.Gender() != c.want.gender {
			t.Errorf("Hero %s has gender %v, expected %v", c.in.Name(), c.in.Gender(), c.want.gender)
		}
	}
}
