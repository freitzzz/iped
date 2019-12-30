package school

// This file contains model views representation for POST functionalities of schools collection

// CreateNewSchoolModelView is the model view representation
// for the create new school functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/schools.md#create-a-new-school
type CreateNewSchoolModelView struct {
	Acronym  string `json:"acronym"`
	Name     string `json:"name"`
	Canteens []struct {
		Name     string             `json:"name"`
		Location postLocationStruct `json:"location"`
	} `json:"canteens"`
}

type postLocationStruct struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
