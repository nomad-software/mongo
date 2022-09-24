package mongo

import (
	"encoding/json"
	"testing"
)

func TestPriceGBP(t *testing.T) {
	p, _ := PriceGBP(6425, 20)
	assertMoneyValue(t, p.Gross(), 6425)
	assertMoneyValue(t, p.Net(), 5354)
	assertMoneyValue(t, p.Tax(), 1071)
}

func TestPriceFromSubunitsError(t *testing.T) {
	p, err := PriceFromSubunits("XXX", 1457, nil)
	if err == nil {
		t.Errorf("PriceFromSubunits failed to error on code 'XXX'")
	}

	p, err = PriceFromSubunits("GBP", 1457, nil)
	if err != nil {
		t.Errorf("PriceFromSubunits failed to recognise code 'GBP'")
	}

	assertMoneyValue(t, p.Gross(), 1457)
	assertMoneyValue(t, p.Net(), 1457)
	assertMoneyValue(t, p.Tax(), 0)

	assertSameMoneyCurrency(p.Gross(), p.Net())
	assertSameMoneyCurrency(p.Net(), p.Tax())
}

func TestPriceFromSubunits(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 1059, nil)
	p.IncludeTaxPercent(20, "VAT")

	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£8.83")
	assertMoneyString(t, p.Tax(), "GBP", "£1.76")

	assertMoneyValue(t, p.Net().Add(p.Tax()), 1059)
	assertMoneyString(t, p.Net().Add(p.Tax()), "GBP", "£10.59")
}

func TestPriceFromSubunitsNoTax(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 1059, nil)

	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£10.59")
	assertMoneyString(t, p.Tax(), "GBP", "£0.00")
}

func TestPriceFromSubunitsMaxTax(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 1059, nil)
	p.IncludeTaxSubunits(1059, "VAT")

	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£0.00")
	assertMoneyString(t, p.Tax(), "GBP", "£10.59")
}

func TestPriceFromStringError(t *testing.T) {
	p, err := PriceFromString("XXX", "£14.57", nil)
	if err == nil {
		t.Errorf("PriceFromString failed to error on code 'XXX'")
	}

	p, err = PriceFromString("GBP", "£14.57", nil)
	if err != nil {
		t.Errorf("PriceFromString failed to recognise code 'GBP'")
	}

	assertMoneyValue(t, p.Gross(), 1457)
	assertMoneyValue(t, p.Net(), 1457)
	assertMoneyValue(t, p.Tax(), 0)

	assertSameMoneyCurrency(p.Gross(), p.Net())
	assertSameMoneyCurrency(p.Net(), p.Tax())
}

func TestPriceFromString(t *testing.T) {
	p, _ := PriceFromString("GBP", "£10.59", nil)
	p.IncludeTaxPercent(20, "VAT")

	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£8.83")
	assertMoneyString(t, p.Tax(), "GBP", "£1.76")

	assertMoneyValue(t, p.Net().Add(p.Tax()), 1059)
	assertMoneyString(t, p.Net().Add(p.Tax()), "GBP", "£10.59")
}

func TestPriceFromStringNoTax(t *testing.T) {
	p, _ := PriceFromString("GBP", "£10.59", nil)

	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£10.59")
	assertMoneyString(t, p.Tax(), "GBP", "£0.00")
}

func TestPriceFromStringMaxTax(t *testing.T) {
	p, _ := PriceFromString("GBP", "£10.59", nil)
	p.IncludeTaxSubunits(1059, "VAT")

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
	assertJSON(t, bytes, `{"currency":"GBP","gross":"£10.99","net":"£9.16","tax":{"formatted":"£1.83","detail":[{"formatted":"£1.83","description":"VAT"}]}}`)

	bytes, _ = json.Marshal(resp)
	assertJSON(t, bytes, `{"name":"Widget","price":{"currency":"GBP","gross":"£10.99","net":"£9.16","tax":{"formatted":"£1.83","detail":[{"formatted":"£1.83","description":"VAT"}]}}}`)
}

func TestPriceString(t *testing.T) {
	price, _ := PriceGBP(1099, 20)
	assertPriceString(t, price, "GBP", "£10.99")
	assertPriceStringNoSymbol(t, price, "GBP", "10.99")
}

func TestAddTaxSubunits(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 5500, nil)
	p.AddTaxSubunits(825, "VAT")

	assertMoneyValue(t, p.Gross(), 6325)
	assertMoneyValue(t, p.Net(), 5500)
	assertMoneyValue(t, p.Tax(), 825)
}

func TestAddTaxPercent(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 5500, nil)
	p.AddTaxPercent(15, "VAT")

	assertMoneyValue(t, p.Gross(), 6325)
	assertMoneyValue(t, p.Net(), 5500)
	assertMoneyValue(t, p.Tax(), 825)
}

func TestIncludeTaxSubunits(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 5500, nil)
	p.IncludeTaxSubunits(825, "VAT")

	assertMoneyValue(t, p.Gross(), 5500)
	assertMoneyValue(t, p.Net(), 4675)
	assertMoneyValue(t, p.Tax(), 825)
}

func TestIncludeTaxPercent(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 2083, nil)

	p.IncludeTaxPercent(21.76, "Tax C")
	assertMoneyValue(t, p.Gross(), 2083)
	assertMoneyValue(t, p.Net(), 1711)
	assertMoneyValue(t, p.Tax(), 372)

	p.IncludeTaxPercent(6.7, "Tax B")
	assertMoneyValue(t, p.Gross(), 2083)
	assertMoneyValue(t, p.Net(), 1604)
	assertMoneyValue(t, p.Tax(), 479)

	p.IncludeTaxPercent(10, "Tax A")
	assertMoneyValue(t, p.Gross(), 2083)
	assertMoneyValue(t, p.Net(), 1458)
	assertMoneyValue(t, p.Tax(), 625)
}
