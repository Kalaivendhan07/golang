package models

// import "time"

// BalanceSheet represents the structure of the balance_sheet table
type BalanceSheet struct {
	ID              int    `json:"id"`
	ExpenseCategory string `json:"expense_category"`
	Amount          int    `json:"amount"`
	DueDate         string `json:"due_date"`
	Status          string `json:"status"`
	PaymentMethod   string `json:"payment_method"`
	EnteredBy       string `json:"entered_by"`
	EnteredDate     string `json:"entered_date"`
	UpdatedBy       string `json:"updated_by,omitempty"`
	UpdatedDate     string `json:"updated_date,omitempty"`
}

