package mongo

// assertSameMoneyCurrency will panic if the arguments are money objects
// containing different currencies.
func assertSameMoneyCurrency(a, b Money) {
	if a.format.code != b.format.code {
		panic("Failed to perform operation on different currencies")
	}
}

// assertSamePriceCurrency will panic if the arguments are price objects
// containing different currencies.
func assertSamePriceCurrency(a, b Price) {
	if a.Gross.format.code != b.Gross.format.code {
		panic("Failed to perform operation on different currencies")
	}

	if a.Net.format.code != b.Net.format.code {
		panic("Failed to perform operation on different currencies")
	}

	if a.Tax.format.code != b.Tax.format.code {
		panic("Failed to perform operation on different currencies")
	}
}
