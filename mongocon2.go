package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var client *mongo.Client

// type BatchDB struct {

// 	OrderID    	    string `json:"OrderID,omitempty" bson:"OrderID,omitempty"`
// 	CompanyName 	string `json:"CompanyName,omitempty" bson:"CompanyName,omitempty"`
// 	Address1        string `json:"Address1,omitempty" bson:"Address1,omitempty"`
// 	Address2        string `json:"Address2,omitempty" bson:"Address2,omitempty"`
// 	Address3        string `json:"Address3,omitempty" bson:"Address3,omitempty"`
// 	City       		string `json:"City,omitempty" bson:"City,omitempty"`
// 	State       	string `json:"State,omitempty" bson:"State,omitempty"`
// 	Country       	string `json:"Country,omitempty" bson:"Country,omitempty"`
// 	ZipCode       	string `json:"ZipCode,omitempty" bson:"ZipCode,omitempty"`	
// 	CarrierCode     string `json:"CarrierCode,omitempty" bson:"CarrierCode,omitempty"`
// 	ClassOfMail     string `json:"ClassOfMail,omitempty" bson:"ClassOfMail,omitempty"`
// 	status     string `json:"status,omitempty" bson:"status,omitempty"`
// }

func getAllPendingOrder(response http.ResponseWriter, request *http.Request) {

	filter := bson.M{"status": "P"}
	projection := bson.M{"status":"0"}	
	coll := client.Database("local").Collection("BatchDB")
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	opts := options.Find().SetProjection(projection)

	cursor, err := coll.Find(ctx, filter, opts)

	response.Header().Set("content-type", "application/json")

	var orders []BatchDB
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var order BatchDB
		cursor.Decode(&order)
		orders = append(orders, order)
	}
	json.NewEncoder(response).Encode(orders)

}


func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/getAllPendingOrder", getAllPendingOrder).Methods("GET")	
	http.ListenAndServe(":8080", router)
}
