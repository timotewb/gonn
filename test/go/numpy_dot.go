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
	fmt.Print("----------------------------------------------------------------------------------------\n")
	fmt.Println("Result:")
	t1, err := numpy.Dot(aX, aY)
	if err != nil {
		fmt.Println("\t", err)
	}
	fmt.Println("\t", t1)
	fmt.Print("----------------------------------------------------------------------------------------\n\n\n\n")

	// Test 02
	bX := [][]float64{{1., 2., 3.}, {1., 2., 3.}}
	bY := [][]float64{{1., 2., 3.}, {1., 2., 3.}}

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("Test 02")
	fmt.Print("----------------------------------------------------------------------------------------\n")
	fmt.Println("Result:")
	t2, err := numpy.Dot(bX, bY)
	if err != nil {
		fmt.Println("\t", err)
	}
	fmt.Println("\t", t2)
	fmt.Print("----------------------------------------------------------------------------------------\n\n\n\n")

	// Test 03
	cX := [][]float64{{1., 2., 3.}, {1., 2., 3.}}
	cY := [][]float64{{1., 1.}, {2., 2.}, {3., 3.}}

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("Test 03")
	fmt.Print("----------------------------------------------------------------------------------------\n")
	fmt.Println("Result:")
	t3, err := numpy.Dot(cX, cY)
	if err != nil {
		fmt.Println("\t", err)
	}
	fmt.Println("\t", t3)
	fmt.Print("----------------------------------------------------------------------------------------\n\n\n\n")

	// Test 04
	dX := [][]float64{{1., 2., 3.}, {1., 2., 3.}}
	dY := [][]float64{{1., 2., 3., 4., 5.}, {1., 2., 3., 4., 5.}, {1., 2., 3., 4., 5.}}

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("Test 04")
	fmt.Print("----------------------------------------------------------------------------------------\n")
	fmt.Println("Result:")
	t4, err := numpy.Dot(dX, dY)
	if err != nil {
		fmt.Println("\t", err)
	}
	fmt.Println("\t", t4)
	fmt.Print("----------------------------------------------------------------------------------------\n\n\n\n")

	// Test 05
	eX := [][][]float64{{{1., 2., 3.}, {1., 2., 3.}, {1., 2., 3.}}, {{1., 2., 3.}, {1., 2., 3.}, {1., 2., 3.}}}
	eY := [][][]float64{{{1., 2., 3.}, {1., 2., 3.}, {1., 2., 3.}}}

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("Test 05")
	fmt.Print("----------------------------------------------------------------------------------------\n")
	fmt.Println("Result:")
	t5, err := numpy.Dot(eX, eY)
	if err != nil {
		fmt.Println("\t", err)
	}
	fmt.Println("\t", t5)
	fmt.Print("----------------------------------------------------------------------------------------\n\n\n\n")
}
