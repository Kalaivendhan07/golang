package controllers

import (
	"encoding/json"
	"net/http"
	"golang/config" 
)


type BalanceSheetData struct {
	ExpenseCategory string `json:"expense_category"`
	Amount          int    `json:"amount"`
	DueDate         string `json:"due_date"`
	Status          string `json:"status"`
	PaymentMethod   string `json:"payment_method"`
	EnteredBy       string `json:"entered_by"`
	Type            string `json:"type"`
}


type BalanceSheetResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}


func AddBalanceSheetEntry(w http.ResponseWriter, r *http.Request) {

	var entry BalanceSheetData


	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO balance_sheet 
		(expense_category, amount, due_date, status, payment_method, entered_by, type) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err = config.DB.Exec(query,
		entry.ExpenseCategory, entry.Amount, entry.DueDate, entry.Status,
		entry.PaymentMethod, entry.EnteredBy, entry.Type,
	)

	if err != nil {
		http.Error(w, "Failed to insert balance sheet entry", http.StatusInternalServerError)
		return
	}


	response := BalanceSheetResponse{
		Status:  "1",
		Message: "Balance sheet entry added successfully",
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
