package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	/*
		t.Run("saying hello to people", func(t *testing.T) {
			got := Hello("Javier")
			want := "Hello, Javier"
			assertCorrectMessage(t, got, want)
		})

		t.Run("saying hello to people", func(t *testing.T) {
			got := Hello("")
			want := "Hello, World"
			assertCorrectMessage(t, got, want)
		})
	*/

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Javier", "Spanish")
		want := "Hola, Javier"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in English", func(t *testing.T) {
		got := Hello("Javier", "English")
		want := "Hello, Javier"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Javier", "French")
		want := "Bonjour, Javier"
		assertCorrectMessage(t, got, want)
	})

}
