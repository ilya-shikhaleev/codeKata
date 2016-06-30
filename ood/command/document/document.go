package document

import (
	"container/list"
	"fmt"
	"github.com/ilya-shikhaleev/codeKata/ood/command/command"
)

type Document struct {
	title   string
	history History
	items   *list.List
}

func (self *Document) InsertParagraph(text string, position int) {
	item := DocumentItem{
		image: nil,
		paragraph: &Paragraph{
			text: text,
		},
	}
	command := NewInsertDocumentItemCommand(self.items, item, position)
	self.history.AddAndExecuteCommand(command)
}

func (self *Document) ReplaceParagraph(text string, position int) {
	item := DocumentItem{
		image: nil,
		paragraph: &Paragraph{
			text: text,
		},
	}
	command := NewReplaceDocumentItemCommand(self.items, item, position)
	self.history.AddAndExecuteCommand(command)
}

func (self *Document) InsertImage(path string, width, height int) {
}

func (self *Document) ItemsCount() int {
	return self.items.Len()
}

func (self *Document) Item(position int) (DocumentItem, error) {
	item, err := GetItemAtPosition(*self.items, position)
	return item.Value.(DocumentItem), err
}

func (self *Document) Items() list.List {
	return *self.items
}

func (self *Document) DeleteItem() {
}

func (self *Document) Title() string {
	return self.title
}

func (self *Document) SetTitle(title string) {
	self.history.AddAndExecuteCommand(command.NewChangeStringCommand(&self.title, &title))
}
func (self *Document) CanUndo() bool {
	return self.history.CanUndo()
}

func (self *Document) Undo() {
	self.history.Undo()
}

func (self *Document) CanRedo() bool {
	return self.history.CanRedo()
}

func (self *Document) Redo() {
	self.history.Redo()
}

func (self *Document) Save() {
}

func NewDocument() *Document {
	d := &Document{}
	d.items = list.New()
	return d
}

func GetItemAtPosition(items list.List, position int) (*list.Element, error) {
	i := 0
	for currentItem := items.Front(); currentItem != nil; currentItem = currentItem.Next() {
		if i == position {
			return currentItem, nil
		}
		i++
	}
	return nil, fmt.Errorf("Not found item")
}

func InsertAtPosition(items *list.List, position int, item *list.Element) {
	i := 0
	isInserted := false
	for currentItem := items.Front(); currentItem != nil; currentItem = currentItem.Next() {
		if i == position {
			items.InsertBefore(item.Value, currentItem)
			isInserted = true
		}
		i++
	}
	if !isInserted {
		items.PushBack(item.Value)
	}
}
