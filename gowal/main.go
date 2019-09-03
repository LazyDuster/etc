package main

import (
	"fmt"
	"os"
	"image"
	_"image/jpeg"
	_"image/png"
	"github.com/generaltso/vibrant"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	
	if len(os.Args) == 1 {
		fmt.Printf("Usage: gowal ~/path/to/img\n")
		return
	}

	fimg, err := os.Open(os.Args[1])
	checkErr(err)
	defer fimg.Close()

	f, err := os.OpenFile("output", os.O_RDWR | os.O_CREATE, 0755)
	checkErr(err)
	defer f.Close()

	img, _, err := image.Decode(fimg)
	checkErr(err)

	palette, err := vibrant.NewPaletteFromImage(img)
	checkErr(err)

	for _, swatch := range palette.ExtractAwesome() {
		var s = swatch.Color.RGBHex()
		f.WriteString(s)
		f.WriteString("\n")
	}
}