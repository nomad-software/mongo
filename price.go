package mongo

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/constraints"
)

// Price is a structure that holds a price and gives information about the
// amount of tax applied to that price.
type Price struct {
	gross Money // The gross.
	taxes taxes // The amount of tax subtracted from the gross to produce the net.
}

// PriceFromSubunits constructs a new price object from an integer. The value
// integer used should represent the subunits of the currency.
// currIsoCode is an ISO 4217 currency code.
// gross is monetary value in subunits.
// roundFunc is a function to be used for division operations.
func PriceFromSubunits[T constraints.Integer](currIsoCode string, gross T, f roundFunc) (Price, error) {
	var price Price
	var err error

	price.gross, err = MoneyFromSubunits(currIsoCode, gross, f)
	if err != nil {
		return Price{}, err
	}

	price.taxes = taxes{
		total:  price.gross.Clone(0),
		detail: make(map[string]Money, 0),
	}

	return price, nil
}

// PriceFromSubunits constructs a new price object from an integer. The value
// integer used should represent the subunits of the currency.
// currIsoCode is an ISO 4217 currency code.
// gross is monetary value expressed as a float.
// roundFunc is a function to be used for division operations.
func PriceFromFloat[T constraints.Float](currIsoCode string, gross T, f roundFunc) (Price, error) {
	var price Price
	var err error

	price.gross, err = MoneyFromFloat(currIsoCode, gross, f)
	if err != nil {
		return Price{}, err
	}

	price.taxes = taxes{
		total:  price.gross.Clone(0),
		detail: make(map[string]Money, 0),
	}

	return price, nil
}

// MoneyFromString constructs a new price object from a string. Everything not
// contained within a number is stripped out before parsing.
// currIsoCode is an ISO 4217 currency code.
// gross is monetary value in subunits.
// roundFunc is a function to be used for division operations.
func PriceFromString(currIsoCode string, gross string, f roundFunc) (Price, error) {
	var price Price
	var err error

	price.gross, err = MoneyFromString(currIsoCode, gross, f)
	if err != nil {
		return Price{}, err
	}

	price.taxes = taxes{
		total:  price.gross.Clone(0),
		detail: make(map[string]Money, 0),
	}

	return price, nil
}

// PriceGBP is a helper function.
// gross is the gross monetary value in subunits.
// vat is a tax percentage that's included in the gross value.
func PriceGBP[T constraints.Integer](gross T, vat float64) (Price, error) {
	price, err := PriceFromSubunits("GBP", gross, nil)
	if err != nil {
		return Price{}, err
	}
	price.IncludeTaxPercent(vat, "VAT")

	return price, err
}

// AddTax adds a tax to the price using a money value.
// This will literally add the money amount to the gross price.
func (p *Price) AddTax(m Money, desc string) {
	assertSameMoneyCurrency(p.gross, m)
	t := p.gross.Clone(m.value)
	p.taxes = p.taxes.add(desc, t)
	p.gross = p.gross.Add(t)
}

// AddTaxPercentage adds a tax to the price using a percentage.
// This will literally add a percentage to the gross price.
func (p *Price) AddTaxPercent(percent float64, desc string) {
	v := (float64(p.gross.value) / 100) * percent
	t := p.gross.Clone(p.gross.round(v))
	p.taxes = p.taxes.add(desc, t)
	p.gross = p.gross.Add(t)
}

// IncludeTax adds a tax to the price using a money value.
// This implies this tax is already included in the gross price.
func (p *Price) IncludeTax(m Money, desc string) {
	assertSameMoneyCurrency(p.gross, m)
	t := p.gross.Clone(m.value)
	p.taxes = p.taxes.add(desc, t)
}

// IncludeTaxPercent adds a tax to the price using a percentage.
// This implies this tax is already included in the gross price.
func (p *Price) IncludeTaxPercent(percent float64, desc string) {
	t := p.Net().Sub(p.Net().Div(1 + (percent / 100)))
	p.taxes = p.taxes.add(desc, t)
}

// IsoCode returns the ISO 4217 currency code.
func (p Price) IsoCode() string {
	return p.gross.format.code
}

// Gross returns the gross monetary value of the price.
func (p Price) Gross() Money {
	return p.gross
}

// Net returns the net monetary value of the price which is equal to the gross
// minus tax.
func (p Price) Net() Money {
	return p.gross.Sub(p.taxes.total)
}

// Tax returns the total amount of tax subtracted from the gross to produce the net.
func (p Price) Tax() Money {
	return p.taxes.total
}

// Add is an arithmetic operator.
func (p Price) Add(v Price) Price {
	p.gross = p.gross.Add(v.gross)
	for k, v := range v.taxes.detail {
		p.taxes = p.taxes.add(k, v)
	}
	return p
}

// Mul is an arithmetic operator.
func (p Price) Mul(n int64) Price {
	p.gross = p.gross.Mul(n)
	p.taxes = p.taxes.mul(n)
	return p
}

// MarshalJSON is an implementation of json.Marshaller.
func (p Price) MarshalJSON() ([]byte, error) {
	tax, err := json.Marshal(p.taxes)
	if err != nil {
		return []byte{}, err
	}

	json := fmt.Sprintf(
		`{"currency": "%s", "gross": "%s", "net": "%s", "tax": %s}`,
		p.IsoCode(),
		p.Gross(),
		p.Net(),
		tax,
	)

	return []byte(json), nil
}

// String is an implementation of fmt.Stringer and returns the string
// formatted representation of the price value.
func (p Price) String() string {
	return p.gross.String()
}

// StringNoSymbol returns the string formatted representation of the price
// value without a currency symbol.
func (p Price) StringNoSymbol() string {
	return p.gross.StringNoSymbol()
}
