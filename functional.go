package functional

import "reflect"

type mapFunc func(interface{}) interface{}

// Map takes array and a function
func Map(array interface{}, fn mapFunc) []interface{} {
	val := reflect.ValueOf(array)
	outputArray := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		outputArray[i] = fn(val.Index(i).Interface())
	}
	return outputArray
}
