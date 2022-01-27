package main

import (
	"projectsBook/goTDD/maths/clockface"
	"testing"
	"time"
)

func TestSecondHandMidnight(t *testing.T) {
	clockTime := time.Date(1337, time.January,1,0,0,0,0,time.UTC)
	want := clockface.Point{X:150, Y:150 - 90}
	got := clockface.SecondHand(clockTime)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}