package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 2)
	expected := strings.Repeat("a", 2)

	if repeated != expected{
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	expected := Repeat("a", 3)
	fmt.Println(expected)
//	Output: aaa

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("hello", 20)
	}
}