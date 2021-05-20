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
	if a.gross.format.code != b.gross.format.code {
		panic("Failed to perform operation on different currencies")
	}

	if a.net.format.code != b.net.format.code {
		panic("Failed to perform operation on different currencies")
	}

	if a.tax.format.code != b.tax.format.code {
		panic("Failed to perform operation on different currencies")
	}
}
