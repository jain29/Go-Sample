package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	model "github.com/jain29/Go-Sample/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//func getOrder(response http.ResponseWriter, request *http.Request) {
func getOrder(client *mongo.Client) func(response http.ResponseWriter, request *http.Request) {

	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json")
		params := mux.Vars(request)
		orderID, _ := params["orderID"]
		var order model.BatchDB
		coll := client.Database("local").Collection("BatchDB")
		ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
		projection := bson.M{"orderID": 1, "companyName": 1, "address1": 1, "address2": 1, "address3": 1, "city": 1, "state": 1, "country": 1, "zipCode": 1, "receipentPhone": 1, "carrierCode": 1, "classOfMail": 1, "ltlClass": 1, "userDefineFld": 1, "email": 1}
		opts := options.Find().SetProjection(projection)
		err := coll.FindOne(ctx, model.BatchDB{OrderID: orderID}, opts).Decode(&order)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(order)
	}
}
