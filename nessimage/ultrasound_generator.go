package nessimage

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/disintegration/gift"
)

// DrawPicture .
func DrawPicture(src image.Image, width int, height int) *image.Gray {
	rect := image.Rect(0, 0, width, height)
	img := image.NewGray(rect)
	rander := rander()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var c color.Gray
			if IsBlack(src, x, y) {
				if y > 1 && IsNotBlack(src, x, y-1) {
					c = borderColor(rander)
				} else {
					c = insideColor(rander)
				}
			} else {
				c = outsideColor(rander)
			}
			x1, y1 := transformCurve(x, y, width, height)
			img.Set(x1, y1, c)
		}
	}
	g := gift.New(
	//gift.GaussianBlur(2),
	)
	img2 := image.NewGray(rect)
	g.Draw(img2, img)
	return img2
}

var minRadius = 10.0
var maxRadius = 800.0
var angle = math.Pi / 3

func transformCurve(x int, y int, width int, height int) (int, int) {
	x_offset := float64(width) / 2
	angle_offset := math.Pi / 6
	angle := 2*angle/float64(width)*float64(x) + angle_offset
	radius := ((maxRadius - minRadius) / float64(height)) * float64(y)
	xr := math.Cos(angle)*radius + x_offset
	yr := math.Sin(angle) * radius
	return int(xr), int(yr)
}

func rander() *rand.Rand {
	s1 := rand.NewSource(time.Now().Unix())
	return rand.New(s1)
}

func insideColor(rander *rand.Rand) color.Gray {
	return color.Gray{uint8(50 + rander.Intn(200))}
}

func outsideColor(rander *rand.Rand) color.Gray {
	return color.Gray{uint8(rander.Intn(5) % 2 * rander.Intn(100))}
}

func borderColor(rander *rand.Rand) color.Gray {
	var c color.Gray
	if rander.Intn(2)%2 == 0 {
		c = insideColor(rander)
	} else {
		c = outsideColor(rander)
	}
	return c
}
