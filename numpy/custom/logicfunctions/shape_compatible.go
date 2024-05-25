package logicfunctions

import (
	"fmt"
	"reflect"
)

// ShapeCompatible checks if two slices are compatible for element-wise operations based on their shapes.
//
// This function compares the shapes of two slices, x and y, to determine if they can be operated upon element-wise.
// Element-wise operations require that the slices either have the same shape or that one slice can be broadcasted to match the shape of the other.
// Specifically, this function checks if the last dimension of x matches the first dimension of y, which is a common requirement for many linear algebra operations.
// If the shapes are compatible, the function returns true; otherwise, it returns false along with an error describing the shape mismatch.
//
// Parameters:
//
//	x, y (interface{}): The slices to be checked for shape compatibility. Both must be of type slice (1D).
//
// Returns:
//
//	(bool, error): A boolean indicating whether the slices are shape-compatible for element-wise operations, and an error if the shapes are not compatible.
//	              If the shapes are compatible, the error is nil. If not, the error describes the shape mismatch.
//
// Errors:
//
//	Returns an error if the shapes of x and y are not compatible for element-wise operations.
func ShapeCompatible(x, y interface{}) (bool, error) {

	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)
	xType := xVal.Type()
	yType := yVal.Type()

	if xType.Kind() == reflect.Slice && yType.Kind() == reflect.Slice {

		if xVal.Index(xVal.Len()-1).Interface() == yVal.Index(0).Interface() {
			return true, nil
		} else {
			return false, fmt.Errorf("x and y are not compatible shapes. x: %v, y: %v", x, y)
		}
	} else {
		return false, fmt.Errorf("x and y must be of type slice (1D). x: %v, y: %v", xType.Kind().String(), yType.Kind().String())
	}
}
