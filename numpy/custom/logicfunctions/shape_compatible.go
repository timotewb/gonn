package logicfunctions

import (
	"fmt"
	"reflect"
)

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
