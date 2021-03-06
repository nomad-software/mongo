package mongo

import (
	"testing"
)

func TestRoundUp(t *testing.T) {
	assertValue(t, roundUp(10.75), 11)
	assertValue(t, roundUp(10.5), 11)
	assertValue(t, roundUp(10.25), 11)
	assertValue(t, roundUp(-10.75), -10)
	assertValue(t, roundUp(-10.5), -10)
	assertValue(t, roundUp(-10.25), -10)
}

func TestRoundDown(t *testing.T) {
	assertValue(t, roundDown(10.75), 10)
	assertValue(t, roundDown(10.5), 10)
	assertValue(t, roundDown(10.25), 10)
	assertValue(t, roundDown(-10.75), -11)
	assertValue(t, roundDown(-10.5), -11)
	assertValue(t, roundDown(-10.25), -11)
}

func TestRoundHalfUp(t *testing.T) {
	assertValue(t, roundHalfUp(10.75), 11)
	assertValue(t, roundHalfUp(10.5), 11)
	assertValue(t, roundHalfUp(10.25), 10)
	assertValue(t, roundHalfUp(-10.75), -11)
	assertValue(t, roundHalfUp(-10.5), -11)
	assertValue(t, roundHalfUp(-10.25), -10)
}

func TestRoundHalfDown(t *testing.T) {
	assertValue(t, roundHalfDown(10.75), 11)
	assertValue(t, roundHalfDown(10.5), 10)
	assertValue(t, roundHalfDown(10.25), 10)
	assertValue(t, roundHalfDown(-10.75), -11)
	assertValue(t, roundHalfDown(-10.5), -10)
	assertValue(t, roundHalfDown(-10.25), -10)
}

func TestRoundHalfToEven(t *testing.T) {
	assertValue(t, roundHalfToEven(11.5), 12)
	assertValue(t, roundHalfToEven(11.25), 11)
	assertValue(t, roundHalfToEven(10.75), 11)
	assertValue(t, roundHalfToEven(10.5), 10)
	assertValue(t, roundHalfToEven(10.25), 10)
	assertValue(t, roundHalfToEven(-11.5), -12)
	assertValue(t, roundHalfToEven(-11.25), -11)
	assertValue(t, roundHalfToEven(-10.75), -11)
	assertValue(t, roundHalfToEven(-10.5), -10)
	assertValue(t, roundHalfToEven(-10.25), -10)
}
