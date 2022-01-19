package strcts_m_i_test

import (
	st "projectsBook/goTDD/strucs-m-i"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := st.Perimeter(10.0, 10.0)
	want := 40.0
	validate(t, got, want)
}

func TestArea(t *testing.T) {
	rectangle := st.Rectangle{Width: 20.0, Height: 15.0}
	got := st.Area(rectangle)
	want := 300.0
	validate(t, got, want)
}

func validate(t *testing.T, got, want float64)  {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
