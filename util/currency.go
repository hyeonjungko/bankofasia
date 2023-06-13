package util

// All supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	KRW = "KRW"
)

// isSupportedCurrency returns true if given currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, KRW:
		return true
	}
	return false
}
