package routes

import (
	"context"
	"encoding/json"
	"fmt"
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

	var response models.MethodResponse
	var method models.Method
	json.NewDecoder(r.Body).Decode(&method)
	err := validateMethod.Struct(&method)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field() + " is a requerid field")

				response = models.MethodResponse{
					Message: err.Field() + " is a requerid field",
					Status:  http.StatusBadRequest,
				}
			}

		}
	} else {
		result, err := methodsCollection.InsertOne(context.TODO(), &method)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		response = models.MethodResponse{
			Status: http.StatusAccepted,
			Method: method,
			Id:     result.InsertedID.(primitive.ObjectID),
		}
	}

	json.NewEncoder(w).Encode(&response)
}

type ResponseGet struct {
	Message string          `json:"message"`
	Status  int             `json:"status"`
	Data    []models.Method `json:"data"`
}

func GetMethodHandler(w http.ResponseWriter, r *http.Request) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var methods []models.Method
	var res ResponseGet
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
	if len(methods) == 0 {
		res = ResponseGet{
			Message: "No se encontraron registros " + params["id"],
			Status:  http.StatusAccepted,
		}
	} else {
		res = ResponseGet{
			Message: "Registros obtenidos correctamente " + params["id"],
			Status:  http.StatusAccepted,
			Data:    methods,
		}
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
	json.NewEncoder(w).Encode(&res)
}

type Response struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Data    models.Method `json:"data"`
}

func UpdateMethodHandler(w http.ResponseWriter, r *http.Request) {
	var method models.Method
	var res Response
	json.NewDecoder(r.Body).Decode(&method)
	err := validateMethod.Struct(&method)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field() + " is a requerid field")

				res = Response{
					Message: err.Field() + " is a requerid field",
					Status:  http.StatusBadRequest,
				}
			}
		}
	} else {
		params := mux.Vars(r)

		objId, _ := primitive.ObjectIDFromHex(params["id"])
		fmt.Println(params)
		// filter := bson.D{{Key: "id", Value: objId}}
		//update := bson.D{{Key: "$set", Value: method}}
		update := bson.M{"user": method.User, "name": method.Name, "titular": method.Titular, "duedate": method.Duedate, "number": method.Number, "type": method.Type, "sucursal": method.Sucursal}
		result, err := methodsCollection.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		fmt.Println("Hello world 3")
		// var updatedMethod []models.Method
		fmt.Println(result.MatchedCount)
		if result.MatchedCount == 1 {
			// filter := bson.D{{Key: "user", Value: params["id"]}}
			// results, err := methodsCollection.Find(context.TODO(), filter)
			// if err != nil {
			// 	w.WriteHeader(http.StatusBadRequest)
			// 	w.Write([]byte(err.Error()))
			// }
			// if err = results.All(context.TODO(), &updatedMethod); err != nil {
			// 	res = Response{
			// 		Message: "error " + err.Error(),
			// 		Status:  http.StatusInternalServerError,
			// 	}
			// } else {
			// 	res = Response{
			// 		Status: http.StatusAccepted,
			// 		Data:   updatedMethod,
			// 	}
			// }

			res = Response{
				Message: "Actualizado correctamente:  " + params["id"],
				Status:  http.StatusAccepted,
				Data:    method,
			}
		} else {
			res = Response{
				Message: "Id no encontrado",
				Status:  http.StatusBadRequest,
			}
		}
		fmt.Println("Hello world 4")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
	// json.NewEncoder(w).Encode(&response)
}
func DeleteMethodHandler(w http.ResponseWriter, r *http.Request) {
	var res Response
	params := mux.Vars(r)
	fmt.Println(params)
	// filter := bson.D{{Key: "_id", Value: params["id"]}}

	objId, _ := primitive.ObjectIDFromHex(params["id"])
	result, err := methodsCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	if result.DeletedCount == 1 {
		res = Response{
			Message: "Eliminado correctamente:  " + params["id"],
			Status:  http.StatusAccepted,
		}
	} else {
		res = Response{
			Message: "Id no encontrado" + params["id"],
			Status:  http.StatusAccepted,
		}
	}
	json.NewEncoder(w).Encode(&res)
}
