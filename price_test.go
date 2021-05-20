package mongo

import (
	"encoding/json"
	"testing"
)

func TestPriceGBP(t *testing.T) {
	p, _ := PriceGBP(6425, 20)
	assertTax(t, p.TaxPercent(), 20)
	assertMoneyValue(t, p.Gross(), 6425)
	assertMoneyValue(t, p.Net(), 5354)
	assertMoneyValue(t, p.Tax(), 1071)
}

func TestPriceEUR(t *testing.T) {
	p, _ := PriceEUR(2624, 20)
	assertTax(t, p.TaxPercent(), 20)
	assertMoneyValue(t, p.Gross(), 2624)
	assertMoneyValue(t, p.Net(), 2187)
	assertMoneyValue(t, p.Tax(), 437)
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

	p, err = PriceFromSubunits("GBP", 1457, 15, nil)
	if err != nil {
		t.Errorf("PriceFromSubunits failed to recognise code 'GBP'")
	}

	assertMoneyValue(t, p.Gross(), 1457)
	assertMoneyValue(t, p.Net(), 1267)
	assertMoneyValue(t, p.Tax(), 190)

	assertSameMoneyCurrency(p.Gross(), p.Net())
	assertSameMoneyCurrency(p.Net(), p.Tax())
}

func TestPriceFromSubunits(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 1059, 20, nil)

	assertTax(t, p.TaxPercent(), 20)
	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£8.83")
	assertMoneyString(t, p.Tax(), "GBP", "£1.76")

	assertMoneyValue(t, p.Net().Add(p.Tax()), 1059)
	assertMoneyString(t, p.Net().Add(p.Tax()), "GBP", "£10.59")
}

func TestPriceFromSubunitsNoTax(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 1059, 0, nil)

	assertTax(t, p.TaxPercent(), 0)
	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£10.59")
	assertMoneyString(t, p.Tax(), "GBP", "£0.00")
}

func TestPriceFromSubunitsMaxTax(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 1059, 100, nil)

	assertTax(t, p.TaxPercent(), 100)
	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£0.00")
	assertMoneyString(t, p.Tax(), "GBP", "£10.59")
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

	assertMoneyValue(t, p.Gross(), 1457)
	assertMoneyValue(t, p.Net(), 1214)
	assertMoneyValue(t, p.Tax(), 243)

	assertSameMoneyCurrency(p.Gross(), p.Net())
	assertSameMoneyCurrency(p.Net(), p.Tax())
}

func TestPriceFromString(t *testing.T) {
	p, _ := PriceFromString("GBP", "£10.59", 20, nil)

	assertTax(t, p.TaxPercent(), 20)
	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£8.83")
	assertMoneyString(t, p.Tax(), "GBP", "£1.76")

	assertMoneyValue(t, p.Net().Add(p.Tax()), 1059)
	assertMoneyString(t, p.Net().Add(p.Tax()), "GBP", "£10.59")
}

func TestPriceFromStringNoTax(t *testing.T) {
	p, _ := PriceFromString("GBP", "£10.59", 0, nil)

	assertTax(t, p.TaxPercent(), 0)
	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£10.59")
	assertMoneyString(t, p.Tax(), "GBP", "£0.00")
}

func TestPriceFromStringMaxTax(t *testing.T) {
	p, _ := PriceFromString("GBP", "£10.59", 100, nil)

	assertTax(t, p.TaxPercent(), 100)
	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£0.00")
	assertMoneyString(t, p.Tax(), "GBP", "£10.59")
}

func TestPriceJsonMarshalling(t *testing.T) {
	type Response struct {
		Name  string `json:"name"`
		Price Price  `json:"price"`
	}
	price, _ := PriceGBP(1099, 20)
	resp := Response{
		Name:  "Widget",
		Price: price,
	}

	bytes, _ := json.Marshal(price)
	assertJSON(t, bytes, `{"currency":"GBP","gross":"£10.99","net":"£9.16","tax":"£1.83","taxPercent":20.000000}`)

	bytes, _ = json.Marshal(resp)
	assertJSON(t, bytes, `{"name":"Widget","price":{"currency":"GBP","gross":"£10.99","net":"£9.16","tax":"£1.83","taxPercent":20.000000}}`)
}

func TestPriceString(t *testing.T) {
	price, _ := PriceGBP(1099, 20)
	assertPriceString(t, price, "GBP", "£10.99")
	assertPriceStringNoSymbol(t, price, "GBP", "10.99")
}
