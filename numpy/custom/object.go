package custom

import (
	"fmt"
	"reflect"
)

// Shape retrieves the shape of a slice, returning a slice of integers representing the dimensions of the input.
//
// This function is designed to work with slices of arbitrary depth, determining the dimensions of the input slice recursively.
// It starts by checking if the input is a slice. If so, it adds the length of the slice to the shape and continues to check the first element of the slice.
// If the first element is also a slice, the function recurses, continuing to add dimensions until it reaches a non-slice element.
// If the input is not a slice, an error is returned indicating that the input must be a slice.
//
// Parameters:
//
//	x (interface{}): The input whose shape is to be retrieved. It must be a slice or a nested slice structure.
//
// Returns:
//
//	(interface{}, error): A tuple containing a slice of integers representing the shape of the input and an error if applicable.
//	                   If the input is a valid slice, the first element of the tuple is the shape. Otherwise, the error describes the issue.
//
// Errors:
//
//	Returns an error if the input is not a slice.
func Shape(x interface{}) (interface{}, error) {

	xVal := reflect.ValueOf(x)
	xType := xVal.Type()
	var shape []int

	// Check if x is a slice
	if xType.Kind() == reflect.Slice {
		shape = append(shape, xVal.Len())

		if xVal.Len() > 0 {
			// Check for multi-dimensional slices
			for {
				// Get the first element of the slice
				firstElem := xVal.Index(0)

				// If the first element is also a slice, increment dimensions and continue checking
				if firstElem.Kind() == reflect.Slice {
					shape = append(shape, firstElem.Len())
					xVal = firstElem
				} else {
					// If the first element is not a slice, break the loop
					break
				}
			}
		} else {
			return shape, nil
		}
	} else {
		return 0., fmt.Errorf("x must be of type slice.\n\tx: %v", xType)
	}
	return shape, nil
}
