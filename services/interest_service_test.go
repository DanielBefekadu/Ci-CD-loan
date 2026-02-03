package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotalInterest(t *testing.T) {
	total := CalculateTotalInterest(1200, 12, 12)
	assert.Equal(t, 120.0, total) // 12% annual / 12 months = 1% per month * 12 months

	totalZero := CalculateTotalInterest(1200, 0, 12)
	assert.Equal(t, 0.0, totalZero) // Zero-interest case
}
