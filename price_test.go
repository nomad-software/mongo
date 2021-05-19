package mongo

import (
	"testing"
)

func TestPriceGBP(t *testing.T) {
	p, _ := PriceGBP(6425, 20)
	assertMoneyValue(t, p.Gross, 6425)
	assertMoneyValue(t, p.Net, 5354)
	assertMoneyValue(t, p.Tax, 1071)
}

func TestPriceEUR(t *testing.T) {
	p, _ := PriceEUR(2624, 20)
	assertMoneyValue(t, p.Gross, 2624)
	assertMoneyValue(t, p.Net, 2187)
	assertMoneyValue(t, p.Tax, 437)
}

func TestPriceFromSubunitsError(t *testing.T) {
	p, err := PriceFromSubunits("XXX", 1457, 20, nil)
	if err == nil {
		t.Errorf("PriceFromSubunits failed to error on code 'XXX'")
	}

	p, err = PriceFromSubunits("GBP", 1457, -0.01, nil)
	if err == nil {
		t.Errorf("PriceFromSubunits failed to error on minus tax percent")
	}

	p, err = PriceFromSubunits("GBP", 1457, 100.01, nil)
	if err == nil {
		t.Errorf("PriceFromSubunits failed to error on over 100 tax percent")
	}

	p, err = PriceFromSubunits("GBP", 1457, 20, nil)
	if err != nil {
		t.Errorf("PriceFromSubunits failed to recognise code 'GBP'")
	}

	assertMoneyValue(t, p.Gross, 1457)
	assertMoneyValue(t, p.Net, 1214)
	assertMoneyValue(t, p.Tax, 243)

	assertSameMoneyCurrency(p.Gross, p.Net)
	assertSameMoneyCurrency(p.Net, p.Tax)
}

func TestPriceFromSubunits(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 1059, 20, nil)

	assertMoneyString(t, p.Gross, "GBP", "£10.59")
	assertMoneyString(t, p.Net, "GBP", "£8.83")
	assertMoneyString(t, p.Tax, "GBP", "£1.76")

	assertMoneyValue(t, p.Net.Add(p.Tax), 1059)
	assertMoneyString(t, p.Net.Add(p.Tax), "GBP", "£10.59")
}

func TestPriceFromSubunitsNoTax(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 1059, 0, nil)

	assertMoneyString(t, p.Gross, "GBP", "£10.59")
	assertMoneyString(t, p.Net, "GBP", "£10.59")
	assertMoneyString(t, p.Tax, "GBP", "£0.00")
}

func TestPriceFromSubunitsMaxTax(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 1059, 100, nil)

	assertMoneyString(t, p.Gross, "GBP", "£10.59")
	assertMoneyString(t, p.Net, "GBP", "£0.00")
	assertMoneyString(t, p.Tax, "GBP", "£10.59")
}

func TestPriceFromStringError(t *testing.T) {
	p, err := PriceFromString("XXX", "£14.57", 20, nil)
	if err == nil {
		t.Errorf("PriceFromString failed to error on code 'XXX'")
	}

	p, err = PriceFromString("GBP", "£14.57", -0.01, nil)
	if err == nil {
		t.Errorf("PriceFromString failed to error on minus tax percent")
	}

	p, err = PriceFromString("GBP", "£14.57", 100.01, nil)
	if err == nil {
		t.Errorf("PriceFromString failed to error on over 100 tax percent")
	}

	p, err = PriceFromString("GBP", "£14.57", 20, nil)
	if err != nil {
		t.Errorf("PriceFromString failed to recognise code 'GBP'")
	}

	assertMoneyValue(t, p.Gross, 1457)
	assertMoneyValue(t, p.Net, 1214)
	assertMoneyValue(t, p.Tax, 243)

	assertSameMoneyCurrency(p.Gross, p.Net)
	assertSameMoneyCurrency(p.Net, p.Tax)
}

func TestPriceFromString(t *testing.T) {
	p, _ := PriceFromString("GBP", "£10.59", 20, nil)

	assertMoneyString(t, p.Gross, "GBP", "£10.59")
	assertMoneyString(t, p.Net, "GBP", "£8.83")
	assertMoneyString(t, p.Tax, "GBP", "£1.76")

	assertMoneyValue(t, p.Net.Add(p.Tax), 1059)
	assertMoneyString(t, p.Net.Add(p.Tax), "GBP", "£10.59")
}

func TestPriceFromStringNoTax(t *testing.T) {
	p, _ := PriceFromString("GBP", "£10.59", 0, nil)

	assertMoneyString(t, p.Gross, "GBP", "£10.59")
	assertMoneyString(t, p.Net, "GBP", "£10.59")
	assertMoneyString(t, p.Tax, "GBP", "£0.00")
}

func TestPriceFromStringMaxTax(t *testing.T) {
	p, _ := PriceFromString("GBP", "£10.59", 100, nil)

	assertMoneyString(t, p.Gross, "GBP", "£10.59")
	assertMoneyString(t, p.Net, "GBP", "£0.00")
	assertMoneyString(t, p.Tax, "GBP", "£10.59")
}

func TestPriceAdd(t *testing.T) {
	x, _ := PriceGBP(67, 20)
	y, _ := PriceGBP(33, 20)
	assertMoneyValue(t, x.Add(y).Gross, 100)
	assertMoneyValue(t, x.Add(y).Net, 84)
	assertMoneyValue(t, x.Add(y).Tax, 16)

	assertMoneyValue(t, x.Add(y).Net.Add(x.Add(y).Tax), 100)
}
