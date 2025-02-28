package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang/config" 
)


type PurchaseData struct {
	ID            int     `json:"id"`
	PurchaseID    string  `json:"purchase_id"`
	VendorID      int     `json:"vendor_id"`
	InvoiceNo     string  `json:"invoice_no"`
	PurchaseDate  string  `json:"purchase_date"`
	ProductName   string  `json:"product_name"`
	SKUCode       string  `json:"sku_code"`
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


type PurchaseResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    []PurchaseData `json:"data"`
}


func GetAllPurchases(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all purchase records...")

	rows, err := config.DB.Query(`
		SELECT id, purchase_id, vendor_id, invoice_no, 
			   DATE_FORMAT(purchase_date, '%Y-%m-%d') AS purchase_date, 
			   product_name, sku_code, amount, tax, total_amount, 
			   status, payment_method, entered_by, 
			   DATE_FORMAT(entered_date, '%Y-%m-%d %H:%i:%s') AS entered_date, 
			   updated_by, DATE_FORMAT(updated_date, '%Y-%m-%d %H:%i:%s') AS updated_date 
		FROM purchases
	`)
	if err != nil {
		http.Error(w, "Error fetching purchase data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var purchaseRecords []PurchaseData

	for rows.Next() {
		var purchase PurchaseData
		err := rows.Scan(
			&purchase.ID, &purchase.PurchaseID, &purchase.VendorID, &purchase.InvoiceNo,
			&purchase.PurchaseDate, &purchase.ProductName, &purchase.SKUCode,
			&purchase.Amount, &purchase.Tax, &purchase.TotalAmount, &purchase.Status,
			&purchase.PaymentMethod, &purchase.EnteredBy, &purchase.EnteredDate,
			&purchase.UpdatedBy, &purchase.UpdatedDate,
		)
		if err != nil {
			http.Error(w, "Error scanning purchase data", http.StatusInternalServerError)
			return
		}
		purchaseRecords = append(purchaseRecords, purchase)
	}

	response := PurchaseResponse{
		Status:  "1",
		Message: "Purchase data fetched successfully",
		Data:    purchaseRecords,
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
