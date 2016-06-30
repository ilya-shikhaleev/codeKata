package main

import (
	"bufio"
	"fmt"
	"github.com/ilya-shikhaleev/codeKata/ood/command/document"
	"io"
	"os"
	"strconv"
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

	items := c.d.Items()
	for currentItem := items.Front(); currentItem != nil; currentItem = currentItem.Next() {
		item := currentItem.Value.(document.DocumentItem)
		if item.Image() != nil {
			fmt.Println("Image")
		}
		if item.Paragraph() != nil {
			fmt.Println("<p>", item.Paragraph().Text(), "</p>")
		}
	}

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

type InsertParagraphCommand struct {
	d *document.Document
}

func (c *InsertParagraphCommand) Execute(input string) {
	params := strings.SplitN(input, " ", 2)
	if len(params) != 2 {
		fmt.Println("Wrong command")
		return
	}
	position, err := preparePosition(params[0], c.d, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.d.InsertParagraph(params[1], position)
}

type ReplaceParagraphCommand struct {
	d *document.Document
}

func (c *ReplaceParagraphCommand) Execute(input string) {
	params := strings.SplitN(input, " ", 2)
	if len(params) != 2 {
		fmt.Println("Wrong command")
		return
	}
	position, err := preparePosition(params[0], c.d, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.d.ReplaceParagraph(params[1], position)
}

func preparePosition(param string, d *document.Document, canExtend bool) (int, error) {

	var position int
	if param == "end" {
		position = document.END_POSITION
	} else {
		i, err := strconv.Atoi(param)
		if err != nil || i < 0 || i > d.ItemsCount() || (!canExtend && i == d.ItemsCount()) {
			return 0, fmt.Errorf("Bad position")
		}
		position = i
	}
	return position, nil
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
	menu.AddItem("setTitle", "Changes title. Args: <заголовок документа>", setTitleCommand.Execute)
	insertParagraphCommand := &InsertParagraphCommand{d}
	menu.AddItem("insertParagraph", "insertParagraph <позиция>|end <текст параграфа>", insertParagraphCommand.Execute)
	replaceParagraphCommand := &ReplaceParagraphCommand{d}
	menu.AddItem("replaceText", "replaceText <позиция>|end <текст параграфа>", replaceParagraphCommand.Execute)

	menu.executeCommand("help")
	menu.Run(bufio.NewReader(os.Stdin))
}
