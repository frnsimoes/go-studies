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
		got := Hello("Fernando")
		want := "Hello, Fernando"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	// trying it out
	t.Run("Testing if name is capitalized", func(t *testing.T) {
		got := Hello("fernando")
		want := "Hello, Fernando"
		assertCorrectMessage(t, got, want)

	})

}
