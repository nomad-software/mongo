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
		Total:  t1,
		Detail: make([]tax, 0),
	}

	taxes.Add(t2, "VAT")
	taxes.Add(t3, "VAT")

	assertMoneyValue(t, taxes.Total, 1844)
	assertMoneyValue(t, taxes.Detail[0].Amount, 1387)
	assertMoneyValue(t, taxes.Detail[1].Amount, 457)
}

func TestTaxJsonMarshalling(t *testing.T) {
	t1, _ := MoneyGBP(0)
	t2, _ := MoneyGBP(1387)

	taxes := taxes{
		Total:  t1,
		Detail: make([]tax, 0),
	}

	taxes.Add(t2, "VAT")

	bytes, _ := json.Marshal(taxes)
	assertJSON(t, bytes, `{"total":{"currency":"GBP","formatted":"£13.87"},"detail":[{"amount":{"currency":"GBP","formatted":"£13.87"},"description":"VAT"}]}`)
}
