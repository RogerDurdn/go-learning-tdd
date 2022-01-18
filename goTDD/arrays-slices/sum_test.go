package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}
	got := sum(numbers)
	want := 6
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := sumAll([]int{1, 2, 3}, []int{1, 2})
	want := []int{6, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkResults := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("AllTails", func(t *testing.T) {
		got := sumAllTails([]int{1, 2, 3}, []int{1, 2})
		want := []int{5, 2}
		checkResults(t, got, want)
	})

	t.Run("AllTails one empty", func(t *testing.T) {
		got := sumAllTails([]int{}, []int{1, 2})
		want := []int{0, 2}
		checkResults(t, got, want)
	})
}
