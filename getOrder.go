package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	model "github.com/jain29/Go-Sample/model"
)

func getOrder(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	orderID, _ := params["orderID"]
	var order model.BatchDB
	coll := client.Database("local").Collection("BatchDB")
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	err := coll.FindOne(ctx, model.BatchDB{OrderID: orderID}).Decode(&order)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(order)

}
