package main

import (
	"fmt"
	"github.com/ilya-shikhaleev/codeKata/ood/decorator/battle"
	"strings"
)

func fight(first, second battle.Character) {
	fightResult := "lose"
	if first.Fight(first, second) {
		fightResult = "win"
	}
	fmt.Printf("Player %s(dmg=%.1f) %v player %s(dmg=%.1f)\n",
		first.Name(), first.Damage(), fightResult, second.Name(), second.Damage())
	fmt.Println(strings.Repeat("-", 40))
}

func main() {
	var firstPlayer battle.Character = battle.NewAssassin("Winter", battle.MALE)
	var secondPlayer battle.Character = battle.NewWizard("Summer", battle.FEMALE)
	fight(firstPlayer, secondPlayer)

	fmt.Printf("%s dresses sword, slippers and cap\n", firstPlayer.Name())
	firstPlayer = battle.AddSword(firstPlayer)
	firstPlayer = battle.AddCap(firstPlayer)
	firstPlayer = battle.AddSlippers(firstPlayer)
	fmt.Printf("%s dresses panties, sneakers and shield\n", secondPlayer.Name())
	secondPlayer = battle.AddPanties(secondPlayer)
	secondPlayer = battle.AddSneakers(secondPlayer)
	secondPlayer = battle.AddShield(secondPlayer)
	fight(firstPlayer, secondPlayer)

	fmt.Printf("%s level uped twice and execrated low damage curse\n", firstPlayer.Name())
	firstPlayer.LevelUp()
	firstPlayer.LevelUp()
	firstPlayer = battle.ExecrateLowDamageCurse(firstPlayer)
	fmt.Printf("%s level uped and execrated female curse\n", secondPlayer.Name())
	secondPlayer = battle.ExecrateGenderCurse(secondPlayer)
	secondPlayer.LevelUp()
	fight(firstPlayer, secondPlayer)

	fmt.Printf("%s level uped and execrated lost level curse\n", firstPlayer.Name())
	firstPlayer.LevelUp()
	firstPlayer = battle.ExecrateLostLevelCurse(firstPlayer)
	fmt.Printf("%s level uped and execrated low damage curse\n", secondPlayer.Name())
	secondPlayer = battle.ExecrateLowDamageCurse(secondPlayer)
	secondPlayer.LevelUp()
	fight(firstPlayer, secondPlayer)
}
