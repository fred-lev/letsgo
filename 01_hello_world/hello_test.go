package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Fred", "")
		want := "Hello, Fred"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Fred", "French")
		want := "Bonjour, Fred"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Fred", "Spanish")
		want := "Hola, Fred"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Fred", "Russian")
		want := "Privet, Fred"
		assertCorrectMessage(t, got, want)
	})

}
