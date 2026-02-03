package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Ci-CD-loan/services"
)

func RepaymentScheduleHandler(w http.ResponseWriter, r *http.Request) {
	principalStr := r.URL.Query().Get("principal")
	rateStr := r.URL.Query().Get("rate")
	termStr := r.URL.Query().Get("term")

	principal, _ := strconv.ParseFloat(principalStr, 64)
	rate, _ := strconv.ParseFloat(rateStr, 64)
	term, _ := strconv.Atoi(termStr)

	schedule := services.CalculateRepayment(principal, rate, term)

	w.Header().Set("Content-Type", "application/json")
	
	// FIX: Capture and handle the error to satisfy errcheck
	if err := json.NewEncoder(w).Encode(schedule); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
