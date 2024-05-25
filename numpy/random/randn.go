package random

import (
	"fmt"

	"github.com/timotewb/gonn/app"
)

type RandnResult interface{}

// Randn generates a Gaussian noise matrix of the specified dimensions.
//
// This function creates a matrix filled with random numbers drawn from a normal distribution (Gaussian noise).
// The dimensions of the matrix are determined by the arguments passed to the function. The function supports generating:
// - A scalar value (single float) if no arguments are provided.
// - A 1D vector (slice of floats) if one argument is provided, specifying the length of the vector.
// - A 2D matrix (slice of slices of floats) if two arguments are provided, specifying the number of rows and columns.
// The generated values are normalized to have zero mean and unit variance.
//
// Parameters:
//
//	args (...int): Variable number of integer arguments specifying the dimensions of the matrix to generate.
//	               No arguments generate a scalar, one argument generates a 1D vector, and two arguments generate a 2D matrix.
//
// Returns:
//
//	RandnResult: An interface{} that holds the generated Gaussian noise matrix. The actual type depends on the number of arguments:
//	            - Scalar float64 if no arguments are provided.
//	            - Slice of float64 if one argument is provided.
//	            - Slice of slices of float64 if two arguments are provided.
//	            If an unsupported number of arguments is provided, nil is returned.
//
// Errors:
//
//	None documented.
func Randn(args ...int) RandnResult {
	switch len(args) {
	case 0:
		return app.GaussianNoise(0, 1)
	case 1:
		size := args[0]
		result := make([]float64, size)
		for i := range result {
			result[i] = app.GaussianNoise(0, 1)
		}
		return result
	case 2:
		rows, cols := args[0], args[1]
		result := make([][]float64, rows)
		for i := range result {
			result[i] = make([]float64, cols)
			for j := range result[i] {
				result[i][j] = app.GaussianNoise(0, 1)
			}
		}
		return result
	default:
		fmt.Printf("%d integers were passed: %v\n", len(args), args)
		return nil
	}
}
