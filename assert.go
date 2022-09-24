package mongo

// assertSameMoneyCurrency will panic if the arguments are money objects
// containing different currencies.
func assertSameMoneyCurrency(a, b Money) {
	if a.format.code != b.format.code {
		panic("Failed to perform operation on different currencies")
	}
}
