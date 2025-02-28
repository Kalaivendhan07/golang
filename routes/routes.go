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
	
	return r
}
