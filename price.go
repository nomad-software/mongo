package mongo

import (
	"encoding/json"
	"fmt"
)

// Price is a structure that holds a price and gives information about the
// amount of tax applied to that price.
type Price struct {
	gross Money // The gross.
	taxes taxes // The amount of tax subtracted from the gross to produce the net.
}

// PriceFromSubunits constructs a new price object from an integer and tax
// percentage. The value integer used should represent the subunits of the
// currency.
// currIsoCode is an ISO 4217 currency code.
// value is monetary value in subunits.
// roundFunc is a function to be used for division operations.
func PriceFromSubunits(currIsoCode string, grossValue int64, f roundFunc) (Price, error) {
	var price Price
	var err error

	price.gross, err = MoneyFromSubunits(currIsoCode, grossValue, f)
	if err != nil {
		return Price{}, err
	}

	price.taxes = taxes{
		total:  price.gross.Clone(0),
		detail: make([]tax, 0),
	}

	return price, nil
}

// MoneyFromString constructs a new price object from a string and tax
// percentage. Everything not contained within a number is stripped out before
// parsing.
// currIsoCode is an ISO 4217 currency code.
// value is monetary value in subunits.
// roundFunc is a function to be used for division operations.
func PriceFromString(currIsoCode string, grossValueStr string, f roundFunc) (Price, error) {
	var price Price
	var err error

	price.gross, err = MoneyFromString(currIsoCode, grossValueStr, f)
	if err != nil {
		return Price{}, err
	}

	price.taxes = taxes{
		total:  price.gross.Clone(0),
		detail: make([]tax, 0),
	}

	return price, nil
}

// PriceGBP is a helper function.
func PriceGBP(grossValue int64, vat float64) (Price, error) {
	price, err := PriceFromSubunits("GBP", grossValue, nil)
	if err != nil {
		return Price{}, err
	}
	price.IncludeTaxPercent(vat, "VAT")

	return price, err
}

// AddTaxSubunits adds a tax to the price using subunits.
// This will literally add a subunit amount to the gross price.
func (p *Price) AddTaxSubunits(value int64, desc string) {
	t := p.gross.Clone(value)
	p.taxes.Add(t, desc)
	p.gross = p.gross.Add(t)
}

// AddTaxPercentage adds a tax to the price using a percentage.
// This will literally add a percentage to the gross price.
func (p *Price) AddTaxPercent(percent float64, desc string) {
	v := float64(p.gross.value/100) * percent
	t := p.gross.Clone(p.gross.round(v))
	p.taxes.Add(t, desc)
	p.gross = p.gross.Add(t)
}

// IncludeTaxSubunits adds a tax to the price using subunits.
// This implies this tax is already included in the gross price.
func (p *Price) IncludeTaxSubunits(value int64, desc string) {
	t := p.gross.Clone(value)
	p.taxes.Add(t, desc)
}

// AddTaxPercentage adds a tax to the price using a percentage.
// This implies this tax is already included in the gross price.
func (p *Price) IncludeTaxPercent(percent float64, desc string) {
	t := p.Net().Sub(p.Net().Div(1 + (percent / 100)))
	p.taxes.Add(t, desc)
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
