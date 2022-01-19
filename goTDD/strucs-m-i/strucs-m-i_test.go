package strcts_m_i_test

import (
	st "projectsBook/goTDD/strucs-m-i"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := st.Rectangle{Width: 10.0, Height: 10.0}
	got := st.Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %g, hasArea %g", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape st.Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, hasArea %g", got, want)
		}
	}
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := st.Rectangle{Width: 20.0, Height: 15.0}
		checkArea(t, rectangle, 300.0)
	})
	t.Run("Circle", func(t *testing.T) {
		circle := st.Circle{Radius: 10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestTableArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   st.Shape
		hasArea float64
	}{
		{name: "RectangleArea", shape: st.Rectangle{Width: 11, Height: 6}, hasArea: 72.0},
		{name: "CircleArea", shape: st.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "TriangleArea", shape: st.Triangle{Base: 10, Height: 10}, hasArea: 50},
	}
	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("%#v got %g, want %g", test.shape, got, test.hasArea)
			}
		})
	}
}
