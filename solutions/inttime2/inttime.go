package inttime2

import "time"

type Second int64
type Nano int64

func (s Second) ToNano() Nano {
	return Nano(s * 1e9)
}

func (n Nano) RoundToSecond() Second {
	return Second(n / 1e9)
}

func FromTime(t time.Time) Nano {
	return Nano(t.UnixNano())
}
