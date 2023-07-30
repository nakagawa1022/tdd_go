package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	t.Run("Buffer Greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	//t.Run("IO Greet", func(t *testing.T) {
	//	Greet(os.Stdout, "Elodie")
	//
	//
	//})

}
