package http

import (
	"net/http"

	"github.com/Strke12i/Go_project/Go_project/adapter/http/actuator"
	"github.com/Strke12i/Go_project/Go_project/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Init é a Funçõa inicial que chamara as demais func
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)
	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
