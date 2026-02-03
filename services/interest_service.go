package services

// CalculateTotalInterest calculates the total interest over the term
func CalculateTotalInterest(principal, rate float64, term int) float64 {
	if rate <= 0 {
		return 0
	}

	monthlyInterest := rate / 12 / 100
	totalInterest := 0.0

	for i := 0; i < term; i++ {
		interest := principal * monthlyInterest
		totalInterest += interest
	}

	return totalInterest
}
