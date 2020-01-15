package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // tells error to fall upon line where this function is called, not this actual line
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to world if no people mentioned", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

}
