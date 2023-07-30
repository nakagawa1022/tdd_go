package add

import "testing"

func TestHello(t *testing.T) {

	asserCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("渡された引数に対してHelloという", func(t *testing.T) {

		got := Hello("Chris", "")
		want := "Hello, Chris"

		asserCorrectMessage(t, got, want)
	})

	t.Run("何も渡されなかった時,Hello, Worldという", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		asserCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		asserCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Frenchman", "French")
		want := "Bonjour, Frenchman"

		asserCorrectMessage(t, got, want)
	})
}
