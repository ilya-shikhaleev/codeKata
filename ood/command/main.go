package main

import (
	"bufio"
	"fmt"
	"github.com/ilya-shikhaleev/codeKata/ood/command/document"
	"io"
	"os"
	"strings"
)

type MenuCommand func(input string)

type Menu struct {
	exit  bool
	items map[string]Item
}

type Item struct {
	shortcut    string
	description string
	command     MenuCommand
}

func (self *Menu) AddItem(shortcut, description string, command MenuCommand) {
	self.items[shortcut] = Item{shortcut, description, command}
}

func (self *Menu) Run(input *bufio.Reader) {
	for {
		s, isPrefix, err := input.ReadLine()
		if err == io.EOF {
			break
		}
		if isPrefix {
			fmt.Println("Command is too long, try again")
			continue
		}
		if !self.executeCommand(string(s)) {
			break
		}
	}
}

func (self *Menu) executeCommand(line string) bool {
	self.exit = false
	words := strings.Fields(line)
	if len(words) == 0 {
		fmt.Println("Empty command")
	} else {
		item, ok := self.items[words[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			item.command(strings.Trim(line[len(words[0]):], " "))
		}
	}
	return !self.exit
}

func (self *Menu) ShowInstructions() {
	fmt.Println("Commands list:")
	for _, item := range self.items {
		fmt.Printf("\t%v: %v\n", item.shortcut, item.description)
	}
}

func NewMenu() *Menu {
	m := &Menu{}
	m.items = make(map[string]Item)
	return m
}

func (self *Menu) Exit() {
	self.exit = true
}

type MenuHelpCommand struct {
	menu *Menu
}

func (c *MenuHelpCommand) Execute(input string) {
	c.menu.ShowInstructions()
}

type ExitMenuCommand struct {
	menu *Menu
}

func (c *ExitMenuCommand) Execute(input string) {
	c.menu.Exit()
}

type ShowListCommand struct {
	d *document.Document
}

func (c *ShowListCommand) Execute(input string) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(c.d.Title())
	fmt.Println(strings.Repeat("-", 20))
}

type UndoCommand struct {
	d *document.Document
}

func (c *UndoCommand) Execute(input string) {
	if c.d.CanUndo() {
		c.d.Undo()
	} else {
		fmt.Println("Can't undo")
	}
}

type RedoCommand struct {
	d *document.Document
}

func (c *RedoCommand) Execute(input string) {
	if c.d.CanRedo() {
		c.d.Redo()
	} else {
		fmt.Println("Can't redo")
	}
}

type SetTitleCommand struct {
	d *document.Document
}

func (c *SetTitleCommand) Execute(input string) {
	c.d.SetTitle(input)
}

func main() {
	d := document.NewDocument()
	menu := NewMenu()

	menuHelpCommand := &MenuHelpCommand{menu}
	menu.AddItem("help", "Show instructions", menuHelpCommand.Execute)
	exitMenuCommand := &ExitMenuCommand{menu}
	menu.AddItem("exit", "Exit from this menu", exitMenuCommand.Execute)
	showListCommand := &ShowListCommand{d}
	menu.AddItem("list", "Show document", showListCommand.Execute)
	undoCommand := &UndoCommand{d}
	menu.AddItem("undo", "Undo command", undoCommand.Execute)
	redoCommand := &RedoCommand{d}
	menu.AddItem("redo", "Redo undone command", redoCommand.Execute)
	setTitleCommand := &SetTitleCommand{d}
	menu.AddItem("setTitle", "Changes title. Args: <new title>", setTitleCommand.Execute)

	menu.executeCommand("help")
	menu.Run(bufio.NewReader(os.Stdin))
}
