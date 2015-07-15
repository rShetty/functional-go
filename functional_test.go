package functional

import "testing"

func TestMap(t *testing.T) {
	inputList := []int{1, 2, 3}

	square := func(x interface{}) interface{} {
		return x.(int) * x.(int)
	}

	expected := []int{1, 4, 9}
	actual := Map(square, inputList)

	for i, e := range expected {
		if e != actual[i] {
			t.Errorf("expected %v != actual %v", expected, actual)
		}
	}
}

func TestFilter(t *testing.T) {
	inputList := []int{1, 2, 3}

	modulotwo := func(x interface{}) bool {
		return x.(int)%2 == 0
	}

	expected := []int{2}
	actual := Filter(modulotwo, inputList)

	for i, e := range expected {
		if e != actual[i] {
			t.Errorf("expected %v != actual %v", expected, actual)
		}
	}
}

func TestFoldl(t *testing.T) {
	inputList := []int{1, 2, 3}

	funcFoldl := func(x interface{}, acc interface{}) interface{} {
		return x.(int) + acc.(int)
	}

	expected := 6
	accumulator := 0
	actual := Foldl(funcFoldl, inputList, accumulator)

	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}
