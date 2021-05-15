package mongo

import (
	"testing"
)

func TestJPYString(t *testing.T) {
	m, _ := FromSubunits("JPY", 1, RoundHalfUp)
	assertMoneyString(t, m, "JPY", "¥1")
	assertMoneyStringNoSymbol(t, m, "JPY", "1")

	m, _ = FromSubunits("JPY", 12, RoundHalfUp)
	assertMoneyString(t, m, "JPY", "¥12")
	assertMoneyStringNoSymbol(t, m, "JPY", "12")

	m, _ = FromSubunits("JPY", 123, RoundHalfUp)
	assertMoneyString(t, m, "JPY", "¥123")
	assertMoneyStringNoSymbol(t, m, "JPY", "123")

	m, _ = FromSubunits("JPY", 1234, RoundHalfUp)
	assertMoneyString(t, m, "JPY", "¥1,234")
	assertMoneyStringNoSymbol(t, m, "JPY", "1,234")

	m, _ = FromSubunits("JPY", 12345, RoundHalfUp)
	assertMoneyString(t, m, "JPY", "¥12,345")
	assertMoneyStringNoSymbol(t, m, "JPY", "12,345")

	m, _ = FromSubunits("JPY", 123456, RoundHalfUp)
	assertMoneyString(t, m, "JPY", "¥123,456")
	assertMoneyStringNoSymbol(t, m, "JPY", "123,456")

	m, _ = FromSubunits("JPY", 1234567, RoundHalfUp)
	assertMoneyString(t, m, "JPY", "¥1,234,567")
	assertMoneyStringNoSymbol(t, m, "JPY", "1,234,567")

	m, _ = FromSubunits("JPY", 12345678, RoundHalfUp)
	assertMoneyString(t, m, "JPY", "¥12,345,678")
	assertMoneyStringNoSymbol(t, m, "JPY", "12,345,678")

	m, _ = FromSubunits("JPY", 123456789, RoundHalfUp)
	assertMoneyString(t, m, "JPY", "¥123,456,789")
	assertMoneyStringNoSymbol(t, m, "JPY", "123,456,789")
}

func TestUSDString(t *testing.T) {
	m, _ := FromSubunits("USD", 1, RoundHalfUp)
	assertMoneyString(t, m, "USD", "$0.01")
	assertMoneyStringNoSymbol(t, m, "USD", "0.01")

	m, _ = FromSubunits("USD", 12, RoundHalfUp)
	assertMoneyString(t, m, "USD", "$0.12")
	assertMoneyStringNoSymbol(t, m, "USD", "0.12")

	m, _ = FromSubunits("USD", 123, RoundHalfUp)
	assertMoneyString(t, m, "USD", "$1.23")
	assertMoneyStringNoSymbol(t, m, "USD", "1.23")

	m, _ = FromSubunits("USD", 1234, RoundHalfUp)
	assertMoneyString(t, m, "USD", "$12.34")
	assertMoneyStringNoSymbol(t, m, "USD", "12.34")

	m, _ = FromSubunits("USD", 12345, RoundHalfUp)
	assertMoneyString(t, m, "USD", "$123.45")
	assertMoneyStringNoSymbol(t, m, "USD", "123.45")

	m, _ = FromSubunits("USD", 123456, RoundHalfUp)
	assertMoneyString(t, m, "USD", "$1,234.56")
	assertMoneyStringNoSymbol(t, m, "USD", "1,234.56")

	m, _ = FromSubunits("USD", 1234567, RoundHalfUp)
	assertMoneyString(t, m, "USD", "$12,345.67")
	assertMoneyStringNoSymbol(t, m, "USD", "12,345.67")

	m, _ = FromSubunits("USD", 12345678, RoundHalfUp)
	assertMoneyString(t, m, "USD", "$123,456.78")
	assertMoneyStringNoSymbol(t, m, "USD", "123,456.78")

	m, _ = FromSubunits("USD", 123456789, RoundHalfUp)
	assertMoneyString(t, m, "USD", "$1,234,567.89")
	assertMoneyStringNoSymbol(t, m, "USD", "1,234,567.89")
}

func TestBYNString(t *testing.T) {
	m, _ := FromSubunits("BYN", 1, RoundHalfUp)
	assertMoneyString(t, m, "BYN", "0,01 p.")
	assertMoneyStringNoSymbol(t, m, "BYN", "0,01")

	m, _ = FromSubunits("BYN", 12, RoundHalfUp)
	assertMoneyString(t, m, "BYN", "0,12 p.")
	assertMoneyStringNoSymbol(t, m, "BYN", "0,12")

	m, _ = FromSubunits("BYN", 123, RoundHalfUp)
	assertMoneyString(t, m, "BYN", "1,23 p.")
	assertMoneyStringNoSymbol(t, m, "BYN", "1,23")

	m, _ = FromSubunits("BYN", 1234, RoundHalfUp)
	assertMoneyString(t, m, "BYN", "12,34 p.")
	assertMoneyStringNoSymbol(t, m, "BYN", "12,34")

	m, _ = FromSubunits("BYN", 12345, RoundHalfUp)
	assertMoneyString(t, m, "BYN", "123,45 p.")
	assertMoneyStringNoSymbol(t, m, "BYN", "123,45")

	m, _ = FromSubunits("BYN", 123456, RoundHalfUp)
	assertMoneyString(t, m, "BYN", "1 234,56 p.")
	assertMoneyStringNoSymbol(t, m, "BYN", "1 234,56")

	m, _ = FromSubunits("BYN", 1234567, RoundHalfUp)
	assertMoneyString(t, m, "BYN", "12 345,67 p.")
	assertMoneyStringNoSymbol(t, m, "BYN", "12 345,67")

	m, _ = FromSubunits("BYN", 12345678, RoundHalfUp)
	assertMoneyString(t, m, "BYN", "123 456,78 p.")
	assertMoneyStringNoSymbol(t, m, "BYN", "123 456,78")

	m, _ = FromSubunits("BYN", 123456789, RoundHalfUp)
	assertMoneyString(t, m, "BYN", "1 234 567,89 p.")
	assertMoneyStringNoSymbol(t, m, "BYN", "1 234 567,89")
}

func TestBHDString(t *testing.T) {
	m, _ := FromSubunits("BHD", 1, RoundHalfUp)
	assertMoneyString(t, m, "BHD", "0.001 .د.ب ")
	assertMoneyStringNoSymbol(t, m, "BHD", "0.001")

	m, _ = FromSubunits("BHD", 12, RoundHalfUp)
	assertMoneyString(t, m, "BHD", "0.012 .د.ب ")
	assertMoneyStringNoSymbol(t, m, "BHD", "0.012")

	m, _ = FromSubunits("BHD", 123, RoundHalfUp)
	assertMoneyString(t, m, "BHD", "0.123 .د.ب ")
	assertMoneyStringNoSymbol(t, m, "BHD", "0.123")

	m, _ = FromSubunits("BHD", 1234, RoundHalfUp)
	assertMoneyString(t, m, "BHD", "1.234 .د.ب ")
	assertMoneyStringNoSymbol(t, m, "BHD", "1.234")

	m, _ = FromSubunits("BHD", 12345, RoundHalfUp)
	assertMoneyString(t, m, "BHD", "12.345 .د.ب ")
	assertMoneyStringNoSymbol(t, m, "BHD", "12.345")

	m, _ = FromSubunits("BHD", 123456, RoundHalfUp)
	assertMoneyString(t, m, "BHD", "123.456 .د.ب ")
	assertMoneyStringNoSymbol(t, m, "BHD", "123.456")

	m, _ = FromSubunits("BHD", 1234567, RoundHalfUp)
	assertMoneyString(t, m, "BHD", "1,234.567 .د.ب ")
	assertMoneyStringNoSymbol(t, m, "BHD", "1,234.567")

	m, _ = FromSubunits("BHD", 12345678, RoundHalfUp)
	assertMoneyString(t, m, "BHD", "12,345.678 .د.ب ")
	assertMoneyStringNoSymbol(t, m, "BHD", "12,345.678")

	m, _ = FromSubunits("BHD", 123456789, RoundHalfUp)
	assertMoneyString(t, m, "BHD", "123,456.789 .د.ب ")
	assertMoneyStringNoSymbol(t, m, "BHD", "123,456.789")
}

func TestCLFString(t *testing.T) {
	m, _ := FromSubunits("CLF", 1, RoundHalfUp)
	assertMoneyString(t, m, "CLF", "UF0,0001")
	assertMoneyStringNoSymbol(t, m, "CLF", "0,0001")

	m, _ = FromSubunits("CLF", 12, RoundHalfUp)
	assertMoneyString(t, m, "CLF", "UF0,0012")
	assertMoneyStringNoSymbol(t, m, "CLF", "0,0012")

	m, _ = FromSubunits("CLF", 123, RoundHalfUp)
	assertMoneyString(t, m, "CLF", "UF0,0123")
	assertMoneyStringNoSymbol(t, m, "CLF", "0,0123")

	m, _ = FromSubunits("CLF", 1234, RoundHalfUp)
	assertMoneyString(t, m, "CLF", "UF0,1234")
	assertMoneyStringNoSymbol(t, m, "CLF", "0,1234")

	m, _ = FromSubunits("CLF", 12345, RoundHalfUp)
	assertMoneyString(t, m, "CLF", "UF1,2345")
	assertMoneyStringNoSymbol(t, m, "CLF", "1,2345")

	m, _ = FromSubunits("CLF", 123456, RoundHalfUp)
	assertMoneyString(t, m, "CLF", "UF12,3456")
	assertMoneyStringNoSymbol(t, m, "CLF", "12,3456")

	m, _ = FromSubunits("CLF", 1234567, RoundHalfUp)
	assertMoneyString(t, m, "CLF", "UF123,4567")
	assertMoneyStringNoSymbol(t, m, "CLF", "123,4567")

	m, _ = FromSubunits("CLF", 12345678, RoundHalfUp)
	assertMoneyString(t, m, "CLF", "UF1.234,5678")
	assertMoneyStringNoSymbol(t, m, "CLF", "1.234,5678")

	m, _ = FromSubunits("CLF", 123456789, RoundHalfUp)
	assertMoneyString(t, m, "CLF", "UF12.345,6789")
	assertMoneyStringNoSymbol(t, m, "CLF", "12.345,6789")
}
