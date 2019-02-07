package conv

import (
	"math"
	"testing"
)

func TestFoot(t *testing.T) {
	f10 := MeterToFoot(Meter(10))
	if f10 < 30 || 40 < f10 {
		t.Errorf("30 < MeterToFoot(10) < 40 (got %v)", f10)
	}
	if 0.0001 < math.Abs(float64(FootToMeter(f10)-10)) {
		t.Errorf("FootToMeter(MeterToFoot(10)) != 10 (got: %v)",
			FootToMeter(f10))
	}
}
