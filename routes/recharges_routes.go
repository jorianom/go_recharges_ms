package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jorianom/go-recharges-ms/driver"
	"github.com/jorianom/go-recharges-ms/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var rechargeCollection *mongo.Collection = driver.GetCollection(driver.DB, "recharges")
var validate = validator.New()

func RechargeHandler(w http.ResponseWriter, r *http.Request) {
	var recharge models.Recharge

	json.NewDecoder(r.Body).Decode(&recharge)
	err := validate.Struct(&recharge)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field() + " is a requerid field")
			}
			return
		}
	}

	result, err := rechargeCollection.InsertOne(context.TODO(), &recharge)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&result)
	//w.Write([]byte("Hello World 2"))
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// var recharge models.Recharge
	var recharges []models.Recharge
	defer cancel()
	params := mux.Vars(r)
	fmt.Println(params)
	filter := bson.D{{Key: "user", Value: params["id"]}}
	// Retrieves documents that match the query filer
	results, err := rechargeCollection.Find(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.Recharge
		if err = results.Decode(&singleUser); err != nil {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		recharges = append(recharges, singleUser)
	}
	json.NewEncoder(w).Encode(&recharges)
}
