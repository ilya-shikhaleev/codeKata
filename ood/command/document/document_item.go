package document

type DocumentItem struct {
	image     *Image
	paragraph *Paragraph
}

func (self *DocumentItem) Paragraph() *Paragraph {
	return self.paragraph
}

func (self *DocumentItem) Image() *Image {
	return self.image
}
