package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorianom/go-recharges-ms/driver"
	"github.com/jorianom/go-recharges-ms/routes"
)

func main() {
	fmt.Println("Hello world")
	client := driver.Connection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	router := mux.NewRouter()
	router.HandleFunc("/api/recharge", routes.RechargeHandler).Methods("POST")
	router.HandleFunc("/api/recharges/{id}", routes.HistoryHandler).Methods("GET")
	router.HandleFunc("/api/method", routes.HistoryHandler).Methods("POST")
	router.HandleFunc("/api/my-methods", routes.HistoryHandler).Methods("GET")

	http.ListenAndServe(":3000", router)
}
