package controllers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jorianom/go-recharges-ms/driver"
	"github.com/jorianom/go-recharges-ms/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = driver.GetCollection(driver.DB, "recharges")
var validate = validator.New()

func CreateRecharge(recharge models.Recharge) (*mongo.InsertOneResult, error) {
	// var user models.Recharge
	//use the validator library to validate required fields
	fmt.Println(recharge)
	if validationErr := validate.Struct(&recharge); validationErr != nil {
		fmt.Println("safa")
	}
	return nil, nil
	// // result, err := userCollection.InsertOne(context.TODO(), newRecharge)
	// if err != nil {
	// 	// panic(err)
	// 	fmt.Println(err)
	// }

	// return result, err

	// return c.Status(http.StatusCreated).JSON(models.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
