package mongo

import "math"

// RoundFunc is the type all the following rounding functions satisfy.
type roundFunc func(float64) int64

// RoundUp is a standard rounding function that always round up.
func roundUp(f float64) int64 {
	return int64(math.Ceil(f))
}

// RoundDown is a standard rounding function that always round down.
func roundDown(f float64) int64 {
	return int64(math.Floor(f))
}

// RoundHalfUp is a standard rounding function that rounds 0.5 and above up.
func roundHalfUp(f float64) int64 {
	return int64(math.Round(f))
}

// RoundHalfDown is a standard rounding function that rounds 0.5 and below down.
func roundHalfDown(f float64) int64 {
	t := math.Trunc(f)
	if math.Abs(f-t) <= 0.5 {
		return int64(t)
	}
	return int64(t + math.Copysign(1, f))
}

// RoundHalfToEven is a standard rounding function that rounds 0.5 to the
// nearest even number. This is sometimes called bankers rounding.
func roundHalfToEven(f float64) int64 {
	return int64(math.RoundToEven(f))
}
