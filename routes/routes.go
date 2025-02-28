package routes

import (
	"fmt"
	"golang/controllers"
	"golang/middleware"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	fmt.Println("checking kalai")

	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	

	secured := r.PathPrefix("/api").Subrouter()
	secured.Use(middleware.AuthMiddleware)

	r.HandleFunc("/api/account_data", controllers.GetBalanceSheet).Methods("GET")

	// kpi start
	r.HandleFunc("/api/salesKPI1", controllers.GetYearToDateProfit).Methods("GET")

	r.HandleFunc("/api/GetExpensesVsSales", controllers.GetExpensesVsSales).Methods("GET")

	r.HandleFunc("/api/topFiveProduct", controllers.GetTopFiveProducts).Methods("GET")

	r.HandleFunc("/api/GetKPIData3", controllers.GetKPIData3).Methods("GET")

	r.HandleFunc("/api/FetchMetricsData", controllers.FetchMetricsData).Methods("GET")

	// fetch sales
	r.HandleFunc("/api/GetAllSales", controllers.GetAllSales).Methods("GET")

	// fetch purchases
	r.HandleFunc("/api/GetAllPurchases", controllers.GetAllPurchases).Methods("GET")

	// add sales
	r.HandleFunc("/api/sales/add", controllers.AddSales).Methods("POST")

	// add purchases
	r.HandleFunc("/api/purchases/add", controllers.AddPurchase).Methods("POST")

	// add balacesheet
	r.HandleFunc("/api/balance_sheet/add", controllers.AddBalanceSheetEntry).Methods("POST")

	
	return r
}
