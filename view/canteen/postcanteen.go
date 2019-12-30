package canteen

// This file contains model views representation for POST functionalities of canteens collection

// CreateNewCanteenModelView is the model view representation
// for the create new canteen functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/canteens.md#create-a-new-canteen
type CreateNewCanteenModelView struct {
	Name     string             `json:"name"`
	Location postlocationStruct `json:"location"`
}

type postlocationStruct struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
