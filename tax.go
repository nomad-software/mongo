package mongo

import (
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type detail map[string]Money

// Adds a new tax to the detail collection.
// If the tax description exists the money value will be added to it.
// This method returns a new map to make sure this operation is immutable
// across price types.
func (d detail) add(desc string, m Money) detail {
	result := make(map[string]Money, 0)

	for k, v := range d {
		result[k] = v
	}

	if _, ok := result[desc]; ok {
		result[desc] = result[desc].Add(m)
	} else {
		result[desc] = m
	}

	return result
}

// Mul multiplies all the values in the tax details collection by the passed
// amount.
func (d detail) mul(n int64) detail {
	result := make(map[string]Money, 0)

	for k, v := range d {
		result[k] = v.Mul(n)
	}

	return result
}

// MarshalJSON is an implementation of json.Marshaller.
func (d detail) MarshalJSON() ([]byte, error) {
	json := make([]string, 0)

	for k, v := range d {
		json = append(json, fmt.Sprintf(`{"description": "%s", "formatted": "%s"}`, k, v.String()))
	}

	// Because the json map's order is non-deterministic, sort for deterministic output.
	slices.Sort(json)

	return []byte("[" + strings.Join(json, ",") + "]"), nil
}

// Taxes is a structure that holds the taxes information of a price.
type taxes struct {
	total  Money  // The total tax.
	detail detail // The breakdown of individual taxes.
}

// Add adds a new tax to the taxes collection.
func (t taxes) add(desc string, m Money) taxes {
	t.detail = t.detail.add(desc, m)
	t.total = t.total.Add(m)
	return t
}

// Mul multiplies taxes in the taxes collection.
func (t taxes) mul(n int64) taxes {
	t.detail = t.detail.mul(n)
	t.total = t.total.Mul(n)
	return t
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
