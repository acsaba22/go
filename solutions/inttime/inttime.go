package inttime

import "time"

type Second int64
type Nano int64

func ToNano(s Second) Nano {
	return Nano(s * 1e9)
}

func RoundToSecond(n Nano) Second {
	return Second(n / 1e9)
}

func FromTime(t time.Time) Nano {
	return Nano(t.UnixNano())
}
