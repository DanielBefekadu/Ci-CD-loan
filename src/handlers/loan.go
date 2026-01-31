package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanielBefekadu/Ci-CD-loan/services"
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
	json.NewEncoder(w).Encode(schedule)
}
