package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{}

const w, h = 256, 256

func (p *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, w, h)
}

func (p *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p *Image) At(x, y int) color.Color {
	v := uint8(x + y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := &Image{}
	pic.ShowImage(m)
}
