package mongo

// Format represents the currency's currencyFormat.
type currencyFormat struct {
	code     string // The ISO 4217 currency code.
	subunits int    // The number of subunits.
	thouSep  string // The thousand separator.
	subSep   string // The subunit separator.
	template string // The string format template.
}

// CurrencyFormats contain a map of all recognised currency formats.
var currencyFormats = map[string]currencyFormat{
	"AED": {code: "AED", subunits: 2, thouSep: ",", subSep: ".", template: "0 د.إ"},
	"AFN": {code: "AFN", subunits: 2, thouSep: ",", subSep: ".", template: "0 ؋"},
	"ALL": {code: "ALL", subunits: 2, thouSep: ",", subSep: ".", template: "L0"},
	"AMD": {code: "AMD", subunits: 2, thouSep: ",", subSep: ".", template: "0 ֏"},
	"ANG": {code: "ANG", subunits: 2, thouSep: ".", subSep: ",", template: "ƒ0"},
	"AOA": {code: "AOA", subunits: 2, thouSep: ",", subSep: ".", template: "0Kz"},
	"ARS": {code: "ARS", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"AUD": {code: "AUD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"AWG": {code: "AWG", subunits: 2, thouSep: ",", subSep: ".", template: "0ƒ"},
	"AZN": {code: "AZN", subunits: 2, thouSep: ",", subSep: ".", template: "m0"},
	"BAM": {code: "BAM", subunits: 2, thouSep: ",", subSep: ".", template: "KM0"},
	"BBD": {code: "BBD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"BDT": {code: "BDT", subunits: 2, thouSep: ",", subSep: ".", template: "৳0"},
	"BGN": {code: "BGN", subunits: 2, thouSep: ",", subSep: ".", template: "лв0"},
	"BHD": {code: "BHD", subunits: 3, thouSep: ",", subSep: ".", template: "0 .د.ب "},
	"BIF": {code: "BIF", subunits: 0, thouSep: ",", subSep: ".", template: "0Fr"},
	"BMD": {code: "BMD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"BND": {code: "BND", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"BOB": {code: "BOB", subunits: 2, thouSep: ",", subSep: ".", template: "Bs.0"},
	"BRL": {code: "BRL", subunits: 2, thouSep: ".", subSep: ",", template: "R$0"},
	"BSD": {code: "BSD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"BTN": {code: "BTN", subunits: 2, thouSep: ",", subSep: ".", template: "0Nu."},
	"BWP": {code: "BWP", subunits: 2, thouSep: ",", subSep: ".", template: "P0"},
	"BYN": {code: "BYN", subunits: 2, thouSep: " ", subSep: ",", template: "0 p."},
	"BYR": {code: "BYR", subunits: 0, thouSep: " ", subSep: ",", template: "0 p."},
	"BZD": {code: "BZD", subunits: 2, thouSep: ",", subSep: ".", template: "BZ$0"},
	"CAD": {code: "CAD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"CDF": {code: "CDF", subunits: 2, thouSep: ",", subSep: ".", template: "0FC"},
	"CHF": {code: "CHF", subunits: 2, thouSep: ",", subSep: ".", template: "0 CHF"},
	"CLF": {code: "CLF", subunits: 4, thouSep: ".", subSep: ",", template: "UF0"},
	"CLP": {code: "CLP", subunits: 0, thouSep: ".", subSep: ",", template: "$0"},
	"CNY": {code: "CNY", subunits: 2, thouSep: ",", subSep: ".", template: "0 ¥"},
	"COP": {code: "COP", subunits: 2, thouSep: ".", subSep: ",", template: "$0"},
	"CRC": {code: "CRC", subunits: 2, thouSep: ",", subSep: ".", template: "₡0"},
	"CUC": {code: "CUC", subunits: 2, thouSep: ",", subSep: ".", template: "0$"},
	"CUP": {code: "CUP", subunits: 2, thouSep: ",", subSep: ".", template: "$MN0"},
	"CVE": {code: "CVE", subunits: 2, thouSep: ",", subSep: ".", template: "0$"},
	"CZK": {code: "CZK", subunits: 2, thouSep: ",", subSep: ".", template: "0 Kč"},
	"DJF": {code: "DJF", subunits: 0, thouSep: ",", subSep: ".", template: "0 Fdj"},
	"DKK": {code: "DKK", subunits: 2, thouSep: ".", subSep: ",", template: "kr 1"},
	"DOP": {code: "DOP", subunits: 2, thouSep: ",", subSep: ".", template: "RD$0"},
	"DZD": {code: "DZD", subunits: 2, thouSep: ",", subSep: ".", template: "0 دج "},
	"EEK": {code: "EEK", subunits: 2, thouSep: ",", subSep: ".", template: "kr0"},
	"EGP": {code: "EGP", subunits: 2, thouSep: ",", subSep: ".", template: "ج.م 0"},
	"ERN": {code: "ERN", subunits: 2, thouSep: ",", subSep: ".", template: "0 Nfk"},
	"ETB": {code: "ETB", subunits: 2, thouSep: ",", subSep: ".", template: "0 Br"},
	"EUR": {code: "EUR", subunits: 2, thouSep: ",", subSep: ".", template: "€0"},
	"FJD": {code: "FJD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"FKP": {code: "FKP", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"GBP": {code: "GBP", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"GEL": {code: "GEL", subunits: 2, thouSep: ",", subSep: ".", template: "0 ლ"},
	"GGP": {code: "GGP", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"GHC": {code: "GHC", subunits: 2, thouSep: ",", subSep: ".", template: "GH₵0"},
	"GHS": {code: "GHS", subunits: 2, thouSep: ",", subSep: ".", template: "GH₵0"},
	"GIP": {code: "GIP", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"GMD": {code: "GMD", subunits: 2, thouSep: ",", subSep: ".", template: "0 D"},
	"GNF": {code: "GNF", subunits: 0, thouSep: ",", subSep: ".", template: "0 FG"},
	"GTQ": {code: "GTQ", subunits: 2, thouSep: ",", subSep: ".", template: "Q0"},
	"GYD": {code: "GYD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"HKD": {code: "HKD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"HNL": {code: "HNL", subunits: 2, thouSep: ",", subSep: ".", template: "L0"},
	"HRK": {code: "HRK", subunits: 2, thouSep: ".", subSep: ",", template: "0 Kn"},
	"HTG": {code: "HTG", subunits: 2, thouSep: ".", subSep: ",", template: "0 G"},
	"HUF": {code: "HUF", subunits: 0, thouSep: ",", subSep: ".", template: "Ft0"},
	"IDR": {code: "IDR", subunits: 2, thouSep: ",", subSep: ".", template: "Rp0"},
	"ILS": {code: "ILS", subunits: 2, thouSep: ",", subSep: ".", template: "₪0"},
	"IMP": {code: "IMP", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"INR": {code: "INR", subunits: 2, thouSep: ",", subSep: ".", template: "₹0"},
	"IQD": {code: "IQD", subunits: 3, thouSep: ",", subSep: ".", template: "0 د.ع"},
	"IRR": {code: "IRR", subunits: 2, thouSep: ",", subSep: ".", template: "0 ﷼"},
	"ISK": {code: "ISK", subunits: 0, thouSep: ".", subSep: ",", template: "Kr0"},
	"JEP": {code: "JEP", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"JMD": {code: "JMD", subunits: 2, thouSep: ",", subSep: ".", template: "J$0"},
	"JOD": {code: "JOD", subunits: 3, thouSep: ",", subSep: ".", template: "0 د.أ"},
	"JPY": {code: "JPY", subunits: 0, thouSep: ",", subSep: ".", template: "¥0"},
	"KES": {code: "KES", subunits: 2, thouSep: ",", subSep: ".", template: "KSh0"},
	"KGS": {code: "KGS", subunits: 2, thouSep: ",", subSep: ".", template: "С̲0"},
	"KHR": {code: "KHR", subunits: 2, thouSep: ",", subSep: ".", template: "៛0"},
	"KMF": {code: "KMF", subunits: 0, thouSep: ",", subSep: ".", template: "CF0"},
	"KPW": {code: "KPW", subunits: 0, thouSep: ",", subSep: ".", template: "₩0"},
	"KRW": {code: "KRW", subunits: 0, thouSep: ",", subSep: ".", template: "₩0"},
	"KWD": {code: "KWD", subunits: 3, thouSep: ",", subSep: ".", template: "0 د.ك"},
	"KYD": {code: "KYD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"KZT": {code: "KZT", subunits: 2, thouSep: ",", subSep: ".", template: "₸0"},
	"LAK": {code: "LAK", subunits: 2, thouSep: ",", subSep: ".", template: "₭0"},
	"LBP": {code: "LBP", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"LKR": {code: "LKR", subunits: 2, thouSep: ",", subSep: ".", template: "රු, ரூ0"},
	"LRD": {code: "LRD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"LSL": {code: "LSL", subunits: 2, thouSep: ",", subSep: ".", template: "L0"},
	"LTL": {code: "LTL", subunits: 2, thouSep: ",", subSep: ".", template: "Lt0"},
	"LVL": {code: "LVL", subunits: 2, thouSep: ",", subSep: ".", template: "0 Ls"},
	"LYD": {code: "LYD", subunits: 3, thouSep: ",", subSep: ".", template: "0 ل.د"},
	"MAD": {code: "MAD", subunits: 2, thouSep: ",", subSep: ".", template: "0 DH"},
	"MDL": {code: "MDL", subunits: 2, thouSep: ",", subSep: ".", template: "0 lei"},
	"MKD": {code: "MKD", subunits: 2, thouSep: ",", subSep: ".", template: "ден0"},
	"MMK": {code: "MMK", subunits: 2, thouSep: ",", subSep: ".", template: "K0"},
	"MNT": {code: "MNT", subunits: 2, thouSep: ",", subSep: ".", template: "₮0"},
	"MOP": {code: "MOP", subunits: 2, thouSep: ",", subSep: ".", template: "0 P"},
	"MUR": {code: "MUR", subunits: 2, thouSep: ",", subSep: ".", template: "₨0"},
	"MVR": {code: "MVR", subunits: 2, thouSep: ",", subSep: ".", template: "0 MVR"},
	"MWK": {code: "MWK", subunits: 2, thouSep: ",", subSep: ".", template: "MK0"},
	"MXN": {code: "MXN", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"MYR": {code: "MYR", subunits: 2, thouSep: ",", subSep: ".", template: "RM0"},
	"MZN": {code: "MZN", subunits: 2, thouSep: ",", subSep: ".", template: "MT0"},
	"NAD": {code: "NAD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"NGN": {code: "NGN", subunits: 2, thouSep: ",", subSep: ".", template: "₦0"},
	"NIO": {code: "NIO", subunits: 2, thouSep: ",", subSep: ".", template: "C$0"},
	"NOK": {code: "NOK", subunits: 2, thouSep: ",", subSep: ".", template: "0 Kr"},
	"NPR": {code: "NPR", subunits: 2, thouSep: ",", subSep: ".", template: "रु0"},
	"NZD": {code: "NZD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"OMR": {code: "OMR", subunits: 3, thouSep: ",", subSep: ".", template: "0 ر.ع."},
	"PAB": {code: "PAB", subunits: 2, thouSep: ",", subSep: ".", template: "B/.0"},
	"PEN": {code: "PEN", subunits: 2, thouSep: ",", subSep: ".", template: "S/0"},
	"PGK": {code: "PGK", subunits: 2, thouSep: ",", subSep: ".", template: "0 K"},
	"PHP": {code: "PHP", subunits: 2, thouSep: ",", subSep: ".", template: "₱0"},
	"PKR": {code: "PKR", subunits: 2, thouSep: ",", subSep: ".", template: "₨0"},
	"PLN": {code: "PLN", subunits: 2, thouSep: ",", subSep: ".", template: "0 zł"},
	"PYG": {code: "PYG", subunits: 0, thouSep: ",", subSep: ".", template: "0Gs"},
	"QAR": {code: "QAR", subunits: 2, thouSep: ",", subSep: ".", template: "0 ر.ق"},
	"RON": {code: "RON", subunits: 2, thouSep: ",", subSep: ".", template: "lei0"},
	"RSD": {code: "RSD", subunits: 2, thouSep: ",", subSep: ".", template: "дин0"},
	"RUB": {code: "RUB", subunits: 2, thouSep: ",", subSep: ".", template: "0 ₽"},
	"RUR": {code: "RUR", subunits: 2, thouSep: ",", subSep: ".", template: "0 ₽"},
	"RWF": {code: "RWF", subunits: 0, thouSep: ",", subSep: ".", template: "0 FRw"},
	"SAR": {code: "SAR", subunits: 2, thouSep: ",", subSep: ".", template: "0 ر.س"},
	"SBD": {code: "SBD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"SCR": {code: "SCR", subunits: 2, thouSep: ",", subSep: ".", template: "SCR0"},
	"SDG": {code: "SDG", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"SEK": {code: "SEK", subunits: 2, thouSep: ",", subSep: ".", template: "0 Kr"},
	"SGD": {code: "SGD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"SHP": {code: "SHP", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"SKK": {code: "SKK", subunits: 2, thouSep: ",", subSep: ".", template: "Sk0"},
	"SLL": {code: "SLL", subunits: 2, thouSep: ",", subSep: ".", template: "0 Le"},
	"SOS": {code: "SOS", subunits: 2, thouSep: ",", subSep: ".", template: "0 Sh"},
	"SRD": {code: "SRD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"SSP": {code: "SSP", subunits: 2, thouSep: ",", subSep: ".", template: "0 £"},
	"STD": {code: "STD", subunits: 2, thouSep: ",", subSep: ".", template: "0 Db"},
	"SVC": {code: "SVC", subunits: 2, thouSep: ",", subSep: ".", template: "₡0"},
	"SYP": {code: "SYP", subunits: 2, thouSep: ",", subSep: ".", template: "0 £"},
	"SZL": {code: "SZL", subunits: 2, thouSep: ",", subSep: ".", template: "£0"},
	"THB": {code: "THB", subunits: 2, thouSep: ",", subSep: ".", template: "฿ 20"},
	"TJS": {code: "TJS", subunits: 2, thouSep: ",", subSep: ".", template: "0 SM"},
	"TMT": {code: "TMT", subunits: 2, thouSep: ",", subSep: ".", template: "0 T"},
	"TND": {code: "TND", subunits: 3, thouSep: ",", subSep: ".", template: "0 د.ت"},
	"TOP": {code: "TOP", subunits: 2, thouSep: ",", subSep: ".", template: "T$0"},
	"TRL": {code: "TRL", subunits: 2, thouSep: ",", subSep: ".", template: "₺0"},
	"TRY": {code: "TRY", subunits: 2, thouSep: ",", subSep: ".", template: "₺0"},
	"TTD": {code: "TTD", subunits: 2, thouSep: ",", subSep: ".", template: "TT$0"},
	"TWD": {code: "TWD", subunits: 2, thouSep: ",", subSep: ".", template: "NT$0"},
	"TZS": {code: "TZS", subunits: 0, thouSep: ",", subSep: ".", template: "TSh0"},
	"UAH": {code: "UAH", subunits: 2, thouSep: ",", subSep: ".", template: "0 ₴"},
	"UGX": {code: "UGX", subunits: 0, thouSep: ",", subSep: ".", template: "0 USh"},
	"USD": {code: "USD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"UYU": {code: "UYU", subunits: 2, thouSep: ",", subSep: ".", template: "U$0"},
	"UZS": {code: "UZS", subunits: 2, thouSep: ",", subSep: ".", template: "сум0"},
	"VEF": {code: "VEF", subunits: 2, thouSep: ",", subSep: ".", template: "Bs0"},
	"VND": {code: "VND", subunits: 0, thouSep: ",", subSep: ".", template: "0 ₫"},
	"VUV": {code: "VUV", subunits: 0, thouSep: ",", subSep: ".", template: "Vt0"},
	"WST": {code: "WST", subunits: 2, thouSep: ",", subSep: ".", template: "0 T"},
	"XAF": {code: "XAF", subunits: 0, thouSep: ",", subSep: ".", template: "0 Fr"},
	"XAG": {code: "XAG", subunits: 0, thouSep: ",", subSep: ".", template: "0 oz t"},
	"XAU": {code: "XAU", subunits: 0, thouSep: ",", subSep: ".", template: "0 oz t"},
	"XCD": {code: "XCD", subunits: 2, thouSep: ",", subSep: ".", template: "$0"},
	"XDR": {code: "XDR", subunits: 0, thouSep: ",", subSep: ".", template: "0 SDR"},
	"XPF": {code: "XPF", subunits: 0, thouSep: ",", subSep: ".", template: "0 ₣"},
	"YER": {code: "YER", subunits: 2, thouSep: ",", subSep: ".", template: "0 ر.ي, ﷼"},
	"ZAR": {code: "ZAR", subunits: 2, thouSep: ",", subSep: ".", template: "R0"},
	"ZMW": {code: "ZMW", subunits: 2, thouSep: ",", subSep: ".", template: "ZK0"},
	"ZWD": {code: "ZWD", subunits: 2, thouSep: ",", subSep: ".", template: "Z$0"},
}
