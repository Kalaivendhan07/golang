package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"golang/config" 
)


type PurchaseData2 struct {
	PurchaseID    string `json:"purchase_id"`
	VendorID      int    `json:"vendor_id"`
	InvoiceNo     string `json:"invoice_no"`
	PurchaseDate  string `json:"purchase_date"`
	ProductName   string `json:"product_name"`
	SKUCode       string `json:"sku_code"`
	Amount        int    `json:"amount"`
	Tax           int    `json:"tax"`
	TotalAmount   int    `json:"total_amount"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
	EnteredBy     string `json:"entered_by"`
}


type PurchaseResponse2 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}


func AddPurchase(w http.ResponseWriter, r *http.Request) {


	var purchase PurchaseData2


	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}


	query := `
		INSERT INTO purchases 
		(purchase_id, vendor_id, invoice_no, purchase_date, product_name, sku_code, amount, tax, total_amount, status, payment_method, entered_by) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = config.DB.Exec(query,
		purchase.PurchaseID, purchase.VendorID, purchase.InvoiceNo, purchase.PurchaseDate,
		purchase.ProductName, purchase.SKUCode, purchase.Amount, purchase.Tax,
		purchase.TotalAmount, purchase.Status, purchase.PaymentMethod, purchase.EnteredBy,
	)

	if err != nil {
		http.Error(w, "Failed to insert purchase record", http.StatusInternalServerError)
		return
	}


	response := PurchaseResponse2{
		Status:  "1",
		Message: "Purchase record added successfully",
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
