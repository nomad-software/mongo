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
		detail: make([]tax, 0),
	}

	taxes.Add(t2, "VAT")
	taxes.Add(t3, "VAT")

	assertMoneyValue(t, taxes.total, 1844)
	assertMoneyValue(t, taxes.detail[0].amount, 1387)
	assertMoneyValue(t, taxes.detail[1].amount, 457)
}

func TestTaxJsonMarshalling(t *testing.T) {
	t1, _ := MoneyGBP(0)
	t2, _ := MoneyGBP(1387)

	taxes := taxes{
		total:  t1,
		detail: make([]tax, 0),
	}

	taxes.Add(t2, "VAT")

	bytes, _ := json.Marshal(taxes)
	assertJSON(t, bytes, `{"formatted":"£13.87","detail":[{"formatted":"£13.87","description":"VAT"}]}`)
}
