package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/yarmand/nessimage/nessimage"
)

func main() {
	fmt.Printf("args: %v\n", os.Args[1])
	infile, err := os.Open(os.Args[1])
	check(err)
	defer infile.Close()
	img, err := png.Decode(infile)
	check(err)
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	fmt.Printf("width: %d, height: %d\n", width, height)

	out_img := nessimage.DrawPicture(img, width, height)
	outfile, err := os.Create(os.Args[2])
	check(err)
	defer outfile.Close()
	err = png.Encode(outfile, out_img)
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
