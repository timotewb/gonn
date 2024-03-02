package app

import (
	"reflect"
)

// CheckIsSliceOfType checks if the provided type is a slice of a specific kind.
// It returns true if the type is a slice and its element type matches the specified kind, false otherwise.
//
// This function is used to ensure that inputs to the Multiply function are of type slice and contain
// only float64 values.
//
// Parameters:
// - t: The reflect.Type to check.
// - c: The reflect.Kind to compare against the element type of the slice.
//
// Returns:
// - bool: True if the type is a slice and its element type matches the specified kind, false otherwise.
func CheckIsSliceOfType(t reflect.Type, c reflect.Kind) bool {
	// Check if `t` is a slice
	if t.Kind() != reflect.Slice {
		return false
	}
	elemType := t.Elem()
	if elemType.Kind() == reflect.Slice {
		// Check if the element is of type `c`
		return elemType.Elem().Kind() == c
	}
	return false
}

// CheckSlicesOfSameShape checks if two slices have the same shape, including their dimensions and the
// dimensions of their sub-slices recursively. It returns true if the slices have the same shape,
// and false otherwise.
//
// This function is used to ensure that the inputs to the Multiply function have the same shape
// when they are slices.
//
// Parameters:
// - x, y: interface{} containing the single/multi-dimensional slice
//
// Returns:
// - bool: True
func CheckSlicesOfSameShape(x, y reflect.Value) bool {
	if x.Kind() != reflect.Slice || y.Kind() != reflect.Slice {
		return false
	}
	if x.Len() != y.Len() {
		return false
	}
	for i := 0; i < x.Len(); i++ {
		if x.Index(i).Kind() != reflect.Slice || y.Index(i).Kind() != reflect.Slice {
			if x.Index(i).Kind() != reflect.Float64 && y.Index(i).Kind() != reflect.Float64 {
				return false
			}
			return true
		}
		if x.Index(i).Len() != y.Index(i).Len() {
			return false
		}
		if !CheckSlicesOfSameShape(x.Index(i), y.Index(i)) {
			return false
		}
	}
	return true
}
