package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "hello world"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestHelloWithRecipient(t *testing.T) {
	got := helloRecipient("recipient")
	want := "hello recipient"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
