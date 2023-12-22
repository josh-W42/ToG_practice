package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	ans := [][]uint8{}

	for i := 0; i < dy; i++ {
		inner := make([]uint8, dx)
		ans = append(ans, inner)
	}

	return ans
}

func main() {
	pic.Show(Pic)
}
