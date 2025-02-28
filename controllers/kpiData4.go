package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"golang/config" 
)


type MetricsData struct {
	ExpenseToRevenueRatio   float64 `json:"expense_to_revenue_ratio"`
	RevenueLostDueToStock   float64 `json:"revenue_lost_due_to_stock"`
	ProfitMargin           float64 `json:"profit_margin"`
	TopSellingSKU          string  `json:"top_selling_sku"`
	InventoryCost          float64 `json:"inventory_cost"`
}


type MetricsResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    MetricsData  `json:"data"`
}

func FetchMetricsData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching Business Metrics...")

	var revenueTotal, expenseTotal, profitTotal int
	var inventoryCost, revenueLostDueToStock float64
	var mostPopularSKU string
	var totalUnitsPurchased, costPerUnit float64


	err := config.DB.QueryRow(`SELECT COALESCE(SUM(total_amount), 0) FROM sales WHERE status = 'Completed'`).Scan(&revenueTotal)
	if err != nil {
		http.Error(w, "Error fetching revenue data", http.StatusInternalServerError)
		return
	}


	err = config.DB.QueryRow(`SELECT COALESCE(SUM(expense_amount), 0) FROM expenses`).Scan(&expenseTotal)
	if err != nil {
		expenseTotal = 0 
	}


	err = config.DB.QueryRow(`SELECT COALESCE(SUM(profit), 0) FROM profits`).Scan(&profitTotal)
	if err != nil {
		profitTotal = revenueTotal - expenseTotal 
	}

	currentMonth := time.Now().Format("2006-01") 
	err = config.DB.QueryRow(`
		SELECT sku 
		FROM sales 
		WHERE DATE_FORMAT(sale_date, '%Y-%m') = ? 
		GROUP BY sku 
		ORDER BY COUNT(*) DESC 
		LIMIT 1
	`, currentMonth).Scan(&mostPopularSKU)

	if err != nil {
		mostPopularSKU = "No Data Available"
	}


	err = config.DB.QueryRow(`SELECT COALESCE(SUM(amount), 0) FROM purchases WHERE status = 'Completed'`).Scan(&totalUnitsPurchased)
	if err != nil {
		http.Error(w, "Error fetching purchased unit data", http.StatusInternalServerError)
		return
	}


	costPerUnit = 5.0 
	inventoryCost = costPerUnit * totalUnitsPurchased

	var avgDailySales, outOfStockDays, avgSkuPrice float64
	err = config.DB.QueryRow(`
		SELECT 
			COALESCE(AVG(daily_sales), 0), 
			COALESCE(SUM(days_out_of_stock), 0), 
			COALESCE(AVG(price), 0) 
		FROM sku_stockouts
	`).Scan(&avgDailySales, &outOfStockDays, &avgSkuPrice)

	if err != nil {
		avgDailySales, outOfStockDays, avgSkuPrice = 0, 0, 0
	}

	revenueLostDueToStock = avgDailySales * outOfStockDays * avgSkuPrice


	var expenseToRevenueRatio float64
	if revenueTotal > 0 {
		expenseToRevenueRatio = (float64(expenseTotal) / float64(revenueTotal)) * 100
	} else {
		expenseToRevenueRatio = 0
	}

	var profitMargin float64
	if revenueTotal > 0 {
		profitMargin = (float64(profitTotal) / float64(revenueTotal)) * 100
	} else {
		profitMargin = 0
	}

	response := MetricsResponse{
		Status:  "1",
		Message: "Business Metrics Retrieved",
		Data: MetricsData{
			ExpenseToRevenueRatio:   expenseToRevenueRatio,
			RevenueLostDueToStock:   revenueLostDueToStock,
			ProfitMargin:            profitMargin,
			TopSellingSKU:           mostPopularSKU,
			InventoryCost:           inventoryCost,
		},
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
