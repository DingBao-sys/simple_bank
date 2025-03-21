package utils

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	SGD = "SGD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, SGD:
		return true
	}
	return false
}
