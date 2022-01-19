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

func TestHelloDefault(t *testing.T) {
	t.Run("hello with default", func(t *testing.T) {
		got := helloDefault("")
		want := "hello world"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("hello with recipient", func(t *testing.T) {
		got := helloDefault("roger")
		want := "hello roger"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

// using a function to no repeat code
// refactored to use a common function
func TestAssertDefault(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("hello with default", func(t *testing.T) {
		got := helloDefault("")
		want := "hello world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("hello with recipient", func(t *testing.T) {
		got := helloDefault("roger")
		want := "hello roger"
		assertCorrectMessage(t, got, want)
	})
}
func TestAssertLanguage(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("hello with recipient and language", func(t *testing.T) {
		got := helloLanguage("roger", "spanish")
		want := "hola roger"
		assertCorrectMessage(t, got, want)
	})
	t.Run("hello with recipient and language default", func(t *testing.T) {
		got := helloLanguage("roger", "chinese")
		want := "hello roger"
		assertCorrectMessage(t, got, want)
	})
	t.Run("hello with recipient and language french", func(t *testing.T) {
		got := helloLanguage("roger", "french")
		want := "bonjour roger"
		assertCorrectMessage(t, got, want)
	})
}
