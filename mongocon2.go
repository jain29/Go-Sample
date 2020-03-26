package main

import (
	"context"
	"fmt"
	"log"
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
// 	projection := bson.M{"orderID": 1, "companyName": 1, "address1": 1, "address2": 1, "address3": 1, "city": 1, "state": 1, "country": 1, "zipCode": 1, "receipentPhone": 1, "carrierCode": 1, "classOfMail": 1, "ltlClass": 1, "userDefineFld": 1, "email": 1}
// 	err := coll.FindOne(ctx, model.BatchDB{OrderID: orderID}, options.FindOne().SetProjection(projection)).Decode(&order)
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(order)

// }

// func getAllPendingOrder(response http.ResponseWriter, request *http.Request) {

// 	filter := bson.M{"status": "P"}
// 	projection := bson.M{"orderID": 1, "companyName": 1, "address1": 1, "address2": 1, "address3": 1, "city": 1, "state": 1, "country": 1, "zipCode": 1, "carrierCode": 1, "classOfMail": 1}
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

func disconnect() error {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	defer disconnect()

	router := mux.NewRouter()
	//router.HandleFunc("/getAllPendingOrder", getAllPendingOrder).Methods("GET")
	router.HandleFunc("/order/{orderID}", getOrder(client)).Methods("GET")
	http.ListenAndServe(":8080", router)
}
