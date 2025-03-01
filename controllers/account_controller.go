package controllers

import (
	"encoding/json"
	"golang/config"
	"net/http"
)


type BalanceSheet struct {
	ID              int    `json:"id"`
	ExpenseCategory string `json:"expense_category"`
	Type            string `json:"type"`
	Amount          int    `json:"amount"`
	DueDate         string `json:"due_date"`
	Status          string `json:"status"`
	PaymentMethod   string `json:"payment_method"`
	EnteredBy       string `json:"entered_by"`
	EnteredDate     string `json:"entered_date"`
	UpdatedBy       *string `json:"updated_by,omitempty"` 
	UpdatedDate     *string `json:"updated_date,omitempty"` 
}


func GetBalanceSheet(w http.ResponseWriter, r *http.Request) {
	var balanceSheets []BalanceSheet


	rows, err := config.DB.Query("SELECT id, expense_category, type, Amount, due_date, status, payment_method, entered_by, entered_date, updated_by, updated_date FROM balance_sheet")
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()


	for rows.Next() {
		var bs BalanceSheet
		err := rows.Scan(
			&bs.ID, &bs.ExpenseCategory, &bs.Type, &bs.Amount, 
			&bs.DueDate, &bs.Status, &bs.PaymentMethod, &bs.EnteredBy, 
			&bs.EnteredDate, &bs.UpdatedBy, &bs.UpdatedDate,
		)
		if err != nil {
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		balanceSheets = append(balanceSheets, bs)
	}

	if len(balanceSheets) == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "0",
			"message": "No records found",
			"data":    []BalanceSheet{},
		})
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "1",
		"message": "Balance sheet data fetched successfully",
		"data":    balanceSheets,
	})
}
