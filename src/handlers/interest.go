package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Ci-CD-loan/services"
)

// TotalInterestHandler handles /total-interest requests
func TotalInterestHandler(w http.ResponseWriter, r *http.Request) {
	principalStr := r.URL.Query().Get("principal")
	rateStr := r.URL.Query().Get("rate")
	termStr := r.URL.Query().Get("term")

	principal, _ := strconv.ParseFloat(principalStr, 64)
	rate, _ := strconv.ParseFloat(rateStr, 64)
	term, _ := strconv.Atoi(termStr)

	total := services.CalculateTotalInterest(principal, rate, term)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"total_interest": total})
}
