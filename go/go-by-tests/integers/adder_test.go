package integers

import (
    "testing"
    "fmt"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
    got := Add(1, 1)
    want := 2

    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
