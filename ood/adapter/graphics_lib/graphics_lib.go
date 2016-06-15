package graphics_lib

import "fmt"

type ICanvas interface {
	SetColor(rgbColor uint32)
	MoveTo(x, y int)
	LineTo(x, y int)
}

type Canvas struct{}

func (self *Canvas) SetColor(rgbColor uint32) {
	fmt.Printf("SetColor (#%06X)\n", rgbColor)
}

func (self *Canvas) MoveTo(x, y int) {
	fmt.Printf("MoveTo (%v, %v)\n", x, y)
}

func (self *Canvas) LineTo(x, y int) {
	fmt.Printf("LineTo (%v, %v)\n", x, y)
}

func NewCanvas() *Canvas {
	return &Canvas{}
}
