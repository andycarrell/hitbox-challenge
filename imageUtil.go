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

func getImage(n int) draw.Image {
	cs := strconv.Itoa(n)

	srcImg := readImage()
	newImg := image.NewRGBA(image.Rect(0, 0, (len(cs)+(len(cs)-1)/3)*100, 100))

	commaOffset := 0
	for pos, c := range cs {
		if (len(cs)-pos)%3 == 0 {
			draw.Draw(newImg, image.Rect((pos+commaOffset)*100, 0, (pos+1+commaOffset)*100, 100), srcImg, getStartingPoint(","), draw.Over)
			commaOffset++
		}
		draw.Draw(newImg, image.Rect((pos+commaOffset)*100, 0, (pos+1+commaOffset)*100, 100), srcImg, getStartingPoint(string(c)), draw.Over)
	}
	return newImg
}
