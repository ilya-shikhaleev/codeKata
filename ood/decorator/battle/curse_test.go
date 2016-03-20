package battle

import "testing"

func TestCurseFight(t *testing.T) {
	var cases []CaseFight

	n := ExecrateGenderCurse(NewNinja("Ninja", FEMALE))
	n.LevelUp()

	a := ExecrateLowDamageCurse(ExecrateLostLevelCurse(NewAssassin("Assassin", MALE)))
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

func TestCurseDamage(t *testing.T) {
	var cases []CasePlayer

	n := ExecrateLowDamageCurse(NewNinja("Ninja with LowDamageCurse", MALE))
	n.LevelDown()
	cases = append(cases, CasePlayer{n, WantedPlayer{1, "Ninja with LowDamageCurse", MALE}})

	n2 := ExecrateGenderCurse(NewNinja("Ninja with GenderCurse", FEMALE))
	n2.LevelUp()
	cases = append(cases, CasePlayer{n2, WantedPlayer{1.5, "Ninja with GenderCurse", FEMALE}})

	var a Character = NewAssassin("Assassin with LostLevelCurse", MALE)
	a.LevelUp()
	a = ExecrateLostLevelCurse(a)
	cases = append(cases, CasePlayer{a, WantedPlayer{1.3, "Assassin with LostLevelCurse", MALE}})

	n3 := ExecrateGenderCurse(ExecrateLowDamageCurse(NewNinja("Ninja with LowDamageCurse and GenderCurse", MALE)))
	cases = append(cases, CasePlayer{n3, WantedPlayer{1, "Ninja with LowDamageCurse and GenderCurse", MALE}})

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
