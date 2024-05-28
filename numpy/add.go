package numpy

import (
	"fmt"
	"reflect"

	"github.com/timotewb/gonn/numpy/custom"
	"github.com/timotewb/gonn/numpy/custom/logicfunctions"
)

func Add(x, y interface{}) (interface{}, error) {

	if reflect.TypeOf(x).Kind() == reflect.Slice || reflect.TypeOf(y).Kind() == reflect.Float64 {
		return addAnyDimSliceByFloat(x, y)
	}

	if reflect.TypeOf(x).Kind() == reflect.Slice || reflect.TypeOf(y).Kind() == reflect.Slice {
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
			return nil, err
		}
		return nil, err
	}
	return nil, nil
}

func addAnyDimSliceByFloat(x, y interface{}) (interface{}, error) {

	// Obtain reflect.Value representations of x and y for manipulation
	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)

	xShape, err := custom.Shape(x)
	if err != nil {
		return nil, err
	}

	// Create an output slice with the appropriate shape for storing the result
	result := Zeros(xShape, false)

	// Perform the operation across the multi-dimensional structure
	err = loopOverSliceAddFloat(xVal, yVal, []int{}, &result)
	if err != nil {
		return 0., err
	}

	return result, nil
}

func loopOverSliceAddFloat(x, y reflect.Value, p []int, r *interface{}) error {

	// Base case: if the first element of x is not a slice, proceed with the dot product calculation
	if x.Index(0).Kind() == reflect.Slice {
		for i := 0; i < x.Len(); i++ {
			// Recursively call loopOverAnyDimSlice for each sub-slice, appending the current index to the path
			err := loopOverSliceAddFloat(x.Index(i), y, append(p, i), r)
			if err != nil {
				return err
			}
		}
	} else {
		// Calculate the dot product for the current level of nesting
		for i := 0; i < x.Len(); i++ {

			// Update the result array with the computed dot product value
			err := custom.UpdateValueInMultiDimArray(r, append(p, i), x.Index(i).Float()+y.Float())
			if err != nil {
				return fmt.Errorf("error when updating result array %v", err)
			}
		}
	}

	return nil
}
