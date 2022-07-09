package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func (t testing.TB, got, want string)  {

		t.Helper()
		
		if got != want {

			t.Errorf("got '%s' want '%s'", got, want)

		}

	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ahmed", "english")
		want := "Hello, Ahmed!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		got := Hello("", "english")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {

		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {

		got := Hello("Sara", "French")
		want := "Bonjour, Sara!"

		assertCorrectMessage(t, got, want)
	})
}

// func TestHello(t *testing.T) {

// 	got := Hello()
// 	want := "Hello, world!"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }
