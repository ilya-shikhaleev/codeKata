package main

import (
	"github.com/ilya-shikhaleev/codeKata/ood/adapter/graphics_lib"
	"github.com/ilya-shikhaleev/codeKata/ood/adapter/modern_graphics_lib"
	. "github.com/ilya-shikhaleev/codeKata/ood/adapter/shape_drawing_lib"
	"os"
)

func paintPicture(painter *CanvasPainter) {
	triangle := NewTriangle(
		&Point{10, 15},
		&Point{100, 200},
		&Point{150, 250},
		0,
	)

	rectangle := NewRectangle(
		&Point{30, 40},
		18,
		24,
		123456,
	)

	painter.Draw(triangle)
	painter.Draw(rectangle)
}

func paintPictureOnCanvas() {
	simpleCanvas := graphics_lib.NewCanvas()
	painter := NewCanvasPainter(simpleCanvas)
	paintPicture(painter)
}

func paintPictureOnModernGraphicsRenderer() {
	renderer := modern_graphics_lib.NewModernGraphicsRenderer(os.Stdout)
	canvasAdapter := modern_graphics_lib.NewModernGraphicsRendererAdapter(renderer)
	defer canvasAdapter.Close()

	painter := NewCanvasPainter(canvasAdapter)
	paintPicture(painter)
}

func paintPictureOnModernGraphicsEmbeddedRenderer() {
	canvasAdapter := modern_graphics_lib.NewModernGraphicsRendererEmbeddedAdapter(os.Stdout)
	defer canvasAdapter.Close()

	painter := NewCanvasPainter(canvasAdapter)
	paintPicture(painter)
}

func main() {
	paintPictureOnCanvas()
	paintPictureOnModernGraphicsRenderer()
	paintPictureOnModernGraphicsEmbeddedRenderer()
}
