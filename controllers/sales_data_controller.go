package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang/config" 
)

type SalesData struct {
	ID            int     `json:"id"`
	InvoiceNo     string  `json:"invoice_no"`
	CustomerID    int     `json:"customer_id"`
	Product       string  `json:"product"`
	SKU           string  `json:"sku"`
	OrderedDate   string  `json:"ordered_date"`
	Amount        int     `json:"amount"`
	Tax           int     `json:"tax"`
	TotalAmount   int     `json:"total_amount"`
	Status        string  `json:"status"`
	PaymentMethod string  `json:"payment_method"`
	EnteredBy     string  `json:"entered_by"`
	EnteredDate   string  `json:"entered_date"`
	UpdatedBy     *string `json:"updated_by,omitempty"`
	UpdatedDate   string  `json:"updated_date"`
}


type SalesResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []SalesData `json:"data"`
}


func GetAllSales(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all sales records...")

	rows, err := config.DB.Query(`
		SELECT id, invoice_no, customer_id, product, sku, 
			   DATE_FORMAT(ordered_date, '%Y-%m-%d') AS ordered_date, 
			   amount, tax, total_amount, status, payment_method, 
			   entered_by, DATE_FORMAT(entered_date, '%Y-%m-%d %H:%i:%s') AS entered_date, 
			   updated_by, DATE_FORMAT(updated_date, '%Y-%m-%d %H:%i:%s') AS updated_date 
		FROM sales
	`)
	if err != nil {
		http.Error(w, "Error fetching sales data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var salesRecords []SalesData

	for rows.Next() {
		var sale SalesData
		err := rows.Scan(
			&sale.ID, &sale.InvoiceNo, &sale.CustomerID, &sale.Product, &sale.SKU,
			&sale.OrderedDate, &sale.Amount, &sale.Tax, &sale.TotalAmount, &sale.Status,
			&sale.PaymentMethod, &sale.EnteredBy, &sale.EnteredDate, &sale.UpdatedBy, &sale.UpdatedDate,
		)
		if err != nil {
			http.Error(w, "Error scanning sales data", http.StatusInternalServerError)
			return
		}
		salesRecords = append(salesRecords, sale)
	}

	response := SalesResponse{
		Status:  "1",
		Message: "Sales data fetched successfully",
		Data:    salesRecords,
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
