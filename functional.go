package functional

type mapFunc func(int) int

// Map takes array and a function
func Map(f mapFunc, array []int) []int {
	outputArray := make([]int, len(array))
	for i, a := range array {
		outputArray[i] = f(a)
	}
	return outputArray
}
