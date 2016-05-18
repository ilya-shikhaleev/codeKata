package command

type Executor interface {
	Execute()
}

type Command interface {
	Executor
	Unexecute()
}

type NoRepeatCommand struct {
	c          Command
	isExecuted bool
}

func (self *NoRepeatCommand) Execute() {
	if !self.isExecuted {
		self.c.Execute()
		self.isExecuted = true
	}
}

func (self *NoRepeatCommand) Unexecute() {
	if self.isExecuted {
		self.c.Unexecute()
		self.isExecuted = false
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
