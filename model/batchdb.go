package model

type BatchDB struct {
	OrderID     string `json:"orderID,omitempty" bson:"orderID,omitempty"`
	CompanyName string `json:"companyName,omitempty" bson:"companyName,omitempty"`
	Address1    string `json:"address1,omitempty" bson:"address1,omitempty"`
	Address2    string `json:"address2,omitempty" bson:"address2,omitempty"`
	Address3    string `json:"address3,omitempty" bson:"address3,omitempty"`
	City        string `json:"city,omitempty" bson:"city,omitempty"`
	State       string `json:"state,omitempty" bson:"state,omitempty"`
	Country     string `json:"country,omitempty" bson:"country,omitempty"`
	ZipCode     string `json:"zipCode,omitempty" bson:"zipCode,omitempty"`
	CarrierCode string `json:"carrierCode,omitempty" bson:"carrierCode,omitempty"`
	ClassOfMail string `json:"classOfMail,omitempty" bson:"classOfMail,omitempty"`
	Status      string `json:"status,omitempty" bson:"status,omitempty"`
}
