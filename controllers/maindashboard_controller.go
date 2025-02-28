package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang/config" 
)


type YTDResponse struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	YTD_Sales    int    `json:"ytd_sales"`
	YTD_Purchase int    `json:"ytd_purchase"`
	Net_Profit   int    `json:"net_profit"`
}


func GetYearToDateProfit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching YTD Sales, Purchases, and Net Profit...")

	var ytdSales int
	err := config.DB.QueryRow(`SELECT COALESCE(SUM(total_amount), 0) FROM sales WHERE YEAR(ordered_date) = YEAR(CURDATE())`).Scan(&ytdSales)
	if err != nil {
		http.Error(w, "Database error fetching sales", http.StatusInternalServerError)
		return
	}

	var ytdPurchases int
	err = config.DB.QueryRow(`SELECT COALESCE(SUM(total_amount), 0) FROM purchases WHERE YEAR(purchase_date) = YEAR(CURDATE())`).Scan(&ytdPurchases)
	if err != nil {
		http.Error(w, "Database error fetching purchases", http.StatusInternalServerError)
		return
	}


	netProfit := ytdSales - ytdPurchases

	response := YTDResponse{
		Status:       "1",
		Message:      "YTD Sales, Purchases & Net Profit Calculated",
		YTD_Sales:    ytdSales,
		YTD_Purchase: ytdPurchases,
		Net_Profit:   netProfit,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
