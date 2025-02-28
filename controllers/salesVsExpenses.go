package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang/config" 
)


type MonthlyData struct {
	Month       string `json:"month"`
	TotalSales  int    `json:"total_sales"`
	TotalExpenses int  `json:"total_expenses"`
	NetProfit   int    `json:"net_profit"`
}

type ExpensesVsSalesResponse struct {
	Status      string        `json:"status"`
	Message     string        `json:"message"`
	Data        []MonthlyData `json:"data"`
}

func GetExpensesVsSales(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching Monthly Expenses vs Sales...")


	rows, err := config.DB.Query(`SELECT DATE_FORMAT(s.ordered_date, '%Y-%m') AS month,COALESCE(SUM(s.total_amount), 0) AS total_sales,COALESCE(SUM(p.total_amount), 0) AS total_expenses
		FROM sales s LEFT JOIN  purchases p ON DATE_FORMAT(s.ordered_date, '%Y-%m') = DATE_FORMAT(p.purchase_date, '%Y-%m')
		GROUP BY  DATE_FORMAT(s.ordered_date, '%Y-%m')
		ORDER BY month DESC;`)

	if err != nil {
		http.Error(w, "Database error fetching expenses vs sales", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var data []MonthlyData


	for rows.Next() {
		var record MonthlyData
		if err := rows.Scan(&record.Month, &record.TotalSales, &record.TotalExpenses); err != nil {
			http.Error(w, "Error scanning expenses vs sales data", http.StatusInternalServerError)
			return
		}

		record.NetProfit = record.TotalSales - record.TotalExpenses
		data = append(data, record)
	}


	response := ExpensesVsSalesResponse{
		Status:  "1",
		Message: "Monthly Expenses vs Sales Fetched",
		Data:    data,
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
