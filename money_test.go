package mongo

import (
	"encoding/json"
	"math"
	"testing"
)

func TestGBP(t *testing.T) {
	m, _ := GBP(11)
	assertMoneyFormat(t, m, "GBP", "£0.11")
	assertMoneyValue(t, m, 11)
}

func TestEUR(t *testing.T) {
	m, _ := EUR(22)
	assertMoneyFormat(t, m, "EUR", "€0.22")
	assertMoneyValue(t, m, 22)
}

func TestScenario(t *testing.T) {
	gross, _ := GBP(1059)
	net := gross.Div(1.2)
	tax := gross.Sub(net)

	assertMoneyFormat(t, gross, "GBP", "£10.59")
	assertMoneyFormat(t, net, "GBP", "£8.83")
	assertMoneyFormat(t, tax, "GBP", "£1.76")

	assertMoneyValue(t, net.Add(tax), 1059)
	assertMoneyFormat(t, net.Add(tax), "GBP", "£10.59")
}

func TestFromSubunits(t *testing.T) {
	m, err := FromSubunits("XXX", 1457, roundHalfUp)
	if err == nil {
		t.Errorf("FromSubunits failed to error on code 'XXX'")
	}

	m, err = FromSubunits("GBP", 1457, roundHalfUp)
	if err != nil {
		t.Errorf("FromSubunits failed to recognise code 'GBP'")
	}

	assertMoneyValue(t, m, 1457)
}

func TestFromString(t *testing.T) {
	m, err := FromString("XXX", "14.57", roundHalfUp)
	if err == nil {
		t.Errorf("FromSubunits failed to error on code 'XXX'")
	}

	m, err = FromString("GBP", "14.57", roundHalfUp)
	if err != nil {
		t.Errorf("FromSubunits failed to recognise code 'GBP'")
	}

	assertMoneyValue(t, m, 1457)
}

func TestFromStringFormats(t *testing.T) {
	m, _ := FromString("JPY", "¥145139", roundHalfUp)
	assertMoneyValue(t, m, 145139)

	m, _ = FromString("EUR", "€14.57", roundHalfUp)
	assertMoneyValue(t, m, 1457)

	m, _ = FromString("JOD", "2,462.486 د.أ", roundHalfUp)
	assertMoneyValue(t, m, 2462486)

	m, _ = FromString("CLF", "UF157.896,4418", roundHalfUp)
	assertMoneyValue(t, m, 1578964418)
}

func TestFromStringErrors(t *testing.T) {
	_, err := FromString("JPY", "145139.0", roundHalfUp)
	if err == nil {
		t.Errorf("FromString failed to error on subunits on a currency that doesn't support them")
	}

	_, err = FromString("EUR", "14.570", roundHalfUp)
	if err == nil {
		t.Errorf("FromString failed to error on too many subunits defined")
	}

	_, err = FromString("JOD", "2,462.48", roundHalfUp)
	if err == nil {
		t.Errorf("FromString failed to error on too few subunits defined")
	}

	_, err = FromString("CLF", "1578964418", roundHalfUp)
	if err == nil {
		t.Errorf("FromString failed to error on no subunits defined")
	}
}

func TestImmutability(t *testing.T) {
	x, _ := GBP(69)
	y, _ := GBP(50)
	z := x.Add(y)
	z = x.Sub(y)
	z = x.Mul(20.5)
	z = x.Div(7)
	s := x.Split(3)

	assertMoneyValue(t, z, 10)
	assertMoneyValue(t, s[0], 23)
	assertMoneyValue(t, s[1], 23)
	assertMoneyValue(t, s[2], 23)
	assertMoneyValue(t, x, 69)
}

func TestPanicOnAdd(t *testing.T) {
	x, _ := GBP(69)
	y, _ := EUR(50)

	defer assertPanic(t)
	x.Add(y)
}

func TestPanicOnSub(t *testing.T) {
	x, _ := GBP(69)
	y, _ := EUR(50)

	defer assertPanic(t)
	x.Sub(y)
}

func TestPanicOnEq(t *testing.T) {
	x, _ := GBP(69)
	y, _ := EUR(50)

	defer assertPanic(t)
	x.Eq(y)
}

func TestPanicOnGt(t *testing.T) {
	x, _ := GBP(69)
	y, _ := EUR(50)

	defer assertPanic(t)
	x.Gt(y)
}

func TestPanicOnGte(t *testing.T) {
	x, _ := GBP(69)
	y, _ := EUR(50)

	defer assertPanic(t)
	x.Gte(y)
}

func TestPanicOnLt(t *testing.T) {
	x, _ := GBP(69)
	y, _ := EUR(50)

	defer assertPanic(t)
	x.Lt(y)
}

func TestPanicOnLte(t *testing.T) {
	x, _ := GBP(69)
	y, _ := EUR(50)

	defer assertPanic(t)
	x.Lte(y)
}

func TestAdd(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(33)
	assertMoneyValue(t, x.Add(y), 100)

	x, _ = GBP(-1)
	y, _ = GBP(-2)
	assertMoneyValue(t, x.Add(y), -3)
}

func TestSub(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(33)
	assertMoneyValue(t, x.Sub(y), 34)

	x, _ = GBP(-51)
	y, _ = GBP(-45)
	assertMoneyValue(t, x.Sub(y), -6)
}

func TestMul(t *testing.T) {
	x, _ := GBP(1337)
	assertMoneyValue(t, x.Mul(1.2457), 1666)
	assertMoneyValue(t, x.Mul(0.872), 1166)
	assertMoneyValue(t, x.Mul(541.544), 724044)

	x, _ = GBP(-114)
	assertMoneyValue(t, x.Mul(1.2457), -142)
	assertMoneyValue(t, x.Mul(0.872), -99)
	assertMoneyValue(t, x.Mul(541.543), -61736)

	x, _ = GBP(100)
	assertMoneyValue(t, x.Mul(5), 500)
}

func TestDiv(t *testing.T) {
	x, _ := GBP(1337)
	assertMoneyValue(t, x.Div(1.2457), 1073)
	assertMoneyValue(t, x.Div(0.871), 1535)
	assertMoneyValue(t, x.Div(541.544), 2)

	x, _ = GBP(-114)
	assertMoneyValue(t, x.Div(1.2457), -92)
	assertMoneyValue(t, x.Div(0.872), -131)
	assertMoneyValue(t, x.Div(541.543), 0)
}

func TestAbs(t *testing.T) {
	w, _ := GBP(-5434651)
	x, _ := GBP(2464125665)
	y, _ := GBP(-9007199254740992646)
	z, _ := GBP(-math.MaxInt64)
	assertMoneyValue(t, w.Abs(), 5434651)
	assertMoneyValue(t, x.Abs(), 2464125665)
	assertMoneyValue(t, y.Abs(), 9007199254740992646)
	assertMoneyValue(t, z.Abs(), math.MaxInt64)
}

func TestEq(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(33)
	assert(t, !x.Eq(y))

	x, _ = GBP(12)
	y, _ = GBP(12)
	assert(t, x.Eq(y))
}

func TestGt(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(33)
	assert(t, x.Gt(y))

	x, _ = GBP(12)
	y, _ = GBP(12)
	assert(t, !x.Gt(y))

	x, _ = GBP(5)
	y, _ = GBP(12)
	assert(t, !x.Gt(y))
}

func TestGte(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(33)
	assert(t, x.Gte(y))

	x, _ = GBP(12)
	y, _ = GBP(12)
	assert(t, x.Gte(y))

	x, _ = GBP(5)
	y, _ = GBP(12)
	assert(t, !x.Gte(y))
}

func TestLt(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(33)
	assert(t, !x.Lt(y))

	x, _ = GBP(12)
	y, _ = GBP(12)
	assert(t, !x.Lt(y))

	x, _ = GBP(5)
	y, _ = GBP(12)
	assert(t, x.Lt(y))
}

func TestLte(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(33)
	assert(t, !x.Lte(y))

	x, _ = GBP(12)
	y, _ = GBP(12)
	assert(t, x.Lte(y))

	x, _ = GBP(5)
	y, _ = GBP(12)
	assert(t, x.Lte(y))
}

func TestIsZero(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(0)
	assert(t, !x.IsZero())
	assert(t, y.IsZero())
}

func TestIsPositive(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(0)
	z, _ := GBP(-5)
	assert(t, x.IsPos())
	assert(t, y.IsPos())
	assert(t, !z.IsPos())
}

func TestIsNegative(t *testing.T) {
	x, _ := GBP(67)
	y, _ := GBP(0)
	z, _ := GBP(-5)
	assert(t, !x.IsNeg())
	assert(t, !y.IsNeg())
	assert(t, z.IsNeg())
}

func TestSplitDivideByZero(t *testing.T) {
	x, _ := GBP(100)
	defer assertPanic(t)
	x.Split(0)
}

func TestSplit(t *testing.T) {
	x, _ := GBP(100)
	s := x.Split(3)
	assertMoneyValue(t, s[0], 34)
	assertMoneyValue(t, s[1], 33)
	assertMoneyValue(t, s[2], 33)
	assertMoneyValue(t, s[0].Add(s[1]).Add(s[2]), 100)

	x, _ = GBP(123)
	s = x.Split(4)
	assertMoneyValue(t, s[0], 31)
	assertMoneyValue(t, s[1], 31)
	assertMoneyValue(t, s[2], 31)
	assertMoneyValue(t, s[3], 30)
	assertMoneyValue(t, s[0].Add(s[1]).Add(s[2]).Add(s[3]), 123)
}

func TestAllocateByZero(t *testing.T) {
	x, _ := GBP(100)
	defer assertPanic(t)
	x.Allocate(0)
}

func TestAllocateEmpty(t *testing.T) {
	x, _ := GBP(100)
	defer assertPanic(t)
	x.Allocate()
}

func TestAllocate(t *testing.T) {
	x, _ := GBP(100)
	s := x.Allocate(1, 1, 1)
	assertMoneyValue(t, s[0], 34)
	assertMoneyValue(t, s[1], 33)
	assertMoneyValue(t, s[2], 33)
	assertMoneyValue(t, s[0].Add(s[1]).Add(s[2]), 100)

	x, _ = GBP(123)
	s = x.Allocate(2, 2, 2, 2)
	assertMoneyValue(t, s[0], 31)
	assertMoneyValue(t, s[1], 31)
	assertMoneyValue(t, s[2], 31)
	assertMoneyValue(t, s[3], 30)
	assertMoneyValue(t, s[0].Add(s[1]).Add(s[2]).Add(s[3]), 123)

	x, _ = GBP(1099)
	s = x.Allocate(30, 70)
	assertMoneyValue(t, s[0], 330)
	assertMoneyValue(t, s[1], 769)
	assertMoneyValue(t, s[0].Add(s[1]), 1099)
	s = x.Allocate(305, 695)
	assertMoneyValue(t, s[0], 336)
	assertMoneyValue(t, s[1], 763)
	assertMoneyValue(t, s[0].Add(s[1]), 1099)

	x, _ = GBP(1135354247)
	s = x.Allocate(654, 465, 45565, 65, 4, 6542, 54, 574, 564, 6544, 9, 2342342, 237, 45, 34325, 2221, 111, 577, 7)
	total := s[0]
	for i := 1; i < len(s); i++ {
		total = total.Add(s[i])
	}
	assertMoneyValue(t, total, 1135354247)
}

func TestJsonMarshalling(t *testing.T) {
	type Response struct {
		Name string `json:"name"`
		Cost Money  `json:"cost"`
	}
	cost, _ := GBP(1099)
	resp := Response{
		Name: "Widget",
		Cost: cost,
	}

	bytes, _ := json.Marshal(cost)
	assertJSON(t, bytes, `{"currency":"GBP","formatted":"£10.99"}`)

	bytes, _ = json.Marshal(resp)
	assertJSON(t, bytes, `{"name":"Widget","cost":{"currency":"GBP","formatted":"£10.99"}}`)
}

func assert(t *testing.T, b bool) {
	if !b {
		t.Errorf("Failed asserting true in test '%s'\n", t.Name())
	}
}

func assertValue(t *testing.T, value, expected int64) {
	if value != expected {
		t.Errorf("Failed asserting %d = %d (expected)\n", value, expected)
	}
}

func assertMoneyValue(t *testing.T, m Money, expected int64) {
	if m.value != expected {
		t.Errorf("Failed asserting %d = %d (expected)\n", m.value, expected)
	}
}

func assertMoneyFormat(t *testing.T, m Money, code string, formatted string) {
	if m.format.code != code {
		t.Errorf("Failed asserting %s = %s (expected)\n", m.format.code, code)
	}

	if m.String() != formatted {
		t.Errorf("Failed asserting %s = %s (expected)\n", m.String(), formatted)
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
