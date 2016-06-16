package shape_drawing_lib

import (
	"github.com/ilya-shikhaleev/codeKata/ood/adapter/graphics_lib"
)

type Point struct {
	X, Y int
}

type ICanvasDrawable interface {
	Draw(canvas graphics_lib.ICanvas)
}

type Triangle struct {
	p1, p2, p3 *Point
	color      uint32
}

func (self *Triangle) Draw(canvas graphics_lib.ICanvas) {
	canvas.SetColor(self.color)
	canvas.MoveTo(self.p1.X, self.p1.Y)
	canvas.LineTo(self.p2.X, self.p2.Y)
	canvas.LineTo(self.p3.X, self.p3.Y)
	canvas.LineTo(self.p1.X, self.p1.Y)
}

func NewTriangle(p1, p2, p3 *Point, color uint32) *Triangle {
	return &Triangle{p1, p2, p3, color}
}

type Rectangle struct {
	leftTop *Point
	width   int
	height  int
	color   uint32
}

func (self *Rectangle) Draw(canvas graphics_lib.ICanvas) {
	canvas.SetColor(self.color)
	canvas.MoveTo(self.leftTop.X, self.leftTop.Y)
	canvas.LineTo(self.leftTop.X+self.width, self.leftTop.Y)
	canvas.LineTo(self.leftTop.X+self.width, self.leftTop.Y+self.height)
	canvas.LineTo(self.leftTop.X, self.leftTop.Y+self.height)
	canvas.LineTo(self.leftTop.X, self.leftTop.Y)
}

func NewRectangle(leftTop *Point, width, height int, color uint32) *Rectangle {
	return &Rectangle{leftTop, width, height, color}
}

type CanvasPainter struct {
	canvas graphics_lib.ICanvas
}

func (self *CanvasPainter) Draw(drawable ICanvasDrawable) {
	drawable.Draw(self.canvas)
}

func NewCanvasPainter(canvas graphics_lib.ICanvas) *CanvasPainter {
	return &CanvasPainter{canvas}
}
