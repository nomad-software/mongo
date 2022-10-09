package mongo

import (
	"encoding/json"
	"testing"
)

func TestAddTax(t *testing.T) {
	t1, _ := MoneyGBP(0)
	t2, _ := MoneyGBP(1387)
	t3, _ := MoneyGBP(457)

	taxes := taxes{
		total:  t1,
		detail: make(map[string]Money, 0),
	}

	taxes = taxes.add("VAT", t2)
	taxes = taxes.add("VAT", t3)

	assertMoneyValue(t, taxes.total, 1844)
	assertMoneyValue(t, taxes.detail["VAT"], 1844)
}

func TestTaxJsonMarshalling(t *testing.T) {
	t1, _ := MoneyGBP(0)
	t2, _ := MoneyGBP(1387)

	taxes := taxes{
		total:  t1,
		detail: make(map[string]Money, 0),
	}

	taxes = taxes.add("VAT", t2)

	bytes, _ := json.Marshal(taxes)
	assertJSON(t, bytes, `{"formatted":"£13.87","detail":[{"description":"VAT","formatted":"£13.87"}]}`)
}
