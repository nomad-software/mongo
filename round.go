package mongo

import "math"

// roundFunc is the type all the following rounding functions satisfy.
type roundFunc func(float64) int64

// RoundUp is a standard rounding function that always round up.
func RoundUp(f float64) int64 {
	return int64(math.Ceil(f))
}

// RoundDown is a standard rounding function that always round down.
func RoundDown(f float64) int64 {
	return int64(math.Floor(f))
}

// RoundHalfUp is a standard rounding function that rounds 0.5 and above up.
func RoundHalfUp(f float64) int64 {
	return int64(math.Round(f))
}

// RoundHalfDown is a standard rounding function that rounds 0.5 and below down.
func RoundHalfDown(f float64) int64 {
	t := math.Trunc(f)
	if math.Abs(f-t) <= 0.5 {
		return int64(t)
	}
	return int64(t + math.Copysign(1, f))
}

// RoundHalfToEven is a standard rounding function that rounds 0.5 to the
// nearest even number. This is sometimes called bankers rounding.
func RoundHalfToEven(f float64) int64 {
	return int64(math.RoundToEven(f))
}
