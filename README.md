PACKAGE DOCUMENTATION

package functional

FUNCTIONS

func Filter(array interface{}, fn filterFunc) []interface{}
    Filter filters the array based on the predicate

func Foldl(array interface{}, fn foldrFunc, accumulator interface{}) interface{}
    Folds left the array values (reduction) based on the function

func Map(array interface{}, fn mapFunc) []interface{}
    Map maps the function onto the array


