package mongo

import (
	"fmt"
)

// Price is a structure that holds a price and gives information about the
// amount of tax applied to that price.
type Price struct {
	gross      Money   // The gross.
	net        Money   // The net which is equal to the gross minus tax.
	tax        Money   // The amount of tax subtracted from the gross to produce the net.
	taxPercent float64 // The percentage of tax deducted.
}

// PriceFromSubunits constructs a new price object from an integer and tax
// percentage. The value integer used should represent the subunits of the
// currency.
// currIsoCode is an ISO 4217 currency code.
// value is monetary value in subunits.
// taxPercent is the amount of tax applied to this price.
// roundFunc is a function to be used for division operations.
func PriceFromSubunits(currIsoCode string, grossValue int64, taxPercent float64, f roundFunc) (Price, error) {
	var price Price
	var err error

	price.gross, err = MoneyFromSubunits(currIsoCode, grossValue, f)
	if err != nil {
		return Price{}, err
	}

	price.taxPercent = taxPercent

	if taxPercent < 0.0 || taxPercent > 100.0 {
		return Price{}, fmt.Errorf("tax percent '%f' must be between 1 and 100", taxPercent)
	}

	if taxPercent < 100.0 {
		price.net = price.gross.Div(1 + (taxPercent / 100))
		price.tax = price.gross.Sub(price.net)
	} else {
		price.tax = price.gross
		price.net = price.gross.Sub(price.tax)
	}

	return price, nil
}

// MoneyFromString constructs a new price object from a string and tax
// percentage. Everything not contained within a number is stripped out before
// parsing.
// currIsoCode is an ISO 4217 currency code.
// value is monetary value in subunits.
// taxPercent is the amount of tax applied to this price.
// roundFunc is a function to be used for division operations.
func PriceFromString(currIsoCode string, grossValueStr string, taxPercent float64, f roundFunc) (Price, error) {
	var price Price
	var err error

	price.gross, err = MoneyFromString(currIsoCode, grossValueStr, f)
	if err != nil {
		return Price{}, err
	}

	price.taxPercent = taxPercent

	if taxPercent < 0.0 || taxPercent > 100.0 {
		return Price{}, fmt.Errorf("tax percent '%f' must be between 1 and 100", taxPercent)
	}

	if taxPercent < 100.0 {
		price.net = price.gross.Div(1 + (taxPercent / 100))
		price.tax = price.gross.Sub(price.net)
	} else {
		price.tax = price.gross
		price.net = price.gross.Sub(price.tax)
	}

	return price, nil
}

// PriceGBP is a helper function.
func PriceGBP(value int64, taxPercent float64) (Price, error) {
	return PriceFromSubunits("GBP", value, taxPercent, nil)
}

// PriceEUR is a helper function.
func PriceEUR(value int64, taxPercent float64) (Price, error) {
	return PriceFromSubunits("EUR", value, taxPercent, nil)
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
	return p.net
}

// Tax returns the amount of tax subtracted from the gross to produce the net.
func (p Price) Tax() Money {
	return p.tax
}

// TaxPercent returns the amount of tax deducted on creation of the price as a
// percentage.
func (p Price) TaxPercent() float64 {
	return p.taxPercent
}

// MarshalJSON is an implementation of json.Marshaller.
func (p Price) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf(`{"currency": "%s", "gross": "%s", "net": "%s", "tax": "%s", "taxPercent": %f}`, p.gross.format.code, p.gross, p.net, p.tax, p.taxPercent)
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
