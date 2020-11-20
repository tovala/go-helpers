package helpers

import (
	"fmt"
	"reflect"
)

func In(item interface{}, arrayType interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid data type: %v", arr.Kind()))
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
