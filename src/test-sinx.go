package main


import (
	"image"
	"image/color"
	"math"
	"os"
	"image/png"
)

func main() {
	// size of picture
	const size = 300
	// create image base on size
	pic := image.NewGray(image.Rect(0, 0 , size, size))

	// iterate every point
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// filled with white
			pic.SetGray(x, y, color.Gray{200})
		}
	}

	// generate x coordinate from 0 to size
	for x := 0; x <size; x ++{
		// value sin between 0 and 2Pi
		s := float64(x) * 2 * math.Pi/size

		// half of sin to move below and reversal
		y := size/2 - math.Sin(s) * size/2

		// carve sin foot mark by black
		pic.SetGray(x, int(y), color.Gray{0})

	}



	file, err := os.Create("d:\\home\\gotmp\\sin.png")
	if err!=nil{
		panic(err)
	}
	// write file
	png.Encode(file, pic)

	file.Close()




}
