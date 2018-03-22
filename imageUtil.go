package main

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"strconv"
)

func readImage() image.Image {
	dat, _ := os.Open("./images/numbers.png")
	defer dat.Close()
	img, _, err := image.Decode(dat)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return img
}

func getStartingPoint(c string) image.Point {
	if c == "," {
		return image.Pt(0, 300)
	}

	n, _ := strconv.Atoi(c)

	if n == 0 {
		return image.Pt(100, 300)
	}

	x := (n - 1) % 3
	y := (n - 1) / 3

	return image.Pt(x*100, y*100)
}

func getRect(x, y int) image.Rectangle {
	return image.Rect(x*100, 0, y*100, 100)
}

func drawFor(new draw.Image, src image.Image) func(r image.Rectangle, sp image.Point) {
	return func(r image.Rectangle, sp image.Point) {
		draw.Draw(new, r, src, sp, draw.Over)
	}
}

func shouldPrintComma(count, pos int) bool {
	return count > 3 && (count-pos)%3 == 0
}

func getImage(n int) draw.Image {
	cs := strconv.Itoa(n)

	srcImg := readImage()
	newImg := image.NewRGBA(image.Rect(0, 0, (len(cs)+(len(cs)-1)/3)*100, 100))
	d := drawFor(newImg, srcImg)

	commaOffset := 0
	for pos, c := range cs {
		if shouldPrintComma(len(cs), pos) {
			d(getRect((pos+commaOffset), (pos+1+commaOffset)), getStartingPoint(","))
			commaOffset++
		}
		d(getRect((pos+commaOffset), (pos+1+commaOffset)), getStartingPoint(string(c)))
	}

	return newImg
}
