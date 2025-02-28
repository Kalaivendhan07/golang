package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang/config" 
)


type SalesData2 struct {
	InvoiceNo     string `json:"invoice_no"`
	CustomerID    int    `json:"customer_id"`
	Product       string `json:"product"`
	SKU           string `json:"sku"`
	OrderedDate   string `json:"ordered_date"`
	Amount        int    `json:"amount"`
	Tax           int    `json:"tax"`
	TotalAmount   int    `json:"total_amount"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
	EnteredBy     string `json:"entered_by"`
}

type SalesResponse2 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}


func AddSales(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding new sales record...")

	var sales SalesData2

	err := json.NewDecoder(r.Body).Decode(&sales)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO sales 
		(invoice_no, customer_id, product, sku, ordered_date, amount, tax, total_amount, status, payment_method, entered_by) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = config.DB.Exec(query,
		sales.InvoiceNo, sales.CustomerID, sales.Product, sales.SKU, sales.OrderedDate,
		sales.Amount, sales.Tax, sales.TotalAmount, sales.Status, sales.PaymentMethod, sales.EnteredBy,
	)

	if err != nil {
		http.Error(w, "Failed to insert sales record", http.StatusInternalServerError)
		return
	}


	response := SalesResponse2{
		Status:  "1",
		Message: "Sales record added successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
