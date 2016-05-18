package document

import (
	"github.com/ilya-shikhaleev/codeKata/ood/command/command"
)

type History struct {
	nextCommandIndex int
	commands         []command.Command
}

func (self *History) CanUndo() bool {
	return self.nextCommandIndex != 0
}

func (self *History) Undo() {
	if self.CanUndo() {
		self.commands[self.nextCommandIndex-1].Unexecute()
		self.nextCommandIndex--
	}
}

func (self *History) CanRedo() bool {
	return self.nextCommandIndex != len(self.commands)
}

func (self *History) Redo() {
	if self.CanRedo() {
		self.commands[self.nextCommandIndex].Execute()
		self.nextCommandIndex++
	}
}

func (self *History) AddAndExecuteCommand(c command.Command) {
	c.Execute()
	if self.nextCommandIndex < len(self.commands) {
		self.commands = self.commands[:self.nextCommandIndex]
		//Commands will be removed, add removing resources
	}
	self.commands = append(self.commands, c)
	self.nextCommandIndex++
}
