package main

import (
	"fmt"
	"reflect"
)

// CheckSliceCompatibility checks if two multidimensional slices are compatible for multiplication.
// It returns true if the slices are compatible, false otherwise.
func CheckSliceCompatibility(x, y interface{}) bool {
	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)

	// Check if both inputs are slices
	if xVal.Kind() != reflect.Slice || yVal.Kind() != reflect.Slice {
		return false
	}

	// Get the number of dimensions for both slices
	xDim := getSliceDimension(xVal)
	yDim := getSliceDimension(yVal)

	// Check if both slices have at least two dimensions
	if xDim < 2 || yDim < 2 {
		return false
	}

	// Get the last dimension of the first slice and the second-to-last dimension of the second slice
	xLastDim := getLastDimension(xVal)
	ySecondLastDim := getSecondLastDimension(yVal)

	// Check if the dimensions are compatible for multiplication
	return xLastDim == ySecondLastDim
}

// getSliceDimension returns the number of dimensions of a slice.
func getSliceDimension(v reflect.Value) int {
	dim := 0
	for v.Kind() == reflect.Slice {
		dim++
		v = v.Index(0)
	}
	return dim
}

// getLastDimension returns the size of the last dimension of a slice.
func getLastDimension(v reflect.Value) int {
	if v.Kind() != reflect.Slice {
		return 0
	}
	for v.Kind() == reflect.Slice {
		if v.Index(0).Kind() != reflect.Slice {
			return v.Len()
		}
		v = v.Index(0)
	}
	return 0
}

// getSecondLastDimension returns the size of the second-to-last dimension of a slice.
func getSecondLastDimension(v reflect.Value) int {
	t := v
	if v.Kind() != reflect.Slice {
		return 0
	}
	// v = v.Index(0)
	i := 0
	for v.Kind() == reflect.Slice {
		if v.Index(0).Kind() != reflect.Slice {
			for j := 0; j < i-1; j++ {
				t = t.Index(0)
			}
		}
		v = v.Index(0)
		i++
	}
	return t.Len()
}

func main() {
	// Example usage
	// x := [][]float64{{1, 2, 3}, {4, 5, 6}}
	// y := [][][]float64{{{7, 8}, {9, 10}}, {{11, 12}, {13, 14}}}
	x := [][]float64{{1, 2, 3}, {1, 2, 3}}
	y := [][][]float64{{{1, 2, 3}, {1, 2, 3}}, {{1, 2, 3}, {1, 2, 3}}, {{1, 2, 3}, {1, 2, 3}}}

	fmt.Println(CheckSliceCompatibility(x, y)) // Should print true
}
