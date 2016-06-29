package document

type Image struct {
	path   string
	width  int
	height int
}

func (self *Image) Path() string {
	return self.path
}

func (self *Image) Width() int {
	return self.width
}

func (self *Image) Height() int {
	return self.height
}

func (self *Image) Resize(width, height int) {
	self.width = width
	self.height = height
}

func NewImage(path string, width, height int) *Image {
	image := &Image{}
	image.width = width
	image.height = height
	image.path = path
	return image
}
