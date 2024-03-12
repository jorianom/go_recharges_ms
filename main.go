package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorianom/go-recharges-ms/routes"
)

func main() {
	fmt.Println("Hello world")
	// client := driver.Connection()
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	router := mux.NewRouter()
	//	s := r.PathPrefix("/api").Subrouter()
	router.HandleFunc("/api/recharge", routes.RechargeHandler).Methods("POST")
	router.HandleFunc("/api/recharges/{id}", routes.HistoryHandler).Methods("GET")
	//
	router.HandleFunc("/api/method", routes.PostMethodHandler).Methods("POST")
	router.HandleFunc("/api/get-methods/{id}", routes.GetMethodHandler).Methods("GET")
	router.HandleFunc("/api/method/{id}", routes.UpdateMethodHandler).Methods("PUT")
	router.HandleFunc("/api/method/{id}", routes.DeleteMethodHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
