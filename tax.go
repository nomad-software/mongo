package mongo

// Taxes is a structure that holds the taxes information of a price.
type taxes struct {
	Total  Money `json:"total"`  // The total tax.
	Detail []tax `json:"detail"` // The breakdown of individual taxes.
}

// Tax is a structure that holds tax information.
type tax struct {
	Amount      Money  `json:"amount"`      // The tax amount.
	Description string `json:"description"` // The description of the tax.
}

// Add adds a new tax to the taxes collection.
func (t *taxes) Add(m Money, desc string) {
	t.Detail = append(t.Detail, tax{Amount: m, Description: desc})
	t.Total = t.Total.Add(m)
}
