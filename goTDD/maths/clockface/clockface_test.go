package clockface

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time         time.Time
		lineExpected Line
	}{
		{simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			buffer := bytes.Buffer{}
			SVGWriter(&buffer, c.time)

			svg := SVG{}
			_ = xml.Unmarshal(buffer.Bytes(), &svg)

			assertContainsLine(t, c.lineExpected, svg.Lines)
		})
	}
}

func TestSVGWriterMinutesHand(t *testing.T) {
	cases := []struct {
		time         time.Time
		lineExpected Line
	}{
		{simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},

		{simpleTime(0, 30, 0),
			Line{150, 150, 150, 230}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			buffer := bytes.Buffer{}
			SVGWriter(&buffer, c.time)

			svg := SVG{}
			_ = xml.Unmarshal(buffer.Bytes(), &svg)

			assertContainsLine(t, c.lineExpected, svg.Lines)
		})
	}
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time         time.Time
		lineExpected Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			buffer := bytes.Buffer{}
			SVGWriter(&buffer, c.time)

			svg := SVG{}
			_ = xml.Unmarshal(buffer.Bytes(), &svg)

			assertContainsLine(t, c.lineExpected, svg.Lines)
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.point
			got := hourHandPoint(c.time)

			if !roughlyEqualPoints(want, got) {
				t.Fatalf("Point got: %v, but want: %v", got, want)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.point
			got := minuteHandPoint(c.time)

			if !roughlyEqualPoints(want, got) {
				t.Fatalf("Point got: %v, but want: %v", got, want)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.point
			got := secondHandPoint(c.time)

			if !roughlyEqualPoints(want, got) {
				t.Fatalf("Point got: %v, but want: %v", got, want)
			}
		})
	}
}

func assertContainsLine(t testing.TB, line Line, lines []Line) {
	t.Helper()
	present := false
	for _, l := range lines {
		if l == line {
			present = true
		}
	}
	if !present {
		t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", line, lines)
	}
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.angle
			got := secondsInRadians(c.time)

			if got != want {
				t.Fatalf("Radians got: %v, but want: %v", got, want)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 45, 0), (math.Pi / 2) * 3},
		{simpleTime(0, 7, 0), (math.Pi / 30) * 7},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.angle
			got := minutesInRadians(c.time)

			if got != want {
				t.Fatalf("Radians got: %v, but want: %v", got, want)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.angle
			got := hoursInRadians(c.time)

			if got != want {
				t.Fatalf("Radians got: %v, but want: %v", got, want)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoints(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(312, time.October, 1, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return "test-case:" + t.Format("15:04:05")
}
