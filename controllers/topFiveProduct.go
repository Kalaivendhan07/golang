package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang/config" 
)


type TopProduct struct {
	ProductName string `json:"product_name"`
	SKU         string `json:"sku"`
	UnitsSold   int    `json:"units_sold"`
	TotalSales  int    `json:"total_sales"`
}


type TopProductsResponse struct {
	Status      string        `json:"status"`
	Message     string        `json:"message"`
	TopProducts []TopProduct  `json:"top_products"`
}


func GetTopFiveProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching Top Five Best-Selling Products...")


	rows, err := config.DB.Query(`SELECT product, sku, COUNT(*) as units_sold, SUM(total_amount) as total_sales FROM sales 
		WHERE YEAR(ordered_date) = YEAR(CURDATE()) AND MONTH(ordered_date) = MONTH(CURDATE()) GROUP BY product, sku ORDER BY units_sold DESC
		LIMIT 5
	`)

	if err != nil {
		http.Error(w, "Database error fetching top products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var topProducts []TopProduct

	for rows.Next() {
		var tp TopProduct
		if err := rows.Scan(&tp.ProductName, &tp.SKU, &tp.UnitsSold, &tp.TotalSales); err != nil {
			http.Error(w, "Error scanning top products", http.StatusInternalServerError)
			return
		}
		topProducts = append(topProducts, tp)
	}


	response := TopProductsResponse{
		Status:      "1",
		Message:     "Top 5 Best-Selling Products Fetched",
		TopProducts: topProducts,
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
