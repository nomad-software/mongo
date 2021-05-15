package mongo

import (
	"testing"
)

func TestRoundUp(t *testing.T) {
	assertValue(t, RoundUp(10.75), 11)
	assertValue(t, RoundUp(10.5), 11)
	assertValue(t, RoundUp(10.25), 11)
	assertValue(t, RoundUp(-10.75), -10)
	assertValue(t, RoundUp(-10.5), -10)
	assertValue(t, RoundUp(-10.25), -10)
}

func TestRoundDown(t *testing.T) {
	assertValue(t, RoundDown(10.75), 10)
	assertValue(t, RoundDown(10.5), 10)
	assertValue(t, RoundDown(10.25), 10)
	assertValue(t, RoundDown(-10.75), -11)
	assertValue(t, RoundDown(-10.5), -11)
	assertValue(t, RoundDown(-10.25), -11)
}

func TestRoundHalfUp(t *testing.T) {
	assertValue(t, RoundHalfUp(10.75), 11)
	assertValue(t, RoundHalfUp(10.5), 11)
	assertValue(t, RoundHalfUp(10.25), 10)
	assertValue(t, RoundHalfUp(-10.75), -11)
	assertValue(t, RoundHalfUp(-10.5), -11)
	assertValue(t, RoundHalfUp(-10.25), -10)
}

func TestRoundHalfDown(t *testing.T) {
	assertValue(t, RoundHalfDown(10.75), 11)
	assertValue(t, RoundHalfDown(10.5), 10)
	assertValue(t, RoundHalfDown(10.25), 10)
	assertValue(t, RoundHalfDown(-10.75), -11)
	assertValue(t, RoundHalfDown(-10.5), -10)
	assertValue(t, RoundHalfDown(-10.25), -10)
}

func TestRoundHalfToEven(t *testing.T) {
	assertValue(t, RoundHalfToEven(11.5), 12)
	assertValue(t, RoundHalfToEven(11.25), 11)
	assertValue(t, RoundHalfToEven(10.75), 11)
	assertValue(t, RoundHalfToEven(10.5), 10)
	assertValue(t, RoundHalfToEven(10.25), 10)
	assertValue(t, RoundHalfToEven(-11.5), -12)
	assertValue(t, RoundHalfToEven(-11.25), -11)
	assertValue(t, RoundHalfToEven(-10.75), -11)
	assertValue(t, RoundHalfToEven(-10.5), -10)
	assertValue(t, RoundHalfToEven(-10.25), -10)
}
