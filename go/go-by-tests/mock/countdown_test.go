package main

import (
    "bytes"
    "testing"
    "reflect"
    "time"
)

func TestCountdown(t *testing.T) {
    t.Run("check output", func(t *testing.T) {
        buffer := &bytes.Buffer{}
        spyCountdownOperations := &SpyCountdownOperations{}

        Countdown(buffer, spyCountdownOperations)

        got := buffer.String()
        want := `3
2
1
Go!`

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }

    })

    t.Run("check ordering", func(t *testing.T) {
        spySleepPrinter := &SpyCountdownOperations{}
        Countdown(spySleepPrinter, spySleepPrinter)

        want := []string{
            write,
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
        }

        if ! reflect.DeepEqual(want, spySleepPrinter.Calls) {
            t.Errorf("wanted calls %v, got calls %v", want, spySleepPrinter.Calls)
        }
    })
}

func TestConfigurableSleeper(t *testing.T) {
    sleepTime := 5 * time.Second
    spyTime := &SpyTime{}

    sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
    sleeper.Sleep()

    if spyTime.durationSlept != sleepTime {
        t.Errorf("should have slept %v, but slept for %v", sleepTime, spyTime.durationSlept)
    }

}
