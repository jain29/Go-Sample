package model

type BatchDB struct {
	CustomerID          string `json:"customerID,omitempty" bson:"customerID,omitempty"`
	FileName            string `json:"fileName,omitempty" bson:"fileName,omitempty"`
	OrderID             string `json:"orderID,omitempty" bson:"orderID,omitempty"`
	CompanyName         string `json:"companyName,omitempty" bson:"companyName,omitempty"`
	Address1            string `json:"address1,omitempty" bson:"address1,omitempty"`
	Address2            string `json:"address2,omitempty" bson:"address2,omitempty"`
	Address3            string `json:"address3,omitempty" bson:"address3,omitempty"`
	City                string `json:"city,omitempty" bson:"city,omitempty"`
	State               string `json:"state,omitempty" bson:"state,omitempty"`
	Country             string `json:"country,omitempty" bson:"country,omitempty"`
	ZipCode             string `json:"zipCode,omitempty" bson:"zipCode,omitempty"`
	ReceipentPhone      string `json:"receipentPhone,omitempty" bson:"receipentPhone,omitempty"`
	CarrierCode         string `json:"carrierCode,omitempty" bson:"carrierCode,omitempty"`
	ClassOfMail         string `json:"classOfMail,omitempty" bson:"classOfMail,omitempty"`
	LtlClass            string `json:"ltlClass,omitempty" bson:"ltlClass,omitempty"`
	UserDefineFld       string `json:"userDefineFld,omitempty" bson:"userDefineFld,omitempty"`
	Email               string `json:"email,omitempty" bson:"email,omitempty"`
	TrackingNumber      string `json:"trackingNumber,omitempty" bson:"trackingNumber,omitempty"`
	TotalShipmentCharge string `json:"totalShipmentCharge,omitempty" bson:"totalShipmentCharge,omitempty"`
	ShipmentDate        string `json:"shipmentDate,omitempty" bson:"shipmentDate,omitempty"`
	Status              string `json:"status,omitempty" bson:"status,omitempty"`
	ErrorDetail         string `json:"errorDetail,omitempty" bson:"errorDetail,omitempty"`
	FileProcessingDate  string `json:"fileProcessingDate,omitempty" bson:"fileProcessingDate,omitempty"`
}
