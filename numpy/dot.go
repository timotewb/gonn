package numpy

import (
	"fmt"
	"reflect"

	"github.com/timotewb/gonn/numpy/custom"
	"github.com/timotewb/gonn/numpy/custom/logicfunctions"
)

// Dot performs dot product operations on two inputs, which can be 1D slices, 2D slices,
// or a single float and a multidimensional slice.
func Dot(x, y interface{}) (interface{}, error) {

	// Check if inputs are either slice or float
	if reflect.TypeOf(x).Kind() == reflect.Slice || reflect.TypeOf(x).Kind() == reflect.Float64 {

	}

	// Check if inputs are 1D slices (array)
	if reflect.TypeOf(x).String() == "[]float64" && reflect.TypeOf(y).String() == "[]float64" {
		return multiplyArray(x, y)
	}

	// Check if inputs are 2D slices (matrices)  and a slice (array)
	if reflect.TypeOf(x).Kind() == reflect.Slice && reflect.TypeOf(y).String() == "[]float64" {
		return multiplyAnyDimSliceBySlice(x, y)
	}

	// Check if inputs are 2D slices (matrices)
	if reflect.TypeOf(x).String() == "[][]float64" && reflect.TypeOf(y).String() == "[][]float64" {
		return multiplyMatrices(x, y)
	}

	// Check if inputs are 3D slices (matrices)
	if reflect.TypeOf(x).String() == "[][]float64" && reflect.TypeOf(y).String() == "[][][]float64" {
		return multiplyTensorByMatrix(x, y)
	}

	// Check if inputs are 3D slices (tensors)
	if reflect.TypeOf(x).String() == "testing" && reflect.TypeOf(y).String() == "testing" {
		return multiplyTensors(x, y)
	}

	// Check if x or y is a single float
	if reflect.TypeOf(x).String() == "float64" || reflect.TypeOf(y).String() == "float64" {
		return multiplyFloat(x, y)
	}

	return 0., fmt.Errorf("input combination of x (%v) and y (%v) is not supported", reflect.TypeOf(x).String(), reflect.TypeOf(y).String())
}

// multiplyFloat multiplies a single float value with each element in a multidimensional slice.
// It supports both x and y being the float, and the other being the slice.
func multiplyFloat(x, y interface{}) (interface{}, error) {
	fmt.Println("multiplyFloat()")

	// Determine which input is the float and which is the slice
	var floatVal float64
	var sliceVal reflect.Value
	if reflect.TypeOf(x).Kind() == reflect.Float64 {
		floatVal = x.(float64)
		sliceVal = reflect.ValueOf(y)
	} else if reflect.TypeOf(y).Kind() == reflect.Float64 {
		floatVal = y.(float64)
		sliceVal = reflect.ValueOf(x)
	} else {
		return nil, fmt.Errorf("one of the inputs must be a float64")
	}

	// Check if the slice is a multidimensional slice
	if sliceVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the other input must be a multidimensional slice")
	}

	// Initialize a new slice to hold the result
	result := make([]interface{}, sliceVal.Len())

	// Iterate over the slice and multiply each element by the float
	for i := 0; i < sliceVal.Len(); i++ {
		element := sliceVal.Index(i)
		if element.Kind() == reflect.Float64 {
			result[i] = element.Float() * floatVal
		} else if element.Kind() == reflect.Slice {
			result[i], _ = multiplyFloat(element.Interface(), floatVal)
		} else {
			return nil, fmt.Errorf("unsupported slice element type: %v", element.Kind())
		}
	}

	return result, nil
}

func multiplyAnyDimSliceBySlice(x, y interface{}) (interface{}, error) {

	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)
	xType := reflect.TypeOf(x)
	yType := reflect.TypeOf(y)

	xShape, err := custom.Shape(x)
	if err != nil {
		return nil, err
	}
	yShape, err := custom.Shape(y)
	if err != nil {
		return nil, err
	}

	fmt.Printf("xType %v\n", xType)
	fmt.Printf("yType %v\n", yType)
	fmt.Printf("xShape %v\n", xShape)
	fmt.Printf("yShape %v\n", yShape)

	// check compatibility
	_, err = logicfunctions.ShapeCompatible(xShape, yShape)
	if err != nil {
		return 0., err
	}

	// create output slice
	result := createMultiDimSlice(xShape)

	fmt.Println("result:", result)

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("loopOverAnyDimSlice:")
	fmt.Println("----------------------------------------------------------------------------------------")
	err = loopOverAnyDimSlice(xVal, yVal, []int{}, &result)
	if err != nil {
		return 0., err
	}
	return 0., nil
}

func loopOverAnyDimSlice(x, y reflect.Value, p []int, r *interface{}) error {

	var s float64

	// check first item
	if x.Index(0).Kind() == reflect.Slice {
		fmt.Println("x.Len():", x.Len())
		fmt.Println("")
		for i := 0; i < x.Len(); i++ {
			err := loopOverAnyDimSlice(x.Index(i), y, append(p, i), r)
			if err != nil {
				return err
			}
		}
	} else {
		for i := 0; i < x.Len(); i++ {
			s = s + (x.Index(i).Float() * y.Index(i).Float())
		}
		// with the position and value, update r
		fmt.Println("p, s:", p, s)
		err := updateValueInMultiDimArray(r, p, s)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		fmt.Println(r)

		return nil
	}
	return nil
}

func updateValueInMultiDimArray(r interface{}, p []int, newValue interface{}) error {
	// Convert r to a reflect.Value so we can manipulate it
	rVal := reflect.ValueOf(r)

	// Unwrap r if it's a pointer to an interface
	if rVal.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
	}

	// Ensure r is a multi-dimensional slice
	if rVal.Kind() != reflect.Slice {
		return fmt.Errorf("r must be a multi-dimensional slice")
	}

	// Navigate to the target position using the positions in p
	current := rVal
	for _, pos := range p {
		if current.Kind() != reflect.Slice {
			return fmt.Errorf("position %d is not a slice", pos)
		}
		current = current.Index(pos)
	}

	// Check if the target is a slice itself (indicating a deeper level needs to be accessed)
	if current.Kind() == reflect.Slice {
		return fmt.Errorf("target position is a slice, not a value")
	}

	// Convert newValue to the appropriate type
	targetType := current.Type()
	newValueVal := reflect.New(targetType).Elem()
	newValueVal.Set(reflect.ValueOf(newValue))

	// Update the value at the target position
	if current.CanSet() {
		current.Set(newValueVal)
	} else {
		return fmt.Errorf("cannot set value: target is unaddressable")
	}

	return nil
}

func createMultiDimSlice(shape interface{}) interface{} {
	var child interface{}

	dimSlice, ok := shape.([]int)
	if !ok {
		// Handle the error if shape is not a slice of int
		// For example, return an error or a default value
		return nil
	}
	fmt.Println("dimSlice:", dimSlice)
	// Remove the last element from dimSlice
	dimSlice = dimSlice[:len(dimSlice)-1]

	fmt.Println("dimSlice:", dimSlice)

	// Start with the innermost slice as a slice of floats
	child = make([]float64, dimSlice[len(dimSlice)-1])

	// Loop backwards through the dimensions to build the nested structure
	for i := len(dimSlice) - 2; i >= 0; i-- {
		parent := make([]interface{}, dimSlice[i])
		for j := 0; j < dimSlice[i]; j++ {
			parent[j] = child
		}
		child = parent
	}

	return child
}

// multiplyArray performs element-wise multiplication of two 1D slices of the same length.
func multiplyArray(x, y interface{}) (interface{}, error) {

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

func multiplyTensorByMatrix(x, y interface{}) (interface{}, error) {

	a := reflect.ValueOf(y)
	b := reflect.ValueOf(x)

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
				matrixB := b.Index(i).Interface()
				fmt.Println(matrixA)
				fmt.Println(matrixB)
				fmt.Println("")
				// Assuming resultMatrix is a 2D slice, convert it back to [][]float64
				result[i][j] = nil
			}
		}
	}

	return result, nil
}

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
