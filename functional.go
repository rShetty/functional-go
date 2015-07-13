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

type filterFunc func(interface{}) bool

// Filter filters the array based on the predicate
func Filter(array interface{}, fn filterFunc) []interface{} {
	val := reflect.ValueOf(array)
	var outputArray []interface{}
	for i := 0; i < val.Len(); i++ {
		if fn(val.Index(i).Interface()) {
			outputArray = append(outputArray, val.Index(i).Interface())
		}
	}
	return outputArray
}

type foldrFunc func(interface{}, interface{}) interface{}

// Fold left the array values (reduction)
func Foldl(array interface{}, fn foldrFunc, accumulator interface{}) interface{} {
	val := reflect.ValueOf(array)
	var result = accumulator
	for i := 0; i < val.Len(); i++ {
		result = fn(val.Index(i).Interface(), result)
	}
	return result
}
