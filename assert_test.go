package mongo

import (
	"testing"
)

func assertValue(t *testing.T, value, expected int64) {
	if value != expected {
		t.Errorf("Failed asserting %d = %d (expected)\n", value, expected)
	}
}

func assertMoneyValue(t *testing.T, m Money, expected int64) {
	if m.Value() != expected {
		t.Errorf("Failed asserting %d = %d (expected)\n", m.Value(), expected)
	}
}

func assertMoneyUnits(t *testing.T, m Money, units, subunits int64) {
	if m.Units() != units {
		t.Errorf("Failed asserting units %d = %d (expected)\n", m.Units(), units)
	}

	if m.Subunits() != subunits {
		t.Errorf("Failed asserting subunits %d = %d (expected)\n", m.Subunits(), subunits)
	}
}

func assertMoneyString(t *testing.T, m Money, code string, formatted string) {
	if m.IsoCode() != code {
		t.Errorf("Failed asserting code %s = %s (expected)\n", m.IsoCode(), code)
	}

	if m.String() != formatted {
		t.Errorf("Failed asserting format %s = %s (expected)\n", m.String(), formatted)
	}
}

func assertMoneyStringNoSymbol(t *testing.T, m Money, code string, formatted string) {
	if m.IsoCode() != code {
		t.Errorf("Failed asserting code %s = %s (expected)\n", m.IsoCode(), code)
	}

	if m.StringNoSymbol() != formatted {
		t.Errorf("Failed asserting format %s = %s (expected)\n", m.StringNoSymbol(), formatted)
	}
}

func assertJSON(t *testing.T, value []byte, expected string) {
	if string(value) != expected {
		t.Errorf("Failed asserting %s = %s (expected)\n", value, expected)
	}
}

func assertPanic(t *testing.T) {
	r := recover()
	if r == nil {
		t.Errorf("Failed asserting panic in test '%s'\n", t.Name())
	}
}
