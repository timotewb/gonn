package numpy

import (
	"fmt"
	"reflect"

	"github.com/timotewb/gonn/numpy/custom"
	"github.com/timotewb/gonn/numpy/custom/logicfunctions"
)

// Dot performs dot product operations on two inputs, which can be 1D slices (arrays), 2D slices (matrices), and a slice (array) with a single float, based on their shapes.
//
// The function first checks if the inputs are 1D slices (arrays) and calls the appropriate handler function.
// Then, it checks if one input is a 2D slice (matrix) and the other is a slice (array), calling another handler function.
// Finally, it checks if one of the inputs is a single float, calling the appropriate handler function.
// If none of these conditions are met, it returns an error indicating that the input combination is not supported.
//
// Parameters:
//
//	x, y (interface{}): The input slices on which the dot product is performed. They can be of varying dimensions, and the function handles broadcasting rules.
//
// Returns:
//
//	(interface{}, error): The result of the dot product operation. If successful, the first element is a slice of the same shape as the input slice, with each element multiplied by the corresponding element in the other slice. If an error occurs, nil and the error are returned.
//
// Errors:
//
//	Returns an error if the shapes of x and y are incompatible, or if there are issues during the operation, such as unsupported types or dimensions.
func Dot(x, y interface{}) (interface{}, error) {

	// Check if inputs are 1D slices (arrays)
	if reflect.TypeOf(x).String() == "[]float64" && reflect.TypeOf(y).String() == "[]float64" {
		return multiplyArray(x, y)
	}
	// Check if inputs are 2D slices (matrices) and a slice (array)
	if reflect.TypeOf(x).Kind() == reflect.Slice && reflect.TypeOf(y).String() == "[]float64" {
		return multiplyAnyDimSliceBySlice(x, y)
	}
	// Check if x or y is a single float
	if reflect.TypeOf(x).String() == "float64" || reflect.TypeOf(y).String() == "float64" {
		return multiplyFloat(x, y)
	}
	// Return an error indicating that the input combination is not supported
	return 0., fmt.Errorf("input combination of x (%v) and y (%v) is not supported", reflect.TypeOf(x).String(), reflect.TypeOf(y).String())
}

// multiplyFloat multiplies a float value with each element of a multi-dimensional slice.
//
// This function is designed to handle operations between a float and a multi-dimensional slice, ensuring that the slice is properly formatted and contains only float64 elements.
// It iterates over the slice, multiplying each element by the float value, and stores the result in a new slice of the same shape.
// If the slice contains elements that are not of type float64, the function attempts to recursively apply the operation to these elements.
// If an element is not a slice and not a float64, an error is returned indicating an unsupported type.
//
// Parameters:
//
//	x, y (interface{}): The first parameter is the float value to multiply with each element of the slice. The second parameter is the multi-dimensional slice to be multiplied.
//	The slice must be a multi-dimensional slice containing only float64 elements or slices of the same structure.
//
// Returns:
//
//	(interface{}, error): A tuple containing the result of the multiplication operation. If successful, the first element is a slice of the same shape as the input slice, with each element multiplied by the float value. If an error occurs, nil and the error are returned.
//
// Errors:
//
//	Returns an error if the slice is not a multi-dimensional slice or if it contains elements that are not of type float64.
func multiplyFloat(x, y interface{}) (interface{}, error) {

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

// multiplyAnyDimSliceBySlice performs operations on two multi-dimensional slices, x and y, based on their shapes.
//
// This function is designed to handle operations between slices of varying dimensions, leveraging the concept of broadcasting similar to NumPy.
// It first determines the shapes of x and y, checks their compatibility, and then proceeds to perform the operation across matching dimensions.
// The result is stored in a newly allocated multi-dimensional slice that matches the expected output shape after the operation.
//
// The function utilizes helper functions like Shape, ShapeCompatible, and Zeros to facilitate these steps.
// It also employs a recursive strategy to traverse and operate on the multi-dimensional structures, encapsulated in loopOverAnyDimSlice.
//
// Parameters:
//
//	x, y (interface{}): The input slices on which the operation is performed. They can be of varying dimensions, and the function handles broadcasting rules.
//
// Returns:
//
//	(interface{}, error): The result of the operation, which is a multi-dimensional slice matching the expected output shape. In case of an error, nil and the error are returned.
//
// Errors:
//
//	Returns an error if the shapes of x and y are incompatible, or if there are issues during the operation, such as unsupported types or dimensions.
func multiplyAnyDimSliceBySlice(x, y interface{}) (interface{}, error) {

	// Obtain reflect.Value representations of x and y for manipulation
	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)

	// Determine the shapes of x and y
	xShape, err := custom.Shape(x)
	if err != nil {
		return nil, err
	}
	yShape, err := custom.Shape(y)
	if err != nil {
		return nil, err
	}

	// Check if the shapes are compatible for the operation
	_, err = logicfunctions.ShapeCompatible(xShape, yShape)
	if err != nil {
		return 0., err
	}

	// Create an output slice with the appropriate shape for storing the result
	result := Zeros(xShape, true)

	// Perform the operation across the multi-dimensional structure
	err = loopOverAnyDimSlice(xVal, yVal, []int{}, &result)
	if err != nil {
		return 0., err
	}

	return result, nil
}

// loopOverAnyDimSlice recursively processes a multi-dimensional slice to perform dot product calculations.
//
// This function is part of the implementation of the Dot function, which computes the dot product of two vectors or matrices.
// It operates on two reflect.Value representations of the input slices, x and y, along with a slice of integers p that tracks the current path of indices,
// and a pointer to an interface{} r that points to the result array where the computed values are stored.
// The function checks if the first element of x is another slice; if so, it recursively calls itself for each sub-slice.
// If not, it calculates the dot product for the current level of nesting and updates the result array accordingly.
// This approach allows for handling inputs of varying dimensions, including vectors, matrices, and tensors.
//
// Parameters:
//
//	x, y (reflect.Value): Representations of the input slices for which the dot product is calculated.
//	p ([]int): A slice of integers tracking the current path of indices through the multi-dimensional structure.
//	r (*interface{}): A pointer to the result array where the computed dot product values are stored.
//
// Returns:
//
//	(error): Returns an error if the operation fails due to reasons like invalid indices or incompatible types.
//	         On success, it returns nil.
//
// Errors:
//
//	If the operation encounters an unexpected type or an out-of-bounds index, an error is returned.
func loopOverAnyDimSlice(x, y reflect.Value, p []int, r *interface{}) error {

	// Accumulator for the dot product calculation at the current level of recursion
	var s float64

	// Base case: if the first element of x is not a slice, proceed with the dot product calculation
	if x.Index(0).Kind() == reflect.Slice {
		for i := 0; i < x.Len(); i++ {
			// Recursively call loopOverAnyDimSlice for each sub-slice, appending the current index to the path
			err := loopOverAnyDimSlice(x.Index(i), y, append(p, i), r)
			if err != nil {
				return err
			}
		}
	} else {
		// Calculate the dot product for the current level of nesting
		for i := 0; i < x.Len(); i++ {
			s += x.Index(i).Float() * y.Index(i).Float()
		}

		// Update the result array with the computed dot product value
		err := custom.UpdateValueInMultiDimArray(r, p, s)
		if err != nil {
			return fmt.Errorf("error when updating result array %v", err)
		}
	}

	return nil
}

// multiplyArray performs element-wise multiplication of two 1D slices of the same length.
//
// This function takes two generic interfaces as arguments, which allows for flexibility in the types of data it can process.
// It leverages reflection to determine the types and lengths of the input slices, ensuring they match before proceeding.
// If the slices are of the same type and have the same length, the function iterates over them, multiplying corresponding elements together.
// The results of these multiplications are then summed to produce a single floating-point value.
// If the slices do not meet the criteria (same type and length), an error is returned detailing the mismatch.
//
// Parameters:
//
//	x, y (interface{}): Two generic interfaces representing the slices to be multiplied. They must be of the same type and length.
//
// Returns:
//
//	(interface{}, error): A tuple containing the sum of the element-wise multiplications of the input slices and an error if applicable.
//	                    If the slices are of the correct type and length, the first element of the tuple is a float64 representing the sum.
//	                    Otherwise, the error describes the mismatch between the slices.
//
// Errors:
//
//	An error is returned if the slices are not of the same type or do not have the same length.
func multiplyArray(x, y interface{}) (interface{}, error) {

	// Retrieve the types of the input slices using reflection.
	xType := reflect.TypeOf(x)
	yType := reflect.TypeOf(y)

	// Format the lengths of the input slices for comparison.
	xShape := fmt.Sprintf("%v", reflect.ValueOf(x).Len())
	yShape := fmt.Sprintf("%v", reflect.ValueOf(y).Len())

	// Check if the slices are of the same type and have the same length.
	if xType == yType && xShape == yShape {
		xVal := reflect.ValueOf(x)
		yVal := reflect.ValueOf(y)

		// Initialize a variable to accumulate the sum of the multiplications.
		var r float64

		// Iterate over the slices, multiplying corresponding elements and adding the result to the accumulator.
		for i := 0; i < xVal.Len(); i++ {
			r = r + (xVal.Index(i).Float() * yVal.Index(i).Float())
		}

		// Return the accumulated sum and no error, indicating successful computation.
		return r, nil
	} else {
		// Construct an error message detailing the mismatch between the slices.
		return 0., fmt.Errorf("x and y must be of type slice and the same shape.\n\tx: %v, %v \n\ty: %v, %v", xType, xShape, yType, yShape)
	}
}
