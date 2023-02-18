package mongo

import (
	"encoding/json"
	"math"
	"testing"
)

func TestMoneyGBP(t *testing.T) {
	m, _ := MoneyGBP(11)
	assertMoneyString(t, m, "GBP", "£0.11")
	assertMoneyValue(t, m, 11)
}

func TestMoneyEUR(t *testing.T) {
	m, _ := MoneyEUR(22)
	assertMoneyString(t, m, "EUR", "€0.22")
	assertMoneyValue(t, m, 22)
}

func TestMoneyUSD(t *testing.T) {
	m, _ := MoneyUSD(22)
	assertMoneyString(t, m, "USD", "$0.22")
	assertMoneyValue(t, m, 22)
}

func TestMoneyScenario(t *testing.T) {
	gross, _ := MoneyGBP(1059)
	net := gross.Div(1.2)
	tax := gross.Sub(net)

	assertMoneyString(t, gross, "GBP", "£10.59")
	assertMoneyString(t, net, "GBP", "£8.83")
	assertMoneyString(t, tax, "GBP", "£1.76")

	assertMoneyValue(t, net.Add(tax), 1059)
	assertMoneyString(t, net.Add(tax), "GBP", "£10.59")
}

func TestMoneyFromSubunitsError(t *testing.T) {
	_, err := MoneyFromSubunits("XXX", 1457, RoundHalfUp)
	if err == nil {
		t.Errorf("MoneyFromSubunits failed to error on code 'XXX'")
	}

	m, err := MoneyFromSubunits("GBP", 1457, RoundHalfToEven)
	if err != nil {
		t.Errorf("MoneyFromSubunits failed to recognise code 'GBP'")
	}

	assertMoneyValue(t, m, 1457)
}

func TestMoneyFromStringError(t *testing.T) {
	_, err := MoneyFromString("XXX", "14.57", RoundHalfUp)
	if err == nil {
		t.Errorf("MoneyFromSubunits failed to error on code 'XXX'")
	}

	m, err := MoneyFromString("GBP", "14.57", RoundHalfUp)
	if err != nil {
		t.Errorf("MoneyFromSubunits failed to recognise code 'GBP'")
	}

	assertMoneyValue(t, m, 1457)
}

func TestMoneyFromStringFormats(t *testing.T) {
	m, _ := MoneyFromString("GBP", "-£1,451.39", RoundHalfUp)
	assertMoneyValue(t, m, -145139)

	m, _ = MoneyFromString("EUR", "-127.54 €", RoundHalfUp)
	assertMoneyValue(t, m, -12754)

	m, _ = MoneyFromString("DKK", "kr-127,54", RoundHalfUp)
	assertMoneyValue(t, m, -12754)

	m, _ = MoneyFromString("EUR", "€ 127.54-", RoundHalfUp)
	assertMoneyValue(t, m, -12754)

	m, _ = MoneyFromString("JPY", "¥145139", RoundHalfUp)
	assertMoneyValue(t, m, 145139)

	m, _ = MoneyFromString("EUR", "€14.57", RoundHalfUp)
	assertMoneyValue(t, m, 1457)

	m, _ = MoneyFromString("JOD", "2,462.486 د.أ", RoundHalfUp)
	assertMoneyValue(t, m, 2462486)

	m, _ = MoneyFromString("CLF", "UF157.896,4418", RoundHalfUp)
	assertMoneyValue(t, m, 1578964418)
}

func TestMoneyFromStringErrors(t *testing.T) {
	_, err := MoneyFromString("JPY", "145139.0", RoundHalfUp)
	if err == nil {
		t.Errorf("MoneyFromString failed to error on subunits on a currency that doesn't support them")
	}

	_, err = MoneyFromString("EUR", "14.570", RoundHalfUp)
	if err == nil {
		t.Errorf("MoneyFromString failed to error on too many subunits defined")
	}

	_, err = MoneyFromString("JOD", "2,462.48", RoundHalfUp)
	if err == nil {
		t.Errorf("MoneyFromString failed to error on too few subunits defined")
	}

	_, err = MoneyFromString("CLF", "1578964418", RoundHalfUp)
	if err == nil {
		t.Errorf("MoneyFromString failed to error on no subunits defined")
	}
}

func TestMoney0Subunits(t *testing.T) {
	a, _ := MoneyFromSubunits("JPY", 4, nil)
	b, _ := MoneyFromSubunits("JPY", -5, nil)
	c, _ := MoneyFromSubunits("JPY", 74, nil)
	d, _ := MoneyFromSubunits("JPY", -54, nil)
	e, _ := MoneyFromSubunits("JPY", 235, nil)
	f, _ := MoneyFromSubunits("JPY", -547, nil)

	assertMoneyUnits(t, a, 4, 0)
	assertMoneyUnits(t, b, -5, 0)
	assertMoneyUnits(t, c, 74, 0)
	assertMoneyUnits(t, d, -54, 0)
	assertMoneyUnits(t, e, 235, 0)
	assertMoneyUnits(t, f, -547, 0)
}

func TestMoney2Subunits(t *testing.T) {
	a, _ := MoneyGBP(4)
	b, _ := MoneyGBP(-5)
	c, _ := MoneyGBP(74)
	d, _ := MoneyGBP(-54)
	e, _ := MoneyGBP(235)
	f, _ := MoneyGBP(-547)

	assertMoneyUnits(t, a, 0, 4)
	assertMoneyUnits(t, b, 0, -5)
	assertMoneyUnits(t, c, 0, 74)
	assertMoneyUnits(t, d, 0, -54)
	assertMoneyUnits(t, e, 2, 35)
	assertMoneyUnits(t, f, -5, -47)
}

func TestMoney3Subunits(t *testing.T) {
	a, _ := MoneyFromSubunits("BHD", 4, nil)
	b, _ := MoneyFromSubunits("BHD", -5, nil)
	c, _ := MoneyFromSubunits("BHD", 74, nil)
	d, _ := MoneyFromSubunits("BHD", -54, nil)
	e, _ := MoneyFromSubunits("BHD", 235, nil)
	f, _ := MoneyFromSubunits("BHD", -547, nil)
	g, _ := MoneyFromSubunits("BHD", 2571, nil)
	h, _ := MoneyFromSubunits("BHD", -5741, nil)

	assertMoneyUnits(t, a, 0, 4)
	assertMoneyUnits(t, b, 0, -5)
	assertMoneyUnits(t, c, 0, 74)
	assertMoneyUnits(t, d, 0, -54)
	assertMoneyUnits(t, e, 0, 235)
	assertMoneyUnits(t, f, 0, -547)
	assertMoneyUnits(t, g, 2, 571)
	assertMoneyUnits(t, h, -5, -741)
}

func TestMoney4Subunits(t *testing.T) {
	a, _ := MoneyFromSubunits("CLF", 4, nil)
	b, _ := MoneyFromSubunits("CLF", -5, nil)
	c, _ := MoneyFromSubunits("CLF", 74, nil)
	d, _ := MoneyFromSubunits("CLF", -54, nil)
	e, _ := MoneyFromSubunits("CLF", 235, nil)
	f, _ := MoneyFromSubunits("CLF", -547, nil)
	g, _ := MoneyFromSubunits("CLF", 2571, nil)
	h, _ := MoneyFromSubunits("CLF", -5741, nil)
	i, _ := MoneyFromSubunits("CLF", 57374, nil)
	j, _ := MoneyFromSubunits("CLF", -75728, nil)

	assertMoneyUnits(t, a, 0, 4)
	assertMoneyUnits(t, b, 0, -5)
	assertMoneyUnits(t, c, 0, 74)
	assertMoneyUnits(t, d, 0, -54)
	assertMoneyUnits(t, e, 0, 235)
	assertMoneyUnits(t, f, 0, -547)
	assertMoneyUnits(t, g, 0, 2571)
	assertMoneyUnits(t, h, 0, -5741)
	assertMoneyUnits(t, i, 5, 7374)
	assertMoneyUnits(t, j, -7, -5728)
}

func TestMoneyImmutability(t *testing.T) {
	x, _ := MoneyGBP(69)
	y, _ := MoneyGBP(50)
	z := x.Add(y)
	z = x.Sub(y)
	z = x.Mul(20)
	z = x.Div(7)
	z = x.FlipSign()
	s := x.Split(3)

	assertMoneyValue(t, z, -69)
	assertMoneyValue(t, s[0], 23)
	assertMoneyValue(t, s[1], 23)
	assertMoneyValue(t, s[2], 23)
	assertMoneyValue(t, x, 69)
}

func TestCurrencyPanicOnAdd(t *testing.T) {
	x, _ := MoneyGBP(69)
	y, _ := MoneyEUR(50)

	defer assertPanic(t)
	x.Add(y)
}

func TestCurrencyPanicOnSub(t *testing.T) {
	x, _ := MoneyGBP(69)
	y, _ := MoneyEUR(50)

	defer assertPanic(t)
	x.Sub(y)
}

func TestCurrencyPanicOnEq(t *testing.T) {
	x, _ := MoneyGBP(69)
	y, _ := MoneyEUR(50)

	defer assertPanic(t)
	x.Eq(y)
}

func TestCurrencyPanicOnGt(t *testing.T) {
	x, _ := MoneyGBP(69)
	y, _ := MoneyEUR(50)

	defer assertPanic(t)
	x.Gt(y)
}

func TestCurrencyPanicOnGte(t *testing.T) {
	x, _ := MoneyGBP(69)
	y, _ := MoneyEUR(50)

	defer assertPanic(t)
	x.Gte(y)
}

func TestCurrencyPanicOnLt(t *testing.T) {
	x, _ := MoneyGBP(69)
	y, _ := MoneyEUR(50)

	defer assertPanic(t)
	x.Lt(y)
}

func TestCurrencyPanicOnLte(t *testing.T) {
	x, _ := MoneyGBP(69)
	y, _ := MoneyEUR(50)

	defer assertPanic(t)
	x.Lte(y)
}

func TestMoneyAdd(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(33)
	assertMoneyValue(t, x.Add(y), 100)

	x, _ = MoneyGBP(-1)
	y, _ = MoneyGBP(-2)
	assertMoneyValue(t, x.Add(y), -3)
}

func TestMoneySub(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(33)
	assertMoneyValue(t, x.Sub(y), 34)

	x, _ = MoneyGBP(-51)
	y, _ = MoneyGBP(-45)
	assertMoneyValue(t, x.Sub(y), -6)
}

func TestMoneyMul(t *testing.T) {
	x, _ := MoneyGBP(1337)
	assertMoneyValue(t, x.Mul(0), 0)
	assertMoneyValue(t, x.Mul(1), 1337)
	assertMoneyValue(t, x.Mul(3), 4011)

	x, _ = MoneyGBP(-114)
	assertMoneyValue(t, x.Mul(0), 0)
	assertMoneyValue(t, x.Mul(1), -114)
	assertMoneyValue(t, x.Mul(3), -342)

	x, _ = MoneyGBP(100)
	assertMoneyValue(t, x.Mul(5), 500)
}

func TestMoneyDiv(t *testing.T) {
	x, _ := MoneyGBP(1337)
	assertMoneyValue(t, x.Div(1.2457), 1073)
	assertMoneyValue(t, x.Div(0.871), 1535)
	assertMoneyValue(t, x.Div(541.544), 2)

	x, _ = MoneyFromSubunits("GBP", 1337, RoundUp)
	assertMoneyValue(t, x.Div(1.2457), 1074)
	assertMoneyValue(t, x.Div(0.871), 1536)
	assertMoneyValue(t, x.Div(541.544), 3)

	x, _ = MoneyGBP(-114)
	assertMoneyValue(t, x.Div(1.2457), -92)
	assertMoneyValue(t, x.Div(0.872), -131)
	assertMoneyValue(t, x.Div(541.543), 0)

	x, _ = MoneyFromSubunits("GBP", -114, RoundUp)
	assertMoneyValue(t, x.Div(1.2457), -91)
	assertMoneyValue(t, x.Div(0.872), -130)
	assertMoneyValue(t, x.Div(541.543), 0)
}

func TestMoneyAbs(t *testing.T) {
	w, _ := MoneyGBP(-5434651)
	x, _ := MoneyGBP(2464125665)
	y, _ := MoneyGBP(-9007199254740992646)
	z, _ := MoneyGBP(-math.MaxInt64)
	assertMoneyValue(t, w.Abs(), 5434651)
	assertMoneyValue(t, x.Abs(), 2464125665)
	assertMoneyValue(t, y.Abs(), 9007199254740992646)
	assertMoneyValue(t, z.Abs(), math.MaxInt64)
}

func TestMoneyFlipSign(t *testing.T) {
	x, _ := MoneyGBP(1337)
	assertMoneyValue(t, x.FlipSign(), -1337)

	x, _ = MoneyGBP(-114)
	assertMoneyValue(t, x.FlipSign(), 114)
}

func TestMoneyEq(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(33)
	assert(t, !x.Eq(y))

	x, _ = MoneyGBP(12)
	y, _ = MoneyGBP(12)
	assert(t, x.Eq(y))
}

func TestMoneyNeq(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(33)
	assert(t, x.Neq(y))

	x, _ = MoneyGBP(12)
	y, _ = MoneyGBP(12)
	assert(t, !x.Neq(y))
}

func TestMoneyGt(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(33)
	assert(t, x.Gt(y))

	x, _ = MoneyGBP(12)
	y, _ = MoneyGBP(12)
	assert(t, !x.Gt(y))

	x, _ = MoneyGBP(5)
	y, _ = MoneyGBP(12)
	assert(t, !x.Gt(y))
}

func TestMoneyGte(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(33)
	assert(t, x.Gte(y))

	x, _ = MoneyGBP(12)
	y, _ = MoneyGBP(12)
	assert(t, x.Gte(y))

	x, _ = MoneyGBP(5)
	y, _ = MoneyGBP(12)
	assert(t, !x.Gte(y))
}

func TestMoneyLt(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(33)
	assert(t, !x.Lt(y))

	x, _ = MoneyGBP(12)
	y, _ = MoneyGBP(12)
	assert(t, !x.Lt(y))

	x, _ = MoneyGBP(5)
	y, _ = MoneyGBP(12)
	assert(t, x.Lt(y))
}

func TestMoneyLte(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(33)
	assert(t, !x.Lte(y))

	x, _ = MoneyGBP(12)
	y, _ = MoneyGBP(12)
	assert(t, x.Lte(y))

	x, _ = MoneyGBP(5)
	y, _ = MoneyGBP(12)
	assert(t, x.Lte(y))
}

func TestMoneyIsZero(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(0)
	assert(t, !x.IsZero())
	assert(t, y.IsZero())
}

func TestMoneyIsPositive(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(0)
	z, _ := MoneyGBP(-5)
	assert(t, x.IsPos())
	assert(t, y.IsPos())
	assert(t, !z.IsPos())
}

func TestMoneyIsNegative(t *testing.T) {
	x, _ := MoneyGBP(67)
	y, _ := MoneyGBP(0)
	z, _ := MoneyGBP(-5)
	assert(t, !x.IsNeg())
	assert(t, !y.IsNeg())
	assert(t, z.IsNeg())
}

func TestMoneySplitDivideByZero(t *testing.T) {
	x, _ := MoneyGBP(100)
	defer assertPanic(t)
	x.Split(0)
}

func TestMoneySplit(t *testing.T) {
	x, _ := MoneyGBP(100)
	s := x.Split(3)
	assertMoneyValue(t, s[0], 34)
	assertMoneyValue(t, s[1], 33)
	assertMoneyValue(t, s[2], 33)
	assertMoneyValue(t, s[0].Add(s[1]).Add(s[2]), 100)

	x, _ = MoneyGBP(123)
	s = x.Split(4)
	assertMoneyValue(t, s[0], 31)
	assertMoneyValue(t, s[1], 31)
	assertMoneyValue(t, s[2], 31)
	assertMoneyValue(t, s[3], 30)
	assertMoneyValue(t, s[0].Add(s[1]).Add(s[2]).Add(s[3]), 123)
}

func TestMoneyAllocateByZero(t *testing.T) {
	x, _ := MoneyGBP(100)
	defer assertPanic(t)
	x.Allocate(0)
}

func TestMoneyAllocateEmpty(t *testing.T) {
	x, _ := MoneyGBP(100)
	defer assertPanic(t)
	x.Allocate()
}

func TestMoneyAllocate(t *testing.T) {
	x, _ := MoneyGBP(100)
	s := x.Allocate(1, 1, 1)
	assertMoneyValue(t, s[0], 34)
	assertMoneyValue(t, s[1], 33)
	assertMoneyValue(t, s[2], 33)
	assertMoneyValue(t, s[0].Add(s[1]).Add(s[2]), 100)

	x, _ = MoneyGBP(123)
	s = x.Allocate(2, 2, 2, 2)
	assertMoneyValue(t, s[0], 31)
	assertMoneyValue(t, s[1], 31)
	assertMoneyValue(t, s[2], 31)
	assertMoneyValue(t, s[3], 30)
	assertMoneyValue(t, s[0].Add(s[1]).Add(s[2]).Add(s[3]), 123)

	x, _ = MoneyGBP(1099)
	s = x.Allocate(30, 70)
	assertMoneyValue(t, s[0], 330)
	assertMoneyValue(t, s[1], 769)
	assertMoneyValue(t, s[0].Add(s[1]), 1099)
	s = x.Allocate(305, 695)
	assertMoneyValue(t, s[0], 336)
	assertMoneyValue(t, s[1], 763)
	assertMoneyValue(t, s[0].Add(s[1]), 1099)

	x, _ = MoneyGBP(1135354247)
	s = x.Allocate(654, 465, 45565, 65, 4, 6542, 54, 574, 564, 6544, 9, 2342342, 237, 45, 34325, 2221, 111, 577, 7)
	total := s[0]
	for i := 1; i < len(s); i++ {
		total = total.Add(s[i])
	}
	assertMoneyValue(t, total, 1135354247)
}

func TestMoneyJsonMarshalling(t *testing.T) {
	type Response struct {
		Name string `json:"name"`
		Cost Money  `json:"cost"`
	}
	cost, _ := MoneyGBP(1099)
	resp := Response{
		Name: "Widget",
		Cost: cost,
	}

	bytes, _ := json.Marshal(cost)
	assertJSON(t, bytes, `{"currency":"GBP","amount":"£10.99"}`)

	bytes, _ = json.Marshal(resp)
	assertJSON(t, bytes, `{"name":"Widget","cost":{"currency":"GBP","amount":"£10.99"}}`)
}
