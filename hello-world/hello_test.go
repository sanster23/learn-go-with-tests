package main

import "testing"

// Using a constant, safeguard against an accidental typo
const testGreetWorld = "Hello, World"

func TestHello(t *testing.T) {
	got := Hello()
	want := testGreetWorld

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithName(t *testing.T) {
	got := HelloWithName("John")
	want := "Hello, John"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithPrefix(t *testing.T) {
	got := HelloWithPrefix("Jane")
	want := "Hello, Jane"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithCase(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("empty string", func(t *testing.T) {
		got := HelloWithCase("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people ", func(t *testing.T) {
		got := HelloWithCase("Jake", "")
		want := "Hello, Jake"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to rest of the world ", func(t *testing.T) {
		got := HelloWithCase("", "")
		want := testGreetWorld
		assertCorrectMessage(t, got, want)
	})

	t.Run("Sayin hello to Maria in Spanish ", func(t *testing.T) {
		got := HelloWithCase("Maria", "Spanish")
		want := "Hola, Maria"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Sayin hello to Elliot in French ", func(t *testing.T) {
		got := HelloWithCase("Elliot", "French")
		want := "Bonjour, Elliot"
		assertCorrectMessage(t, got, want)
	})
}
