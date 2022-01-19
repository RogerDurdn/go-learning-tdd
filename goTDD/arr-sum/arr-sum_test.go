package arr_sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	arrNumbers := [5]int{1,2,3,4,5}
	result := Sum(arrNumbers)
	expected := 15
	validate(t, result, expected)
}

func TestSumSlice(t *testing.T) {
	numbers := make([]int,3)
	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	numbers = append(numbers, 30)
	result := SumSlice(numbers)
	expected := 60
	validate(t, result, expected)
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 2})
	want := []int{3, 9}
	if reflect.DeepEqual(got, want) {
		t.Errorf("result = %d, want %d", got, want)
	}
}

func validate(t *testing.T, result, expected int)  {
	if result != expected {
		t.Errorf("result = %d, want %d", result, expected)
	}
}
