package main

import "golang.org/x/tour/pic"

func generatePic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8(i * j)
		}
	}
	return pic
}

func main() {
	pic.Show(generatePic)
}
