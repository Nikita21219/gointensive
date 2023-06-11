package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func hexToRGBA(hex string) (color.RGBA, error) {
	var (
		rgba             color.RGBA
		err              error
		errInvalidFormat = fmt.Errorf("Invalid format")
	)

	rgba.A = 0xff
	if hex[0] != '#' {
		return rgba, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(hex) {
	case 7:
		rgba.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		rgba.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		rgba.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		rgba.R = hexToByte(hex[1]) * 17
		rgba.G = hexToByte(hex[2]) * 17
		rgba.B = hexToByte(hex[3]) * 17
	default:
		err = errInvalidFormat
	}
	return rgba, err
}

func createPNG(size int) (*image.RGBA, error) {
	bgColor, err := hexToRGBA("#764abc")
	if err != nil {
		return nil, err
	}
	bg := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(bg, bg.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)
	return bg, err
}

func main() {
	pngFile, err := createPNG(100)
	if err != nil {
		log.Fatalln("Error creating image:", err)
	}
	f, err := os.Create("favicon.ico")
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	err = png.Encode(f, pngFile)
	if err != nil {
		log.Fatalln("Error encoding png:", err)
	}
}
