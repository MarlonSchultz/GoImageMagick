package main

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	callImageMagic()
}

func callImageMagic(){
	imagick.Initialize()
	defer imagick.Terminate()

	magickWand := imagick.NewMagickWand()
	drawingWand := imagick.NewDrawingWand()
	pixelWand := imagick.NewPixelWand()

	magickWand.NewImage(200, 200, pixelWand)
	pixelWand.SetColor("green")
	drawingWand.SetFillColor(pixelWand)
	drawingWand.RoundRectangle(10,10,190,190,5,5)
	drawingWand.SetGravity(imagick.GRAVITY_CENTER)
	drawingWand.SetFont("Arial")
	drawingWand = setGrade(drawingWand, pixelWand, "A")
	drawingWand = setTitle(drawingWand, pixelWand, "PHP")
	magickWand.DrawImage(drawingWand)
	magickWand.WriteImage("output.jpg")
}

func setGrade(drawingWand *imagick.DrawingWand, pixelWand *imagick.PixelWand, grade string) *imagick.DrawingWand {
	drawingWand.SetFontSize(100)
	pixelWand.SetColor("black")
	drawingWand.SetStrokeColor(pixelWand)
	drawingWand.SetFillColor(pixelWand)
	drawingWand.Annotation(0,50, grade)
	return drawingWand
}

func setTitle(drawingWand *imagick.DrawingWand, pixelWand *imagick.PixelWand, title string) *imagick.DrawingWand {
	drawingWand.SetFontSize(50)
	pixelWand.SetColor("None")
	drawingWand.SetFillColor(pixelWand)
	drawingWand.Annotation(0,-30, title)
	return drawingWand
}


