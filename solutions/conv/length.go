package conv

type Meter float64
type Foot float64

const (
	footInMeter = 0.3048
)

func MeterToFoot(m Meter) Foot {
	return Foot(m / footInMeter)
}

func FootToMeter(m Foot) Meter {
	return Meter(m * footInMeter)
}
