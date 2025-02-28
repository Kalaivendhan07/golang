package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang/config" 
)


type KPIData struct {
	SKUHoldingCost        float64 `json:"sku_holding_cost"`
	AverageTransactionValue float64 `json:"average_transaction_value"`
	FootfallConversionRate float64 `json:"footfall_conversion_rate"`
}


type KPIResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    KPIData `json:"data"`
}


func GetKPIData3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching KPI Metrics...")


	var totalRevenue, totalTransactions, totalVisitors int
	var totalSKUUnits, storageCostPerSKU float64


	err := config.DB.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0) AS total_revenue, COUNT(*) AS total_transactions FROM sales WHERE status = 'Completed'
	`).Scan(&totalRevenue, &totalTransactions)

	if err != nil {
		http.Error(w, "Error fetching revenue and transactions", http.StatusInternalServerError)
		return
	}


	err = config.DB.QueryRow(`
		SELECT COALESCE(SUM(amount), 0) AS total_sku_units FROM purchases WHERE status = 'Completed'
	`).Scan(&totalSKUUnits)

	if err != nil {
		http.Error(w, "Error fetching SKU units", http.StatusInternalServerError)
		return
	}


	err = config.DB.QueryRow(`
		SELECT COALESCE(SUM(visitors_count), 0) AS total_visitors FROM store_visitors
	`).Scan(&totalVisitors)

	if err != nil {
		totalVisitors = 1000 
	}


	storageCostPerSKU = 5.0


	skuHoldingCost := storageCostPerSKU * totalSKUUnits

	var avgTransactionValue float64
	if totalTransactions > 0 {
		avgTransactionValue = float64(totalRevenue) / float64(totalTransactions)
	} else {
		avgTransactionValue = 0
	}


	var conversionRate float64
	if totalVisitors > 0 {
		conversionRate = (float64(totalTransactions) / float64(totalVisitors)) * 100
	} else {
		conversionRate = 0
	}


	response := KPIResponse{
		Status:  "1",
		Message: "KPI Metrics Fetched",
		Data: KPIData{
			SKUHoldingCost:        skuHoldingCost,
			AverageTransactionValue: avgTransactionValue,
			FootfallConversionRate: conversionRate,
		},
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
