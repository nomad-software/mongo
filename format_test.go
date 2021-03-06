package mongo

import (
	"testing"
)

func TestJPYString(t *testing.T) {
	m, _ := FromSubunits("JPY", 1, roundHalfUp)
	assertMoneyFormat(t, m, "JPY", "¥1")

	m, _ = FromSubunits("JPY", 12, roundHalfUp)
	assertMoneyFormat(t, m, "JPY", "¥12")

	m, _ = FromSubunits("JPY", 123, roundHalfUp)
	assertMoneyFormat(t, m, "JPY", "¥123")

	m, _ = FromSubunits("JPY", 1234, roundHalfUp)
	assertMoneyFormat(t, m, "JPY", "¥1,234")

	m, _ = FromSubunits("JPY", 12345, roundHalfUp)
	assertMoneyFormat(t, m, "JPY", "¥12,345")

	m, _ = FromSubunits("JPY", 123456, roundHalfUp)
	assertMoneyFormat(t, m, "JPY", "¥123,456")

	m, _ = FromSubunits("JPY", 1234567, roundHalfUp)
	assertMoneyFormat(t, m, "JPY", "¥1,234,567")

	m, _ = FromSubunits("JPY", 12345678, roundHalfUp)
	assertMoneyFormat(t, m, "JPY", "¥12,345,678")

	m, _ = FromSubunits("JPY", 123456789, roundHalfUp)
	assertMoneyFormat(t, m, "JPY", "¥123,456,789")
}

func TestUSDString(t *testing.T) {
	m, _ := FromSubunits("USD", 1, roundHalfUp)
	assertMoneyFormat(t, m, "USD", "$0.01")

	m, _ = FromSubunits("USD", 12, roundHalfUp)
	assertMoneyFormat(t, m, "USD", "$0.12")

	m, _ = FromSubunits("USD", 123, roundHalfUp)
	assertMoneyFormat(t, m, "USD", "$1.23")

	m, _ = FromSubunits("USD", 1234, roundHalfUp)
	assertMoneyFormat(t, m, "USD", "$12.34")

	m, _ = FromSubunits("USD", 12345, roundHalfUp)
	assertMoneyFormat(t, m, "USD", "$123.45")

	m, _ = FromSubunits("USD", 123456, roundHalfUp)
	assertMoneyFormat(t, m, "USD", "$1,234.56")

	m, _ = FromSubunits("USD", 1234567, roundHalfUp)
	assertMoneyFormat(t, m, "USD", "$12,345.67")

	m, _ = FromSubunits("USD", 12345678, roundHalfUp)
	assertMoneyFormat(t, m, "USD", "$123,456.78")

	m, _ = FromSubunits("USD", 123456789, roundHalfUp)
	assertMoneyFormat(t, m, "USD", "$1,234,567.89")
}

func TestBYNString(t *testing.T) {
	m, _ := FromSubunits("BYN", 1, roundHalfUp)
	assertMoneyFormat(t, m, "BYN", "0,01 p.")

	m, _ = FromSubunits("BYN", 12, roundHalfUp)
	assertMoneyFormat(t, m, "BYN", "0,12 p.")

	m, _ = FromSubunits("BYN", 123, roundHalfUp)
	assertMoneyFormat(t, m, "BYN", "1,23 p.")

	m, _ = FromSubunits("BYN", 1234, roundHalfUp)
	assertMoneyFormat(t, m, "BYN", "12,34 p.")

	m, _ = FromSubunits("BYN", 12345, roundHalfUp)
	assertMoneyFormat(t, m, "BYN", "123,45 p.")

	m, _ = FromSubunits("BYN", 123456, roundHalfUp)
	assertMoneyFormat(t, m, "BYN", "1 234,56 p.")

	m, _ = FromSubunits("BYN", 1234567, roundHalfUp)
	assertMoneyFormat(t, m, "BYN", "12 345,67 p.")

	m, _ = FromSubunits("BYN", 12345678, roundHalfUp)
	assertMoneyFormat(t, m, "BYN", "123 456,78 p.")

	m, _ = FromSubunits("BYN", 123456789, roundHalfUp)
	assertMoneyFormat(t, m, "BYN", "1 234 567,89 p.")
}

func TestBHDString(t *testing.T) {
	m, _ := FromSubunits("BHD", 1, roundHalfUp)
	assertMoneyFormat(t, m, "BHD", "0.001 .د.ب ")

	m, _ = FromSubunits("BHD", 12, roundHalfUp)
	assertMoneyFormat(t, m, "BHD", "0.012 .د.ب ")

	m, _ = FromSubunits("BHD", 123, roundHalfUp)
	assertMoneyFormat(t, m, "BHD", "0.123 .د.ب ")

	m, _ = FromSubunits("BHD", 1234, roundHalfUp)
	assertMoneyFormat(t, m, "BHD", "1.234 .د.ب ")

	m, _ = FromSubunits("BHD", 12345, roundHalfUp)
	assertMoneyFormat(t, m, "BHD", "12.345 .د.ب ")

	m, _ = FromSubunits("BHD", 123456, roundHalfUp)
	assertMoneyFormat(t, m, "BHD", "123.456 .د.ب ")

	m, _ = FromSubunits("BHD", 1234567, roundHalfUp)
	assertMoneyFormat(t, m, "BHD", "1,234.567 .د.ب ")

	m, _ = FromSubunits("BHD", 12345678, roundHalfUp)
	assertMoneyFormat(t, m, "BHD", "12,345.678 .د.ب ")

	m, _ = FromSubunits("BHD", 123456789, roundHalfUp)
	assertMoneyFormat(t, m, "BHD", "123,456.789 .د.ب ")
}

func TestCLFString(t *testing.T) {
	m, _ := FromSubunits("CLF", 1, roundHalfUp)
	assertMoneyFormat(t, m, "CLF", "UF0,0001")

	m, _ = FromSubunits("CLF", 12, roundHalfUp)
	assertMoneyFormat(t, m, "CLF", "UF0,0012")

	m, _ = FromSubunits("CLF", 123, roundHalfUp)
	assertMoneyFormat(t, m, "CLF", "UF0,0123")

	m, _ = FromSubunits("CLF", 1234, roundHalfUp)
	assertMoneyFormat(t, m, "CLF", "UF0,1234")

	m, _ = FromSubunits("CLF", 12345, roundHalfUp)
	assertMoneyFormat(t, m, "CLF", "UF1,2345")

	m, _ = FromSubunits("CLF", 123456, roundHalfUp)
	assertMoneyFormat(t, m, "CLF", "UF12,3456")

	m, _ = FromSubunits("CLF", 1234567, roundHalfUp)
	assertMoneyFormat(t, m, "CLF", "UF123,4567")

	m, _ = FromSubunits("CLF", 12345678, roundHalfUp)
	assertMoneyFormat(t, m, "CLF", "UF1.234,5678")

	m, _ = FromSubunits("CLF", 123456789, roundHalfUp)
	assertMoneyFormat(t, m, "CLF", "UF12.345,6789")
}
