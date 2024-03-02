package random

import (
	"fmt"

	"github.com/timotewb/gonn/app"
)

type RandnResult interface{}

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
