package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Executor interface {
	Execute()
}

type MacroCommand struct {
	commands []Executor
}

func (c *MacroCommand) Execute() {
	for _, command := range c.commands {
		command.Execute()
	}
}

type HorseMovingCommand struct {
	MacroCommand
}

func NewHorseMovingCommand(robot *Robot) *HorseMovingCommand {
	c := &HorseMovingCommand{}
	c.commands = []Executor{&WalkCommand{robot, RIGHT}, &WalkCommand{robot, RIGHT}, &WalkCommand{robot, UP}}
	return c
}

type TurnOnCommand struct {
	robot *Robot
}

func (c *TurnOnCommand) Execute() {
	c.robot.TurnOn()
}

type TurnOffCommand struct {
	robot *Robot
}

func (c *TurnOffCommand) Execute() {
	c.robot.TurnOff()
}

type WalkCommand struct {
	robot     *Robot
	direction int64
}

func (c *WalkCommand) Execute() {
	c.robot.Walk(c.direction)
}

type StopCommand struct {
	robot *Robot
}

func (c *StopCommand) Execute() {
	c.robot.Stop()
}

type MenuHelpCommand struct {
	menu *Menu
}

func (c *MenuHelpCommand) Execute() {
	c.menu.ShowInstructions()
}

type ExitMenuCommand struct {
	menu *Menu
}

func (c *ExitMenuCommand) Execute() {
	c.menu.Exit()
}

const (
	UP int64 = iota
	DOWN
	LEFT
	RIGHT
	NO_DIRECTION
)

type Robot struct {
	direction int64
	turnedOn  bool
}

func (r *Robot) TurnOn() {
	if !r.turnedOn {
		r.turnedOn = true
		fmt.Println("It am waiting for your commands")
	}
}

func (r *Robot) TurnOff() {
	if r.turnedOn {
		r.turnedOn = false
		r.direction = NO_DIRECTION
		fmt.Println("It is a pleasure to serve you")
	}
}

func (r *Robot) Walk(direction int64) {
	if r.turnedOn {
		r.direction = direction
		directions := make(map[int64]string)
		directions[UP] = "up"
		directions[DOWN] = "down"
		directions[LEFT] = "left"
		directions[RIGHT] = "right"
		fmt.Printf("Walking %v\n", directions[direction])
	} else {
		fmt.Println("The robot should be turned on first")
	}
}

func (r *Robot) Stop() {
	if r.turnedOn {
		if r.direction != NO_DIRECTION {
			r.direction = NO_DIRECTION
			fmt.Printf("Stopped\n")
		} else {
			fmt.Printf("I am staying still\n")
		}
	} else {
		fmt.Println("The robot should be turned on first")
	}
}

func NewRobot() *Robot {
	return &Robot{NO_DIRECTION, false}
}

type Menu struct {
	exit  bool
	items map[string]Item
}

type Item struct {
	shortcut    string
	description string
	command     Executor
}

func (m *Menu) AddItem(shortcut, description string, command Executor) {
	m.items[shortcut] = Item{shortcut, description, command}
}

func (m *Menu) Run(input *bufio.Reader) {
	for {
		s, isPrefix, err := input.ReadLine()
		if err == io.EOF {
			break
		}
		if isPrefix {
			fmt.Println("Command is too long, try again")
			continue
		}
		if !m.executeCommand(string(s)) {
			break
		}
	}
}

func (m *Menu) executeCommand(word string) bool {
	m.exit = false
	item, ok := m.items[word]
	if !ok {
		fmt.Println("Unknown command")
	} else {
		item.command.Execute()
	}
	return !m.exit
}

func (m *Menu) ShowInstructions() {
	fmt.Println("Commands list:")
	for _, item := range m.items {
		fmt.Printf("\t%v: %v\n", item.shortcut, item.description)
	}
}

func NewMenu() *Menu {
	m := &Menu{}
	m.items = make(map[string]Item)
	return m
}

func (m *Menu) Exit() {
	m.exit = true
}

func main() {
	robot := NewRobot()
	menu := NewMenu()

	menu.AddItem("on", "Turns the Robot on", &TurnOnCommand{robot})
	menu.AddItem("off", "Turns the Robot off", &TurnOffCommand{robot})

	menu.AddItem("up", "Makes the Robot walk up", &WalkCommand{robot, UP})
	menu.AddItem("down", "Makes the Robot walk down", &WalkCommand{robot, DOWN})
	menu.AddItem("left", "Makes the Robot walk left", &WalkCommand{robot, LEFT})
	menu.AddItem("right", "Makes the Robot walk right", &WalkCommand{robot, RIGHT})
	menu.AddItem("horse_moving", "Makes the Robot walk like horse", NewHorseMovingCommand(robot))

	menu.AddItem("stop", "Stops the Robot", &StopCommand{robot})

	menu.AddItem("help", "Show instructions", &MenuHelpCommand{menu})
	menu.AddItem("exit", "Exit from this menu", &ExitMenuCommand{menu})

	menu.executeCommand("help")
	menu.Run(bufio.NewReader(os.Stdin))
}
