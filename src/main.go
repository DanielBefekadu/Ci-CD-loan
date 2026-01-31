package main

import (
	"log"
	"net/http"

	"github.com/DanielBefekadu/Ci-CD-loan/src/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/repayment-schedule", handlers.RepaymentScheduleHandler).Methods("GET")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
