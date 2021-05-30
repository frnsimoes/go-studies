package main

import "testing"

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got int, want int, numbers [5]int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	}

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want, numbers)
	})

}

func TestSumWithSlices(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got int, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	}

	t.Run("Using slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlice(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)
	})
}
