package modern_graphics_lib

import (
	"bytes"
	"fmt"
	"io"
)

type RGBAColor struct {
	r, g, b, a float32
}

type Point struct {
	x, y int
}

type ModernGraphicsRenderer struct {
	out     io.Writer
	drawing bool
}

func (self *ModernGraphicsRenderer) BeginDraw() error {
	if self.drawing {
		return fmt.Errorf("Drawing has already begun")
	}
	buffer := bytes.NewBufferString("<draw>\n")
	buffer.WriteTo(self.out)
	self.drawing = true
	return nil
}

func (self *ModernGraphicsRenderer) DrawLine(start, end Point, color RGBAColor) error {
	if !self.drawing {
		return fmt.Errorf("DrawLine is allowed between BeginDraw()/EndDraw() only")
	}
	str := fmt.Sprintf(
		`<line fromX="%v" fromY="%v" toX="%v" toY="%v">\n
  <color r="%.2f" g="%.2f" b="%.2f" a="%.2f"
</line>%v`, start.x, start.y, end.x, end.y, color.r, color.g, color.b, color.a, "\n")
	buffer := bytes.NewBufferString(str)
	buffer.WriteTo(self.out)
	self.drawing = true
	return nil
}

func (self *ModernGraphicsRenderer) EndDraw() error {
	if !self.drawing {
		return fmt.Errorf("Drawing has not been started")
	}
	buffer := bytes.NewBufferString("</draw>\n")
	buffer.WriteTo(self.out)
	self.drawing = false
	return nil
}

func (self *ModernGraphicsRenderer) Close() {
	if self.drawing {
		if err := self.EndDraw(); err != nil {
			panic(err)
		}
	}
}

func NewModernGraphicsRenderer(out io.Writer) *ModernGraphicsRenderer {
	return &ModernGraphicsRenderer{out, false}
}
