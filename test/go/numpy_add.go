package main

import (
	"fmt"

	"github.com/timotewb/gonn/numpy"
)

func main() {

	aX := [][]float64{{1., 2., 3.}, {1.5, 2.6, 3.7}}
	aY := 2.

	fmt.Println(numpy.Add(aX, aY))
}
