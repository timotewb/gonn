package numpy

// Zeros creates a multi-dimensional slice filled with zeros based on the specified shape.
//
// The shape parameter should be a slice of integers where each integer represents the size of one dimension of the resulting multi-dimensional slice.
// The function constructs the slice starting from the innermost dimension and moving outward, ensuring that all dimensions are correctly sized.
// If the shape parameter cannot be cast to a slice of integers, the function returns nil.
//
// Example usage:
//
//	var shape = []int{2, 3, 4} // Creates a 2x3x4 tensor filled with zeros
//	zeros := Zeros(shape)
//	fmt.Printf("%v\n", zeros) // Output will depend on the environment but will represent a 2x3x4 tensor of zeros
//
// Note: The function uses dynamic typing for the shape parameter and the returned slice, allowing for flexibility but requiring careful handling of types when used.
//
// Parameters:
//
//	shape (interface{}): A slice of integers representing the shape of the desired multi-dimensional slice.
//
// Returns:
//
//	(interface{}): A multi-dimensional slice filled with zeros according to the specified shape. The actual type of the returned slice is determined dynamically based on the shape parameter.
//
// Errors:
//
//	If the shape parameter cannot be cast to a slice of integers, the function returns nil.
func Zeros(shape interface{}, removeLast bool) interface{} {
	var child interface{}

	dimSlice, ok := shape.([]int)
	if !ok {
		// Handle the error if shape is not a slice of int
		// For example, return an error or a default value
		return nil
	}
	// Remove the last element from dimSlice
	if removeLast {
		dimSlice = dimSlice[:len(dimSlice)-1]
	}

	// Start with the innermost slice as a slice of floats
	child = make([]float64, dimSlice[len(dimSlice)-1])

	// Loop backwards through the dimensions to build the nested structure
	for i := len(dimSlice) - 2; i >= 0; i-- {
		parent := make([]interface{}, dimSlice[i])
		for j := 0; j < dimSlice[i]; j++ {
			// make a copy of the child to avoid pointer
			copiedChild := make([]float64, len(child.([]float64)))
			copy(copiedChild, child.([]float64))
			parent[j] = copiedChild
		}
		child = parent
	}
	return child
}
