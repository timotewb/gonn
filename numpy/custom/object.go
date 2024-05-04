package custom

import (
	"fmt"
	"reflect"
)

func Shape(x interface{}) (interface{}, error) {

	xVal := reflect.ValueOf(x)
	xType := xVal.Type()
	var shape []int

	// Check if x is a slice
	if xType.Kind() == reflect.Slice {
		shape = append(shape, xVal.Len())

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
		return 0., fmt.Errorf("x must be of type slice.\n\tx: %v", xType)
	}
	return shape, nil
}
