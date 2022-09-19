package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello World"

	got := Greeting("World", "")

	if got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestHelloWorldParam(t *testing.T) {
	expected := "Hello roger"

	got := Greeting("roger", "")

	if got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestHelloWorlds(t *testing.T) {
	t.Run("say 'Hello World' when empty string", func(t *testing.T) {
		want := "Hello World"
		got := Greeting("", "")
		assertCorrectMessage(t, want, got)
	})

	t.Run("say hello to people", func(t *testing.T) {
		want := "Hello roger"
		got := Greeting("roger", "")
		assertCorrectMessage(t, want, got)
	})

	t.Run("say hello to people in spanish", func(t *testing.T) {
		want := "Hola roger"
		got := Greeting("roger", "spanish")
		assertCorrectMessage(t, want, got)
	})

	t.Run("say hello to people in french", func(t *testing.T) {
		want := "Bonjour roger"
		got := Greeting("roger", "french")
		assertCorrectMessage(t, want, got)
	})
}

func assertCorrectMessage(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
