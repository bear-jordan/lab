package helloworld

import "testing"

func assertCorrectMessage(t testing.TB, got string, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

func Test_Hello(t *testing.T) {
    t.Run("default to world", func(t *testing.T) {
        got := Hello("", "en")
        want := "hello world"

        assertCorrectMessage(t, got, want)
    })

    t.Run("accept custom names", func(t *testing.T) {
        got := Hello("bear", "en")
        want := "hello bear"

        assertCorrectMessage(t, got, want)
    })

    t.Run("say hello world in english", func(t *testing.T) {
        got := Hello("world", "en")
        want := "hello world"

        assertCorrectMessage(t, got, want)
    })

    t.Run("say hello world in spanish", func(t *testing.T) {
        got := Hello("world", "es")
        want := "hola world"

        assertCorrectMessage(t, got, want)
    })

    t.Run("say hello world in french", func(t *testing.T) {
        got := Hello("world", "fr")
        want := "bonjour world"

        assertCorrectMessage(t, got, want)
    })
}
