package main

import "testing"

func TestSearch(t *testing.T) {

	t.Run("unknown word", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}

		_, got := dictionary.Search("unknown")
		want := "this is just a test"
		assertStrings(t, got, want)

	})

	t.Run("known word", func(t *testing.T) {
		_, err := dictionary.Search
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
