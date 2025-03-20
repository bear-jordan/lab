package main

import (
    "testing"
    "bytes"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Bear")

    got := buffer.String()
    want := "Hello, Bear"

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
