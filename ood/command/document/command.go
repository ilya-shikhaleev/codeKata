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
	i := 0
	isInserted := false
	for currentItem := self.items.Front(); currentItem != nil; currentItem = currentItem.Next() {
		if i == self.position {
			self.items.InsertBefore(self.item, currentItem)
			isInserted = true
		}
		i++
	}
	if !isInserted {
		self.items.PushBack(self.item.Value)
	}
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
