package nessimage

import (
	"image"
	"image/color"
	"math"
)

// BlackCap everything below this grey value is considered black.
const BlackCap = 175

//const BlackCap = 100

// ToGrayLuminance give grey luminance valllue based on composite color threshold.
func ToGrayLuminance(c color.Color) color.Gray {
	rr, gg, bb, _ := c.RGBA()
	r := math.Pow(float64(rr), 2.2)
	g := math.Pow(float64(gg), 2.2)
	b := math.Pow(float64(bb), 2.2)
	y := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
	Y := uint16(y + 0.5)
	return color.Gray{uint8(Y >> 8)}
}

// LuminanceAt return the grey lumincance at coord x,y in img.
func LuminanceAt(img image.Image, x int, y int) uint8 {
	c := img.At(x, y)
	grayColor := ToGrayLuminance(c)
	return grayColor.Y
}

// ScaledLuminanceAt return grey luminance at coord on a scaled from 0 to top
func ScaledLuminanceAt(top uint8, img image.Image, x int, y int) uint8 {
	l := LuminanceAt(img, x, y)
	return uint8((float64(l) / 255.0) * float64(top))
}

// IsBlack return tru if a point at x,y of image img is below the BlackCap.
func IsBlack(img image.Image, x int, y int) bool {
	return LuminanceAt(img, x, y) <= BlackCap
}

// IsNotBlack is the opposite to IsBlack.
func IsNotBlack(img image.Image, x int, y int) bool {
	return !IsBlack(img, x, y)
}
