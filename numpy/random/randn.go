package random

import "fmt"

func Randn(args ...int) {
	switch len(args) {
	case 0:
		fmt.Println("No Arguments passed.")
	case 1:
		fmt.Println("1 Argument passed.")
	case 2:
		fmt.Println("2 Argument passed.")
	default:
		fmt.Printf("%d integers were passed (Under development): %v\n", len(args), args)
	}
}
