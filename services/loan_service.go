package services

type Repayment struct {
	Month     int     `json:"month"`
	Principal float64 `json:"principal"`
	Interest  float64 `json:"interest"`
	Total     float64 `json:"total"`
}

func CalculateRepayment(principal, rate float64, term int) []Repayment {
	schedule := make([]Repayment, term)
	monthlyInterest := 0.0
	if rate > 0 {
		monthlyInterest = rate / 12 / 100
	}

	for i := 0; i < term; i++ {
		interest := principal * monthlyInterest
		monthlyPayment := principal/float64(term) + interest
		schedule[i] = Repayment{
			Month:     i + 1,
			Principal: principal / float64(term),
			Interest:  interest,
			Total:     monthlyPayment,
		}
	}
	return schedule
}
