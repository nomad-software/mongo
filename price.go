package mongo

import "fmt"

// Price is a structure that holds a price and gives information about the
// amount of tax applied to that price.
type Price struct {
	Gross Money // The gross.
	Net   Money // The net which is equal to the gross minus tax.
	Tax   Money // The amount of tax subtracted from the gross to produce the net.
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

	price.Gross, err = MoneyFromSubunits(currIsoCode, grossValue, f)
	if err != nil {
		return Price{}, err
	}

	if taxPercent < 0.0 || taxPercent > 100.0 {
		return Price{}, fmt.Errorf("tax percent '%f' must be between 1 and 100", taxPercent)
	}

	if taxPercent < 100.0 {
		price.Net = price.Gross.Div(1 + (taxPercent / 100))
		price.Tax = price.Gross.Sub(price.Net)
	} else {
		price.Tax = price.Gross
		price.Net = price.Gross.Sub(price.Tax)
	}

	return price, nil
}
