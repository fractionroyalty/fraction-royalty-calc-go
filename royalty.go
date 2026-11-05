// Package royaltycalc provides tiny helpers for oil and gas fraction royalty decimal math.
package royaltycalc

// Decimal converts a royalty fraction numerator/denominator to a float64 decimal.
func Decimal(numerator, denominator int64) float64 {
	if denominator == 0 {
		return 0
	}
	return float64(numerator) / float64(denominator)
}
