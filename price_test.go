package mongo

import (
	"encoding/json"
	"testing"
)

func TestGenericIntegerPriceTypes(t *testing.T) {
	p, _ := PriceGBP(int8(104), 20)
	assertMoneyValue(t, p.Gross(), 104)
	assertMoneyValue(t, p.Net(), 87)
	assertMoneyValue(t, p.Tax(), 17)

	p, _ = PriceGBP(uint8(104), 20)
	assertMoneyValue(t, p.Gross(), 104)
	assertMoneyValue(t, p.Net(), 87)
	assertMoneyValue(t, p.Tax(), 17)

	p, _ = PriceGBP(int16(1528), 20)
	assertMoneyValue(t, p.Gross(), 1528)
	assertMoneyValue(t, p.Net(), 1273)
	assertMoneyValue(t, p.Tax(), 255)

	p, _ = PriceGBP(uint16(1528), 20)
	assertMoneyValue(t, p.Gross(), 1528)
	assertMoneyValue(t, p.Net(), 1273)
	assertMoneyValue(t, p.Tax(), 255)

	p, _ = PriceGBP(int32(68242), 20)
	assertMoneyValue(t, p.Gross(), 68242)
	assertMoneyValue(t, p.Net(), 56868)
	assertMoneyValue(t, p.Tax(), 11374)

	p, _ = PriceGBP(uint32(68242), 20)
	assertMoneyValue(t, p.Gross(), 68242)
	assertMoneyValue(t, p.Net(), 56868)
	assertMoneyValue(t, p.Tax(), 11374)

	p, _ = PriceGBP(int64(4656414), 20)
	assertMoneyValue(t, p.Gross(), 4656414)
	assertMoneyValue(t, p.Net(), 3880345)
	assertMoneyValue(t, p.Tax(), 776069)

	p, _ = PriceGBP(uint64(4656414), 20)
	assertMoneyValue(t, p.Gross(), 4656414)
	assertMoneyValue(t, p.Net(), 3880345)
	assertMoneyValue(t, p.Tax(), 776069)
}

func TestPriceGBP(t *testing.T) {
	p, _ := PriceGBP(6425, 20)
	assertMoneyValue(t, p.Gross(), 6425)
	assertMoneyValue(t, p.Net(), 5354)
	assertMoneyValue(t, p.Tax(), 1071)
}

func TestPriceFromSubunitsError(t *testing.T) {
	_, err := PriceFromSubunits("XXX", 1457, nil)
	if err == nil {
		t.Errorf("PriceFromSubunits failed to error on code 'XXX'")
	}

	p, err := PriceFromSubunits("GBP", 1457, nil)
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
	m, _ := MoneyFromSubunits("GBP", 1059, nil)
	p.IncludeTax(m, "VAT")

	assertMoneyString(t, p.Gross(), "GBP", "£10.59")
	assertMoneyString(t, p.Net(), "GBP", "£0.00")
	assertMoneyString(t, p.Tax(), "GBP", "£10.59")
}

func TestPriceFromStringError(t *testing.T) {
	_, err := PriceFromString("XXX", "£14.57", nil)
	if err == nil {
		t.Errorf("PriceFromString failed to error on code 'XXX'")
	}

	p, err := PriceFromString("GBP", "£14.57", nil)
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
	m, _ := MoneyFromSubunits("GBP", 1059, nil)
	p.IncludeTax(m, "VAT")

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
	assertJSON(t, bytes, `{"currency":"GBP","gross":"£10.99","net":"£9.16","tax":{"total":"£1.83","detail":[{"amount":"£1.83","description":"VAT"}]}}`)

	bytes, _ = json.Marshal(resp)
	assertJSON(t, bytes, `{"name":"Widget","price":{"currency":"GBP","gross":"£10.99","net":"£9.16","tax":{"total":"£1.83","detail":[{"amount":"£1.83","description":"VAT"}]}}}`)
}

func TestPriceString(t *testing.T) {
	price, _ := PriceGBP(1099, 20)
	assertPriceString(t, price, "GBP", "£10.99")
	assertPriceStringNoSymbol(t, price, "GBP", "10.99")
}

func TestAddTaxMoney(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 5500, nil)
	m, _ := MoneyFromSubunits("GBP", 825, nil)
	p.AddTax(m, "VAT")

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

func TestIncludeTaxMoney(t *testing.T) {
	p, _ := PriceFromSubunits("GBP", 5500, nil)
	m, _ := MoneyFromSubunits("GBP", 825, nil)
	p.IncludeTax(m, "VAT")

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

func TestAddPriceAndImmutability(t *testing.T) {
	p1, _ := PriceFromSubunits("GBP", 2083, nil)
	p1.AddTaxPercent(15, "VAT")
	p1.AddTaxPercent(5, "Small order")

	assertMoneyValue(t, p1.Gross(), 2515)
	assertMoneyValue(t, p1.Net(), 2083)
	assertMoneyValue(t, p1.Tax(), 432)

	bytes, _ := json.Marshal(p1)
	assertJSON(t, bytes, `{"currency":"GBP","gross":"£25.15","net":"£20.83","tax":{"total":"£4.32","detail":[{"amount":"£1.20","description":"Small order"},{"amount":"£3.12","description":"VAT"}]}}`)

	p2, _ := PriceFromSubunits("GBP", 1545, nil)
	p2.AddTaxPercent(15, "VAT")

	p3 := p1.Add(p2)
	assertMoneyValue(t, p3.Gross(), 4292)
	assertMoneyValue(t, p3.Net(), 3628)
	assertMoneyValue(t, p3.Tax(), 664)

	bytes, _ = json.Marshal(p3)
	assertJSON(t, bytes, `{"currency":"GBP","gross":"£42.92","net":"£36.28","tax":{"total":"£6.64","detail":[{"amount":"£1.20","description":"Small order"},{"amount":"£5.44","description":"VAT"}]}}`)

	assertMoneyValue(t, p1.Gross(), 2515)
	assertMoneyValue(t, p1.Net(), 2083)
	assertMoneyValue(t, p1.Tax(), 432)

	bytes, _ = json.Marshal(p1)
	assertJSON(t, bytes, `{"currency":"GBP","gross":"£25.15","net":"£20.83","tax":{"total":"£4.32","detail":[{"amount":"£1.20","description":"Small order"},{"amount":"£3.12","description":"VAT"}]}}`)
}

func TestMulPriceAndImmutability(t *testing.T) {
	p1, _ := PriceFromSubunits("GBP", 2083, nil)
	p1.AddTaxPercent(15, "VAT")
	p1.AddTaxPercent(5, "Small order")

	assertMoneyValue(t, p1.Gross(), 2515)
	assertMoneyValue(t, p1.Net(), 2083)
	assertMoneyValue(t, p1.Tax(), 432)

	bytes, _ := json.Marshal(p1)
	assertJSON(t, bytes, `{"currency":"GBP","gross":"£25.15","net":"£20.83","tax":{"total":"£4.32","detail":[{"amount":"£1.20","description":"Small order"},{"amount":"£3.12","description":"VAT"}]}}`)

	p2 := p1.Mul(3)
	assertMoneyValue(t, p2.Gross(), 7545)
	assertMoneyValue(t, p2.Net(), 6249)
	assertMoneyValue(t, p2.Tax(), 1296)

	bytes, _ = json.Marshal(p2)
	assertJSON(t, bytes, `{"currency":"GBP","gross":"£75.45","net":"£62.49","tax":{"total":"£12.96","detail":[{"amount":"£3.60","description":"Small order"},{"amount":"£9.36","description":"VAT"}]}}`)

	assertMoneyValue(t, p1.Gross(), 2515)
	assertMoneyValue(t, p1.Net(), 2083)
	assertMoneyValue(t, p1.Tax(), 432)

	bytes, _ = json.Marshal(p1)
	assertJSON(t, bytes, `{"currency":"GBP","gross":"£25.15","net":"£20.83","tax":{"total":"£4.32","detail":[{"amount":"£1.20","description":"Small order"},{"amount":"£3.12","description":"VAT"}]}}`)
}
