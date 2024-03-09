package main

import (
	"fmt"

	"github.com/timotewb/gonn/numpy"
)

func main() {
	// Test 01
	aX := []float64{1., 2., 3.}
	aY := []float64{1., 2., 3.}

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("Test 01")
	fmt.Print("----------------------------------------------------------------------------------------\n\n")
	fmt.Println("Result:")
	fmt.Println("\t", numpy.Dot(aX, aY))
	fmt.Print("\n----------------------------------------------------------------------------------------\n\n\n\n")

	// Test 02
	bX := [][]float64{{1., 2., 3.}, {1., 2., 3.}}
	bY := [][]float64{{1., 2., 3.}, {1., 2., 3.}}

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("Test 02")
	fmt.Print("----------------------------------------------------------------------------------------\n\n")
	fmt.Println("Result:")
	fmt.Println("\t", numpy.Dot(bX, bY))
	fmt.Print("\n----------------------------------------------------------------------------------------\n\n\n\n")
}