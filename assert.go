package mongo

// AssertSameCurrency will panic if the arguments are money objects containing
// different currencies.
func assertSameCurrency(a, b Money) {
	if a.format.code != b.format.code {
		panic("Failed to perform operation on different currencies")
	}
}
