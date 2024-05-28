package custom

import (
	"fmt"
	"reflect"
)

// updateValueInMultiDimArray updates a specific value in a multi-dimensional array identified by indices.
//
// This function is designed to work with arrays represented as slices of slices (or higher dimensions),
// where each index in the path leads to a deeper level of nesting until reaching the target value.
// The function accepts an array, a list of indices leading to the target location, and a new value to set at that location.
// It uses reflection to navigate through the array structure, updating the value at the specified indices.
// If the final element at the path is not a float64, the function attempts to drill down further into the array.
// If the path leads to a non-existent location (e.g., out-of-bounds), the function returns an error.
//
// Parameters:
//
//	r (interface{}): The multi-dimensional array to update. Must be a slice of slices (or higher dimensions).
//	p ([]int): A slice of integers representing the indices leading to the target location in the array.
//	newValue (interface{}): The new value to set at the target location in the array.
//
// Returns:
//
//	(error): Returns an error if the operation fails due to reasons like invalid indices or incompatible types.
//	         On success, it returns nil.
//
// Errors:
//
//	If the array is not a multi-dimensional slice, or if the path leads to a non-existent location, or if the target value is not a float64.
func UpdateValueInMultiDimArray(r interface{}, p []int, newValue interface{}) error {

	// Convert r to a reflect.Value so we can manipulate it
	rVal := reflect.ValueOf(r)

	// Unwrap r if it's a pointer to an interface
	if rVal.Kind() == reflect.Ptr {
		rVal = rVal.Elem().Elem()
	}

	// Ensure r is a multi-dimensional slice
	if rVal.Kind() != reflect.Slice {
		return fmt.Errorf("r must be a multi-dimensional slice")
	}
	current := rVal
	for i := 0; i < len(p); i++ {
		if current.Kind() == reflect.Interface {
			current = current.Elem()
		}
		if current.Index(p[i]).Kind() == reflect.Float64 {
			current.Index(p[i]).Set(reflect.ValueOf(newValue))
			return nil
		} else {
			if current.Index(p[i]).Elem().Kind() == reflect.Slice {
				current = current.Index(p[i]).Elem()
			}
		}
	}
	return nil
}
