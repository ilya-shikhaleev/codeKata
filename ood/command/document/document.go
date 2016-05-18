package document

import "github.com/ilya-shikhaleev/codeKata/ood/command/command"

type Document struct {
	title   string
	history History
}

func (self *Document) InsertParagraph(text string, position int) {
}

func (self *Document) InsertImage(path string, width, height int) {
}

func (self *Document) ItemsCount() int {
	return 0
}

func (self *Document) Item() {
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
	return &Document{}
}
