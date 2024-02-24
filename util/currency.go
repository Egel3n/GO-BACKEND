package util

const (
	EUR = "EUR"
	USD = "USD"
	TRY = "TRY"
)

func IsCurrencySupported(Currency string) bool {
	switch Currency {
	case EUR, USD, TRY:
		return true
	}
	return false
}
