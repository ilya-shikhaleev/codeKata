package main

import (
	"github.com/ilya-shikhaleev/codeKata/ood/adapter/graphics_lib"
	"github.com/ilya-shikhaleev/codeKata/ood/adapter/modern_graphics_lib"
	"github.com/ilya-shikhaleev/codeKata/ood/adapter/shape_drawing_lib"
	"os"
)

func paintPicture(painter *shape_drawing_lib.CanvasPainter) {
	triangle := shape_drawing_lib.NewTriangle(
		&shape_drawing_lib.Point{10, 15},
		&shape_drawing_lib.Point{100, 200},
		&shape_drawing_lib.Point{150, 250},
		0,
	)

	rectangle := shape_drawing_lib.NewRectangle(
		&shape_drawing_lib.Point{30, 40},
		18,
		24,
		123456,
	)

	painter.Draw(triangle)
	painter.Draw(rectangle)
}

func paintPictureOnCanvas() {
	simpleCanvas := graphics_lib.NewCanvas()
	painter := shape_drawing_lib.NewCanvasPainter(simpleCanvas)
	paintPicture(painter)
}

func paintPictureOnModernGraphicsRenderer() {
	renderer := modern_graphics_lib.NewModernGraphicsRenderer(os.Stdout)
	canvasAdapter := modern_graphics_lib.NewModernGraphicsRendererAdapter(renderer)
	defer canvasAdapter.Close()

	painter := shape_drawing_lib.NewCanvasPainter(canvasAdapter)
	paintPicture(painter)
}

func main() {
	paintPictureOnCanvas()
	paintPictureOnModernGraphicsRenderer()
}
