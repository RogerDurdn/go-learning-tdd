package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Roger")

	got := buffer.String()
	want := "hello Roger"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
