package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// func getOrder(response http.ResponseWriter, request *http.Request) {

// 	response.Header().Set("content-type", "application/json")
// 	params := mux.Vars(request)
// 	orderID, _ := params["orderID"]
// 	var order model.BatchDB
// 	coll := client.Database("local").Collection("BatchDB")
// 	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

// 	err := coll.FindOne(ctx, model.BatchDB{OrderID: orderID}).Decode(&order)
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(order)

// }

// func getAllPendingOrder(response http.ResponseWriter, request *http.Request) {

// 	filter := bson.M{"status": "P"}
// 	//filter := bson.M{}
// 	//projection := bson.M{"orderID": 1, "CompanyName": 1,"Address1"}
// 	projection := bson.M{"status": 0}
// 	coll := client.Database("local").Collection("BatchDB")
// 	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

// 	opts := options.Find().SetProjection(projection)

// 	cursor, err := coll.Find(ctx, filter, opts)

// 	response.Header().Set("content-type", "application/json")

// 	var orders []model.BatchDB
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var order model.BatchDB
// 		cursor.Decode(&order)
// 		orders = append(orders, order)
// 	}
// 	json.NewEncoder(response).Encode(orders)

// }

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	//router.HandleFunc("/getAllPendingOrder", getAllPendingOrder).Methods("GET")
	router.HandleFunc("/order/{orderID}", getOrder).Methods("GET")
	http.ListenAndServe(":8080", router)
}
