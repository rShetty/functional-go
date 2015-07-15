package functional

import "reflect"

// Definition of Map function
type mapFunc func(interface{}) interface{}

// Map maps the function onto the array
func Map(fn mapFunc, array interface{}) []interface{} {
	val := reflect.ValueOf(array)
	outputArray := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		outputArray[i] = fn(val.Index(i).Interface())
	}
	return outputArray
}

// Definition of Filter function
type filterFunc func(interface{}) bool

// Filter filters the array based on the predicate
func Filter(fn filterFunc, array interface{}) []interface{} {
	val := reflect.ValueOf(array)
	var outputArray []interface{}
	for i := 0; i < val.Len(); i++ {
		if fn(val.Index(i).Interface()) {
			outputArray = append(outputArray, val.Index(i).Interface())
		}
	}
	return outputArray
}

// Definition of Foldl function
type foldlFunc func(interface{}, interface{}) interface{}

// Folds left the array values (reduction) based on the function
func Foldl(fn foldlFunc, array interface{}, accumulator interface{}) interface{} {
	val := reflect.ValueOf(array)
	var result = accumulator
	for i := 0; i < val.Len(); i++ {
		result = fn(val.Index(i).Interface(), result)
	}
	return result
}
