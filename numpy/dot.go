package numpy

import (
	"fmt"
	"log"
	"reflect"

	"github.com/timotewb/gonn/app"
)

// dotProduct performs a dot product operation on two slices.
// It supports 2D slices (matrices) and 3D slices (tensors).
// 1 = slice of float64
// 2 = single float64
func Dot(x, y interface{}) (interface{}, error) {

	// Check if inputs are 1D slices (array)
	if reflect.TypeOf(x).String() == "[]float64" && reflect.TypeOf(y).String() == "[]float64" {
		return multiplyArray(x, y)
	}

	// Check if inputs are 2D slices (matrices)
	if reflect.TypeOf(x).String() == "[][]float64" && reflect.TypeOf(y).String() == "[][]float64" {
		return multiplyMatrices(x, y)
	}

	// Check if inputs are 3D slices (tensors)
	if reflect.TypeOf(x).String() == "testing" && reflect.TypeOf(y).String() == "testing" {
		return multiplyTensors(x, y)
	}

	return 0., fmt.Errorf("input combination of x (%v) and y (%v) is not supported", reflect.TypeOf(x).String(), reflect.TypeOf(y).String())
}

func multiplyArray(x, y interface{}) (interface{}, error) {
	fmt.Println("multiplyArray()")

	xType := reflect.TypeOf(x)
	yType := reflect.TypeOf(y)
	xShape := fmt.Sprintf("%v", reflect.ValueOf(x).Len())
	yShape := fmt.Sprintf("%v", reflect.ValueOf(y).Len())

	if xType == yType && xShape == yShape {
		// convert x and y to a reflect.Value
		xVal := reflect.ValueOf(x)
		yVal := reflect.ValueOf(y)

		// define result float64
		var r float64
		for i := 0; i < xVal.Len(); i++ {
			r = r + (xVal.Index(i).Float() * yVal.Index(i).Float())
		}
		return r, nil
	} else {
		return 0., fmt.Errorf("x and y must be of type slice and the same shape.\n\tx: %v, %v \n\ty: %v, %v", xType, xShape, yType, yShape)
	}
}

// multiplyMatrices performs matrix multiplication on two 2D slices.
func multiplyMatrices(x, y interface{}) (interface{}, error) {
	fmt.Println("multiplyMatrices()")

	a := reflect.ValueOf(x)
	b := reflect.ValueOf(y)

	rowsA := a.Len()
	colsA := a.Index(0).Len()
	rowsB := b.Len()
	colsB := b.Index(0).Len()

	if colsA != rowsB {
		return 0., fmt.Errorf("number of columns in the first matrix must equal the number of rows in the second matrix")
	}

	result := make([][]float64, rowsA)
	for i := 0; i < rowsA; i++ {
		result[i] = make([]float64, colsB)
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += a.Index(i).Index(k).Float() * b.Index(k).Index(j).Float()
			}
		}
	}

	return result, nil
}

// multiplyTensors performs a series of matrix multiplications across the third dimension of two 3D slices.
func multiplyTensors(x, y interface{}) (interface{}, error) {
	fmt.Println("multiplyTensors()")

	a := reflect.ValueOf(x)
	b := reflect.ValueOf(y)

	// x and y must both be 3D and have the same shape
	if a.Len() != b.Len() || a.Index(0).Len() != b.Index(0).Len() || a.Index(0).Index(0).Len() != b.Index(0).Index(0).Len() {
		return 0., fmt.Errorf("x and y must both be 3D and have the same shape")
	}
	dim1 := a.Len()
	dim2 := a.Index(0).Len()
	dim3 := a.Index(0).Index(0).Len()

	result := make([][][]float64, dim1)
	for i := 0; i < dim1; i++ {
		result[i] = make([][]float64, dim2)
		for j := 0; j < dim2; j++ {
			result[i][j] = make([]float64, dim3)
			for k := 0; k < dim3; k++ {
				// Perform matrix multiplication for each matrix in the third dimension
				matrixA := a.Index(i).Index(j).Interface()
				matrixB := b.Index(i).Index(j).Interface()
				resultMatrix, err := multiplyArray(matrixA, matrixB)
				if err != nil {
					return 0., err
				}
				// Assuming resultMatrix is a 2D slice, convert it back to [][]float64
				result[i][j] = resultMatrix.([]float64)
			}
		}
	}

	return result, nil
}

func dotProduct(x interface{}, y interface{}) interface{} {
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
	}
	if app.CheckIsSliceOfType(reflect.TypeOf(y), reflect.Float64) {
		yType = 1
	} else if reflect.TypeOf(y) == reflect.TypeOf(float64(0)) {
		yType = 2
	} else {
		log.Fatal("Error: y must be a slice of float64 or a single float64 number")
	}

	// process
	if xType == 1 && yType == 1 {
		xShape := app.GetSliceShape(x)
		yShape := app.GetSliceShape(y)
		if len(xShape) < 3 && len(yShape) < 3 {

			if xShape[len(xShape)-1] != yShape[0] {
				log.Fatalf("Error: x and y must have compatible shape. \n\tx: %v \n\ty: %v", xShape, yShape)
			}

			if len(xShape) == 1 && len(yShape) == 1 {
				return rankOne(x, y)
			}

			// return rankTwo(x, y, xShape, yShape)

		} else {
			log.Fatalf("Error: Functionality not yet implemented.")
		}
		return nil
	} else if xType == 1 && yType == 2 {
		return Multiply(y, x)
	} else if xType == 2 && yType == 1 {
		return Multiply(x, y)
	} else if xType == 2 && yType == 2 {
		return Multiply(x, y)
	} else {
		return nil
	}
}

// rankOne multiplies two single dimension slices in the same way numpy.dot() functions.

func rankOne(x, y interface{}) float64 {

	xType := reflect.TypeOf(x)
	yType := reflect.TypeOf(y)
	xShape := fmt.Sprintf("%v", reflect.ValueOf(x).Len())
	yShape := fmt.Sprintf("%v", reflect.ValueOf(y).Len())

	if xType == yType && xShape == yShape {
		// convert x and y to a reflect.Value
		xVal := reflect.ValueOf(x)
		yVal := reflect.ValueOf(y)

		// define result float64
		var r float64
		for i := 0; i < xVal.Len(); i++ {
			r = r + (xVal.Index(i).Float() * yVal.Index(i).Float())
		}
		return r
	} else {
		log.Fatalf("Error: x and y must be of type slice and the same shape.\n\tx: %v, %v \n\ty: %v, %v", xType, xShape, yType, yShape)
		return 0.
	}
}

// func rankTwo(x, y interface{}, xShape, yShape []int) interface{} {
// 	fmt.Print("rankN()\n")
// 	if xShape[0] == 1 {
// 		var r []float64
// 	} else {
// 		var r [][]float64
// 	}

// 	// convert x and y to a reflect.Value
// 	xVal := reflect.ValueOf(x)
// 	// yVal := reflect.ValueOf(y)

// 	// for each array of inputs in x
// 	for i := 0; i <= xShape[0]; i++ {
// 		fmt.Print(xVal.Index(i))
// 	}

// 	return r
// }
