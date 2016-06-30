package document

import (
	"container/list"
	. "github.com/ilya-shikhaleev/codeKata/ood/command/command"
)

const END_POSITION = -1

type InsertDocumentItemCommand struct {
	items    *list.List
	position int
	item     *list.Element
}

func (self *InsertDocumentItemCommand) Execute() {
	InsertAtPosition(self.items, self.position, self.item)
}

func (self *InsertDocumentItemCommand) Unexecute() {
	item, _ := GetItemAtPosition(*self.items, self.position)
	self.items.Remove(item)
}

func NewInsertDocumentItemCommand(items *list.List, item DocumentItem, position int) Command {
	c := &InsertDocumentItemCommand{}
	c.items = items
	c.item = &list.Element{Value: item}
	c.position = position

	if c.position == END_POSITION {
		c.position = c.items.Len()
	}

	return &NoRepeatCommand{c, false}
}

type ReplaceDocumentItemCommand struct {
	items    *list.List
	position int
	item     *list.Element
	prev     *list.Element
}

func (self *ReplaceDocumentItemCommand) Execute() {
	self.prev, _ = GetItemAtPosition(*self.items, self.position)
	InsertAtPosition(self.items, self.position, self.item)
	self.items.Remove(self.prev)
}

func (self *ReplaceDocumentItemCommand) Unexecute() {
	item, _ := GetItemAtPosition(*self.items, self.position)
	InsertAtPosition(self.items, self.position, self.prev)
	self.items.Remove(item)
}

func NewReplaceDocumentItemCommand(items *list.List, item DocumentItem, position int) Command {
	c := &ReplaceDocumentItemCommand{}
	c.items = items
	c.item = &list.Element{Value: item}
	c.position = position

	if c.position == END_POSITION {
		c.position = c.items.Len()
	}

	return &NoRepeatCommand{c, false}
}
