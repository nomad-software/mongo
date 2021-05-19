package mongo

import (
	"testing"
)

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
