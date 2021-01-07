package inttime2

import (
	"testing"
	"time"
)

// https://www.epochconverter.com/

func TakesSec(s Second) {}
func TakesNano(n Nano)  {}

func TestFoot(t *testing.T) {
	format := "2006-01-02 15:04:05.000000000"
	time0, err := time.Parse(format, "2021-01-01 00:00:00.000000123")
	if err != nil {
		t.Errorf("setup failure")
	}
	var n0 Nano
	n0 = FromTime(time0)
	var s0 Second
	s0 = n0.RoundToSecond()

	if s0 != 1609459200 {
		t.Errorf("bad timestamp %d", s0)
	}

	if n0-s0.ToNano() != 123 {
		t.Errorf("bad difference")
	}
}
