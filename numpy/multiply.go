package numpy

import (
	"log"
	"reflect"

	"github.com/timotewb/gonn/app"
)

// Multiply multiplies two inputs, which can be either a single float64 number or a slice of float64 numbers,
// element-wise. The inputs must be of the same shape if they are slices. The function returns the result
// as an interface{}, which can be either a single float64 number or a slice of float64 numbers, depending
// on the inputs.
//
// If both inputs are single float64 numbers, the function returns their product as a float64 number.
// If one input is a single float64 number and the other is a slice of float64 numbers, the function returns
// a new slice where each element is the product of the float64 number and the corresponding element in
// the slice.
// If both inputs are slices of float64 numbers, the function returns a new slice where each element
// is the product of the corresponding elements in the two slices.
//
// The function supports slices of any depth. If the inputs are slices and they do not have the same
// shape, the function logs a fatal error and terminates the program.
//
// Parameters:
// - x, y: interface{} containing the single/multi-dimensional slice
//
// Returns:
// - interface: float64, slice, slice of slice(s)
func Multiply(x interface{}, y interface{}) interface{} {
	// 0 = invalid type
	// 1 = slice of float64
	// 2 = single float64
	var xType int
	var yType int
	if app.CheckIsSliceOfType(reflect.TypeOf(x), reflect.Float64) {
		xType = 1
	} else if reflect.TypeOf(x) == reflect.TypeOf(float64(0)) {
		xType = 2
	} else {
		log.Fatal("Error: x must be a slice of float64 or a single float64 number")
		return nil
	}
	if app.CheckIsSliceOfType(reflect.TypeOf(y), reflect.Float64) {
		yType = 1
	} else if reflect.TypeOf(y) == reflect.TypeOf(float64(0)) {
		yType = 2
	} else {
		log.Fatal("Error: y must be a slice of float64 or a single float64 number")
		return nil
	}

	// process
	if xType == 1 && yType == 1 {
		if !app.CheckSlicesOfSameShape(reflect.ValueOf(x), reflect.ValueOf(y)) {
			log.Fatal("Error: x and y must be the same shape.")
			return nil
		} else {
			return matrixByMatrix(x, y)
		}
	} else if xType == 1 && yType == 2 {
		return float64ByMatrix(y.(float64), x)
	} else if xType == 2 && yType == 1 {
		return float64ByMatrix(x.(float64), y)
	} else if xType == 2 && yType == 2 {
		return x.(float64) * y.(float64)
	} else {
		return nil
	}
}

// float64ByMatrix multiplies a single float64 with a multidimensional slice of any depth.
// It returns a new slice with the sme shape as the input slice where each element is the
// product of the float64 value with each float64 value in the slice.
func float64ByMatrix(n float64, m interface{}) interface{} {
	// recursively iterate over `m`
	// check if `m` is a slice
	if reflect.TypeOf(m).Kind() == reflect.Slice {
		// convert m to a reflect.Value
		mVal := reflect.ValueOf(m)
		// create a new slice of the same type and length as m
		result := reflect.MakeSlice(mVal.Type(), mVal.Len(), mVal.Len())
		// iterate over each element in m
		for i := 0; i < mVal.Len(); i++ {
			result.Index(i).Set(reflect.ValueOf(float64ByMatrix(n, mVal.Index(i).Interface())))
		}
		return result.Interface()
	} else {
		// if we find a float64, then multiply
		return n * m.(float64)
	}
}

// multiplySlicesElementWise multiplies two multidimensional slices of slices
// element-wise. It returns a new slice with the same shape as the input slices,
// where each element is the product of the corresponding elements in the input
// slices.
//
// The function supports slices of any depth. If the input slices do not have
// the same shape, the function returns nil.
func matrixByMatrix(x, y interface{}) interface{} {
	// Convert x and y to reflect.Value
	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)

	// Check if x and y are slices
	if xVal.Kind() != reflect.Slice || yVal.Kind() != reflect.Slice {
		return nil // Return nil if either x or y is not a slice
	}

	// Create a new slice of the same type and length as x
	result := reflect.MakeSlice(xVal.Type(), xVal.Len(), xVal.Len())

	// Iterate over each element in x and y
	for i := 0; i < xVal.Len(); i++ {
		// Check if the elements are slices
		if xVal.Index(i).Kind() == reflect.Slice && yVal.Index(i).Kind() == reflect.Slice {
			// Recursively multiply the slices
			result.Index(i).Set(reflect.ValueOf(matrixByMatrix(xVal.Index(i).Interface(), yVal.Index(i).Interface())))
		} else if xVal.Index(i).Kind() == reflect.Float64 && yVal.Index(i).Kind() == reflect.Float64 {
			// Multiply the float64 values
			result.Index(i).SetFloat(xVal.Index(i).Float() * yVal.Index(i).Float())
		} else {
			// Return nil if the elements are not of the expected types
			return nil
		}
	}
	return result.Interface()
}
