package document

type Paragraph struct {
	text string
}

func (self *Paragraph) SetText(text string) {
	self.text = text
}

func (self *Paragraph) Text() string {
	return self.text
}
