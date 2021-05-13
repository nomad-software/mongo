package mongo

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Money is the main structure that holds a monetary value and how to format it
// as a string.
type Money struct {
	format currencyFormat // The currency format object.
	value  int64          // The monetary value as a integer.
	round  roundFunc      // The rounding function to use for division and multiplication.
}

// FromSubunits constructs a new money object from an integer. The integer used
// should contain the subunits of the currency.
// currIsoCode is an ISO 4217 currency code.
// value is monetary value in subunits.
// roundFunc is a function to be used for division operations.
func FromSubunits(currIsoCode string, value int64, f roundFunc) (Money, error) {
	curr, ok := currencyFormats[currIsoCode]
	if !ok {
		return Money{}, fmt.Errorf("the currency code '%s' is not recognised", currIsoCode)
	}
	if f == nil {
		f = roundHalfUp
	}
	m := Money{
		format: curr,
		value:  value,
		round:  f,
	}
	return m, nil
}

// FromString constructs a new money object from a string. Everything not
// contained within a number is stripped out before parsing.
// currIsoCode is an ISO 4217 currency code.
// value is monetary value in subunits.
// roundFunc is a function to be used for division operations.
func FromString(currIsoCode string, str string, f roundFunc) (Money, error) {
	curr, ok := currencyFormats[currIsoCode]
	if !ok {
		return Money{}, fmt.Errorf("the currency code '%s' is not recognised", currIsoCode)
	}
	if f == nil {
		f = roundHalfUp
	}

	isNegative := strings.Contains(str, "-")

	// Remove everything before the first number and after the last number.
	re := regexp.MustCompile("^.*?([0-9].*[0-9]).*$")
	str = re.ReplaceAllString(str, "$1")

	if curr.subunits > 0 {
		// If the string is longer than the amount of subunits in this
		// currency, we expect to see a subunit separator.
		if len(str) > curr.subunits {
			if string(str[len(str)-(curr.subunits+1)]) != curr.subSep {
				return Money{}, fmt.Errorf("failed to parse string to money, no subunits defined")
			}
			str = strings.ReplaceAll(str, curr.subSep, "")
		}
		str = strings.ReplaceAll(str, curr.thouSep, "")
	}

	if isNegative {
		str = "-" + str
	}

	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return Money{}, err
	}

	m := Money{
		format: curr,
		value:  value,
		round:  f,
	}

	return m, nil
}

// GBP is a helper function.
func GBP(value int64) (Money, error) {
	return FromSubunits("GBP", value, nil)
}

// EUR is a helper function.
func EUR(value int64) (Money, error) {
	return FromSubunits("EUR", value, nil)
}

// IsoCode returns the ISO 4217 currency code.
func (m Money) IsoCode() string {
	return m.format.code
}

// Value returns the entire monetary value expressed in subunits.
// For example, using GBP this would be pence, using EUR would be cents.
func (m Money) Value() int64 {
	return m.value
}

// Units returns only the monetary units.
func (m Money) Units() int64 {
	if m.format.subunits == 0 {
		return m.value
	}
	units := [5]int64{0, 10, 100, 1000, 10000}
	return (m.value - m.Subunits()) / units[m.format.subunits]
}

// Subunits returns only the monetary subunits.
func (m Money) Subunits() int64 {
	if m.format.subunits == 0 {
		return 0
	}
	units := [5]int64{0, 10, 100, 1000, 10000}
	return m.value % units[m.format.subunits]
}

// Add is an arithmetic operator.
func (m Money) Add(v Money) Money {
	assertSameCurrency(m, v)
	m.value += v.value
	return m
}

// Sub is an arithmetic operator.
func (m Money) Sub(v Money) Money {
	assertSameCurrency(m, v)
	m.value -= v.value
	return m
}

// Mul is an arithmetic operator.
func (m Money) Mul(n int64) Money {
	m.value = m.value * n
	return m
}

// Div is an arithmetic operator. This operation will perform rounding of the
// resulting value if necessary. If you need to accurately divide a money object
// with lossless precision, use the Split or Allocate function instead.
func (m Money) Div(f float64) Money {
	m.value = m.round(float64(m.value) / f)
	return m
}

// Abs returns a money object with an absolute value.
func (m Money) Abs() Money {
	if m.value < 0 {
		return m.FlipSign()
	}
	return m
}

// FlipSign flips the sign of the money object's value. Switching positive to
// negative and vice versa.
func (m Money) FlipSign() Money {
	m.value = -m.value
	return m
}

// Eq is a logical operator.
func (m Money) Eq(v Money) bool {
	assertSameCurrency(m, v)
	return m.value == v.value
}

// Neq is a logical operator.
func (m Money) Neq(v Money) bool {
	assertSameCurrency(m, v)
	return m.value != v.value
}

// Gt is a logical operator.
func (m Money) Gt(v Money) bool {
	assertSameCurrency(m, v)
	return m.value > v.value
}

// Gte is a logical operator.
func (m Money) Gte(v Money) bool {
	assertSameCurrency(m, v)
	return m.value >= v.value
}

// Lt is a logical operator.
func (m Money) Lt(v Money) bool {
	assertSameCurrency(m, v)
	return m.value < v.value
}

// Lte is a logical operator.
func (m Money) Lte(v Money) bool {
	assertSameCurrency(m, v)
	return m.value <= v.value
}

// IsZero returns a boolean value if the value is zero.
func (m Money) IsZero() bool {
	return m.value == 0
}

// IsPos returns a boolean value if the value is positive.
func (m Money) IsPos() bool {
	return m.value >= 0
}

// IsNeg returns a boolean value if the value is negative.
func (m Money) IsNeg() bool {
	return m.value < 0
}

// Split returns a slice containing money objects split as evenly as possible by
// 'n' times. This operation is lossless and will account for all remainders.
func (m Money) Split(n int64) []Money {
	if n <= 0 {
		panic("Failed to split money by zero")
	}
	s := make([]Money, 0, n)
	rem := int64(math.Mod(float64(m.value), float64(n)))
	value := int64(m.value / n)
	var i int64
	for i = 0; i < n; i++ {
		if rem > 0 {
			piece, _ := FromSubunits(m.format.code, value+1, m.round)
			rem--
			s = append(s, piece)
		} else {
			piece, _ := FromSubunits(m.format.code, value, m.round)
			s = append(s, piece)
		}
	}
	return s
}

// Allocate returns a slice containing money objects split according to the
// passed ratios. The ratios are completely arbitrary and are calculated as
// percentages of the overall sum. This operation is lossless and will account
// for all remainders.
func (m Money) Allocate(ratios ...int64) []Money {
	var sum int64 = 0
	for _, n := range ratios {
		sum += n
	}
	if sum <= 0 {
		panic("Failed to allocate money, no ratios passed")
	}
	s := make([]Money, 0, len(ratios))
	var allocated int64 = 0
	for _, n := range ratios {
		value := m.value * n / sum
		piece, _ := FromSubunits(m.format.code, value, m.round)
		s = append(s, piece)
		allocated += value
	}
	rem := m.value - allocated
	for i := 0; i < len(ratios) && rem > 0; i++ {
		s[i].value++
		rem--
	}
	return s
}

// MarshalJSON is an implementation of json.Marshaller.
func (m Money) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf(`{"currency": "%s", "formatted":"%s"}`, m.format.code, m.String())
	return []byte(json), nil
}

// String is an implementation of fmt.Stringer and returns the string
// formatted representation of the monetary value.
func (m Money) String() string {
	return strings.Replace(m.format.template, "0", m.StringNoSymbol(), 1)
}

// StringNoSymbol returns the string formatted representation of the monetary
// value without a currency symbol.
func (m Money) StringNoSymbol() string {
	str := strconv.FormatInt(m.value, 10)

	if len(str) <= m.format.subunits {
		str = strings.Repeat("0", m.format.subunits-len(str)+1) + str
	}

	if m.format.thouSep != "" {
		for i := len(str) - m.format.subunits - 3; i > 0; i -= 3 {
			str = str[:i] + m.format.thouSep + str[i:]
		}
	}

	if m.format.subunits > 0 {
		str = str[:len(str)-m.format.subunits] + m.format.subSep + str[len(str)-m.format.subunits:]
	}

	return str
}
