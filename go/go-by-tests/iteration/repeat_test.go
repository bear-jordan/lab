package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 3)
    }
}

func assertStringEquality(t testing.TB, got string, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func TestRepeat(t *testing.T) {
    t.Run("repeat a three times", func(t *testing.T) {
        got := Repeat("a", 3)
        want := "aaa"
        
        assertStringEquality(t, got, want)
    })

    t.Run("repeat a five times", func(t *testing.T) {
        got := Repeat("a", 5)
        want := "aaaaa"
        
        assertStringEquality(t, got, want)
    })
}
