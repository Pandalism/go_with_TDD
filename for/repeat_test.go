package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // tells error to fall upon line where this function is called, not this actual line
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("repeating a five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeating b seven times", func(t *testing.T) {
		repeated := Repeat("b", 7)
		expected := "bbbbbbb"

		assertCorrectMessage(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 6)
	fmt.Println(repeated)
	// Output: cccccc
}
