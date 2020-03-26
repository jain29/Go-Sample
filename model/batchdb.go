package model

type BatchDB struct {

	OrderID    	    string `json:"OrderID,omitempty" bson:"OrderID,omitempty"`
	CompanyName 	string `json:"CompanyName,omitempty" bson:"CompanyName,omitempty"`
	Address1        string `json:"Address1,omitempty" bson:"Address1,omitempty"`
	Address2        string `json:"Address2,omitempty" bson:"Address2,omitempty"`
	Address3        string `json:"Address3,omitempty" bson:"Address3,omitempty"`
	City       		string `json:"City,omitempty" bson:"City,omitempty"`
	State       	string `json:"State,omitempty" bson:"State,omitempty"`
	Country       	string `json:"Country,omitempty" bson:"Country,omitempty"`
	ZipCode       	string `json:"ZipCode,omitempty" bson:"ZipCode,omitempty"`	
	CarrierCode     string `json:"CarrierCode,omitempty" bson:"CarrierCode,omitempty"`
	ClassOfMail     string `json:"ClassOfMail,omitempty" bson:"ClassOfMail,omitempty"`
	status     		string `json:"status,omitempty" bson:"status,omitempty"`
}