package command

type Executor interface {
	Execute()
}

type Command interface {
	Executor
	Unexecute()
}

type NoRepeatCommand struct {
	C          Command
	IsExecuted bool
}

func (self *NoRepeatCommand) Execute() {
	if !self.IsExecuted {
		self.C.Execute()
		self.IsExecuted = true
	}
}

func (self *NoRepeatCommand) Unexecute() {
	if self.IsExecuted {
		self.C.Unexecute()
		self.IsExecuted = false
	}
}

type ChangeStringCommand struct {
	target   *string
	newValue *string
}

func (self *ChangeStringCommand) Execute() {
	*self.newValue, *self.target = *self.target, *self.newValue
}

func (self *ChangeStringCommand) Unexecute() {
	*self.newValue, *self.target = *self.target, *self.newValue
}

func NewChangeStringCommand(target, newValue *string) Command {
	c := &ChangeStringCommand{}
	c.target = target
	c.newValue = newValue
	return &NoRepeatCommand{c, false}
}
