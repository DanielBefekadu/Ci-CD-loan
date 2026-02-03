package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotalInterest(t *testing.T) {
	total := CalculateTotalInterest(1200, 12, 12)
	assert.Equal(t, 144.0, total) // Correct total interest for 12% annual / 12 months

	totalZero := CalculateTotalInterest(1200, 0, 12)
	assert.Equal(t, 0.0, totalZero) // Zero-interest case
}
