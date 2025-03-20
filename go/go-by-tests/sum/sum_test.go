package sum

import (
    "testing"
    "reflect"
)

func assertEqualInt(t testing.TB, got int, want int) {
    t.Helper()

    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func assertDeepCopy(t testing.TB, got []int, want []int) {
    t.Helper()

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestSum(t *testing.T) {
    t.Run("test with a variable size array", func(t *testing.T) {
        numbers := []int{1, 2}

        got := Sum(numbers)
        want := 3

        assertEqualInt(t, got, want)
    })

    t.Run("test with a size 3 array", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        assertEqualInt(t, got, want)
    })
}

func TestSumAll(t *testing.T) {
    t.Run("testing with two arrays", func(t *testing.T) {
        got := SumAll([]int{1, 2}, []int{3, 4})
        want := []int{3, 7}

        if ! reflect.DeepEqual(got, want) {
            t.Errorf("got %d, want %d", got, want)
        }
    })
}

func TestSumAllTails(t *testing.T) {
    t.Run("make the sums of tails of", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{3, 4})
        want := []int{2, 4}

        assertDeepCopy(t, got, want)
    })

    t.Run("safely sum empty slice", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4})
        want := []int{0, 4}

        assertDeepCopy(t, got, want)
    })
}
