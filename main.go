package main

import (
	"fmt"
	"image/png"
	"os"
)

var asciiChars string = "Ñ@#W$9876543210?!abc;:+=-,._"

func main() {
	file, _ := os.Open("go-logo.png")
	img, _ := png.Decode(file)

	bounds := img.Bounds()

	imageWidth := bounds.Dx()
	scaleX := imageWidth / 120
	scaleY := scaleX * 2

	for y := bounds.Min.Y; y < bounds.Max.Y; y += scaleY {
		for x := bounds.Min.X; x < bounds.Max.X; x += scaleX {
			r16, g16, b16, a16 := img.At(x, y).RGBA()
			if float64(a16>>8) == 0 {
				fmt.Print(" ")
			} else {
				r := float64(r16 >> 8)
				g := float64(g16 >> 8)
				b := float64(b16 >> 8)

				brightness := (r + g + b) / 3.0

				index := int((brightness / 255.0) * float64(len(asciiChars)-1))
				fmt.Print(string(asciiChars[index]))
			}
		}
		fmt.Println()
	}
}
