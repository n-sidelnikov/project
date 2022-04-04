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
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

// func TestAdd(t *testing.T) {
// 	sum := Add(2, 4)
// 	expected := 6

// 	if sum != expected {
// 		t.Errorf("expected %d sum %d", expected, sum)
// 	}
// }

// func TestRepeat(t *testing.T) {
// 	repeated := Repeat("a")
// 	expected := "aaaaa"

// 	if repeated != expected {
// 		t.Errorf("expected %q but got %q", expected, repeated)
// 	}
// }
