package main

import "fmt"

// Fly behaviors
type FlyBehavior func()

func FlyWithWings() {
	fmt.Println("I believe I can fly")
}

func FlyNoWay() {}

func MakeFlyWithCount() func() {
	flyCount := 0
	return func() {
		fmt.Printf("I'm flying %v meter(s)\n", flyCount)
		flyCount++
	}
}

// Quack behaviors
type QuackBehavior func()

func Quack() {
	fmt.Println("Quack")
}

func Squeak() {
	fmt.Println("Squeak")
}

func MuteQuack() {}

// Dance behaviors
type DanceBehavior func()

func Waltz() {
	fmt.Println("I'm dancing a waltz")
}

func Minuet() {
	fmt.Println("I'm dancing a minuet")
}

func DanceNoWay() {}

//Ducks
type IDuck interface {
	Fly()
	Quack()
	Swim()
	Dance()
	Display()
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
	danceBehavior DanceBehavior
}

func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

func (d *Duck) SetQuackBehavior(qb QuackBehavior) {
	d.quackBehavior = qb
}

func (d *Duck) SetDanceBehavior(db DanceBehavior) {
	d.danceBehavior = db
}

func (d Duck) Fly() {
	if d.flyBehavior != nil {
		d.flyBehavior()
	}
}

func (d Duck) Swim() {
	fmt.Println("I'm swimming")
}

func (d Duck) Quack() {
	if d.quackBehavior != nil {
		d.quackBehavior()
	}
}

func (d Duck) Dance() {
	if d.danceBehavior != nil {
		d.danceBehavior()
	}
}

//Mallard duck
type MallardDuck struct {
	Duck
}

func (d MallardDuck) Display() {
	fmt.Println("I'm mallard duck")
}

func NewMallardDuck() *MallardDuck {
	d := &MallardDuck{}
	d.SetFlyBehavior(FlyWithWings)
	d.SetQuackBehavior(Quack)
	d.SetDanceBehavior(Waltz)
	return d
}

//Redhead duck
type RedheadDuck struct {
	Duck
}

func (d RedheadDuck) Display() {
	fmt.Println("I'm redhead duck")
}

func NewRedheadDuck() *RedheadDuck {
	d := &RedheadDuck{}
	d.SetFlyBehavior(MakeFlyWithCount())
	d.SetQuackBehavior(Squeak)
	d.SetDanceBehavior(Minuet)
	return d
}

//Rubber duck
type RubberDuck struct {
	Duck
}

func (d RubberDuck) Display() {
	fmt.Println("I'm rubber duck")
}

func NewRubberDuck() *RubberDuck {
	d := &RubberDuck{}
	d.SetFlyBehavior(FlyNoWay)
	d.SetQuackBehavior(Squeak)
	d.SetDanceBehavior(DanceNoWay)
	return d
}

//Decoy duck
type DecoyDuck struct {
	Duck
}

func (d DecoyDuck) Display() {
	fmt.Println("I'm decoy duck")
}

func NewDecoyDuck() *DecoyDuck {
	d := &DecoyDuck{}
	d.SetFlyBehavior(FlyNoWay)
	d.SetQuackBehavior(MuteQuack)
	d.SetDanceBehavior(DanceNoWay)
	return d
}

//Main client
func playWithDuck(d IDuck) {
	d.Display()
	d.Quack()
	d.Dance()
	d.Swim()
	d.Fly()
	d.Fly()
	d.Fly()
}

func main() {
	mallardDuck := NewMallardDuck()
	playWithDuck(mallardDuck)
	fmt.Println()

	redheadDuck := NewRedheadDuck()
	playWithDuck(redheadDuck)
	fmt.Println()

	rubberDuck := NewRubberDuck()
	playWithDuck(rubberDuck)
	fmt.Println()

	decoyDuck := NewDecoyDuck()
	playWithDuck(decoyDuck)
}
