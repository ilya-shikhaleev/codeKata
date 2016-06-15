package modern_graphics_lib

import (
	"fmt"
)

type ModernGraphicsRendererAdapter struct {
	renderer *ModernGraphicsRenderer
	start    Point
	color    RGBAColor
}

func (self *ModernGraphicsRendererAdapter) SetColor(color uint32) {
	self.color.r = float32((color>>16)%0X100) / float32(0XFF)
	self.color.g = float32((color>>8)%0X100) / float32(0XFF)
	self.color.b = float32((color)%0X100) / float32(0XFF)
	self.color.a = 1.0
}

func (self *ModernGraphicsRendererAdapter) MoveTo(x, y int) {
	self.start.x = x
	self.start.y = y
}

func (self *ModernGraphicsRendererAdapter) LineTo(x, y int) {
	if err := self.renderer.DrawLine(self.start, Point{x, y}, self.color); err != nil {
		fmt.Println(err)
	}
}

func (self *ModernGraphicsRendererAdapter) Close() {
	self.renderer.Close()
}

func NewModernGraphicsRendererAdapter(renderer *ModernGraphicsRenderer) *ModernGraphicsRendererAdapter {
	adapter := &ModernGraphicsRendererAdapter{}
	adapter.renderer = renderer
	if err := renderer.BeginDraw(); err != nil {
		fmt.Println(err)
	}
	return adapter
}
