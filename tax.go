package mongo

import (
	"encoding/json"
	"fmt"
)

type taxDetail []tax

// Taxes is a structure that holds the taxes information of a price.
type taxes struct {
	total  Money     // The total tax.
	detail taxDetail // The breakdown of individual taxes.
}

// Add adds a new tax to the taxes collection.
func (t *taxes) Add(m Money, desc string) {
	t.detail = append(t.detail, tax{amount: m, desc: desc})
	t.total = t.total.Add(m)
}

// MarshalJSON is an implementation of json.Marshaller.
func (t taxes) MarshalJSON() ([]byte, error) {
	detail, err := json.Marshal(t.detail)
	if err != nil {
		return []byte{}, err
	}

	json := fmt.Sprintf(
		`{"formatted": "%s", "detail": %s}`,
		t.total.String(),
		detail,
	)

	return []byte(json), nil
}

// Tax is a structure that holds tax information.
type tax struct {
	amount Money  // The tax amount.
	desc   string // The description of the tax.
}

// MarshalJSON is an implementation of json.Marshaller.
func (t tax) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf(
		`{"formatted": "%s", "description": "%s"}`,
		t.amount,
		t.desc,
	)

	return []byte(json), nil
}
