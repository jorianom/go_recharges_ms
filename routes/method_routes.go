package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jorianom/go-recharges-ms/driver"
	"github.com/jorianom/go-recharges-ms/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var methodsCollection *mongo.Collection = driver.GetCollection(driver.DB, "methods")
var validateMethod = validator.New()

func PostMethodHandler(w http.ResponseWriter, r *http.Request) {
	var method models.Method
	json.NewDecoder(r.Body).Decode(&method)
	err := validateMethod.Struct(&method)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field() + " is a requerid field")
			}
			return
		}
	}

	result, err := methodsCollection.InsertOne(context.TODO(), &method)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	response := models.MethodResponse{
		Status: http.StatusAccepted,
		Method: method,
		Id:     result.InsertedID.(primitive.ObjectID),
	}

	json.NewEncoder(w).Encode(&response)
}

func GetMethodHandler(w http.ResponseWriter, r *http.Request) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var methods []models.Method
	// defer cancel()
	params := mux.Vars(r)
	fmt.Println(params)
	filter := bson.D{{Key: "user", Value: params["id"]}}
	// Retrieves documents that match the query filer
	results, err := methodsCollection.Find(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	if err = results.All(context.TODO(), &methods); err != nil {
		panic(err)
	}
	// defer results.Close(ctx)
	// for results.Next(ctx) {
	// 	var singleMethod models.Method
	// 	if err = results.Decode(&singleMethod); err != nil {

	// 		w.WriteHeader(http.StatusBadRequest)
	// 		w.Write([]byte(err.Error()))
	// 	}
	// 	singleMethod.Id = results.
	// 	methods = append(methods, singleMethod)
	// }
	json.NewEncoder(w).Encode(&methods)
}

type Response struct {
	Message string `json:"message"`
}

func UpdateMethodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
	var method models.Method
	json.NewDecoder(r.Body).Decode(&method)
	err := validateMethod.Struct(&method)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field() + " is a requerid field")
			}
			return
		}
	}

	params := mux.Vars(r)
	fmt.Println(params)
	filter := bson.D{{Key: "_id", Value: params["id"]}}
	update := bson.D{{Key: "$set", Value: method}}
	result, err := methodsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil || result == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Printf("Error al actualizar el documento: %v", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
	// var response models.MethodResponse
	// if result != nil {

	// 	response = models.MethodResponse{
	// 		Status: http.StatusAccepted,
	// 		Method: method,
	// 	}

	// } else {

	// 	response = models.MethodResponse{
	// 		Status: http.StatusBadRequest,
	// 	}

	// }
	res := Response{
		Message: "updated " + string(rune(result.ModifiedCount)) + " documents.",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	// json.NewEncoder(w).Encode(&response)
}
func DeleteMethodHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	filter := bson.D{{"_id", params["id"]}}
	result, err := methodsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&result)
}
