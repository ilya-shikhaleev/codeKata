package battle

import "testing"

func TestClothesFight(t *testing.T) {
	var cases []CaseFight

	n := AddCap(AddSneakers(AddSword(NewNinja("Ninja", FEMALE))))
	n.LevelUp()

	a := AddShield(AddSlippers(NewAssassin("Assassin", MALE)))
	a.LevelUp()
	cases = append(cases, CaseFight{[2]Character{n, a}, true})

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

func TestClothesDamage(t *testing.T) {
	var cases []CasePlayer

	n := AddSneakers(NewNinja("Ninja with Sneakers", MALE))
	n.LevelDown()
	cases = append(cases, CasePlayer{n, WantedPlayer{4.5, "Ninja with Sneakers", MALE}})

	n2 := AddPanties(NewNinja("Ninja with Panties", FEMALE))
	n2.LevelUp()
	cases = append(cases, CasePlayer{n2, WantedPlayer{5, "Ninja with Panties", FEMALE}})

	var a Character = NewAssassin("Assassin with Cap", MALE)
	a.LevelUp()
	a = AddCap(a)
	cases = append(cases, CasePlayer{a, WantedPlayer{3.6, "Assassin with Cap", MALE}})

	w := AddSlippers(NewWizard("Wizard with Slippers", FEMALE))
	w.LevelUp()
	cases = append(cases, CasePlayer{w, WantedPlayer{4.5, "Wizard with Slippers", FEMALE}})

	n3 := AddSword(AddShield(NewNinja("Ninja with Sword and Shield", MALE)))
	cases = append(cases, CasePlayer{n3, WantedPlayer{11.5, "Ninja with Sword and Shield", MALE}})

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
