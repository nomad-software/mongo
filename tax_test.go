package mongo

import (
	"encoding/json"
	"testing"
)

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
