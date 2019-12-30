package school

import (
	"strings"

	"github.com/freitzzz/iped/model/school"
)

// This file contains model views representation for GET functionalities of schools collection

// GetAvailableSchoolsModelView is the model view representation
// for the available schools functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/schools.md#available-schools
type GetAvailableSchoolsModelView []struct {
	ID      int    `json:"id"`
	Acronym string `json:"acronym"`
	Name    string `json:"name"`
}

// GetDetailedSchoolInformationModelView is the model view representation
// for the detailed school information functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/schools.md#detailed-school-information
type GetDetailedSchoolInformationModelView struct {
	ID       int                   `json:"id"`
	Acronym  string                `json:"acronym"`
	Name     string                `json:"name"`
	Canteens []innerCanteensStruct `json:"canteens"`
}

type innerCanteensStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ToGetAvailableSchoolsModelView creates a GetAvailableSchoolsModelView using a slice of school models
func ToGetAvailableSchoolsModelView(schools []school.School) GetAvailableSchoolsModelView {
	modelview := make(GetAvailableSchoolsModelView, len(schools))

	for index, school := range schools {
		element := &modelview[index]
		element.ID = int(school.ID)
		element.Acronym = strings.ToUpper(school.Acronym)
		element.Name = school.Name
	}

	return modelview
}

// ToGetDetailedSchoolInformationModelView creates a GetDetailedSchoolInformationModelView using a school model
func ToGetDetailedSchoolInformationModelView(school school.School) GetDetailedSchoolInformationModelView {

	canteens := school.Canteens()

	modelviewCanteens := make([]innerCanteensStruct, len(canteens))

	for index, canteen := range canteens {
		element := &modelviewCanteens[index]
		element.ID = int(canteen.ID)
		element.Name = canteen.Name
	}

	modelview := GetDetailedSchoolInformationModelView{ID: int(school.ID), Name: school.Name, Acronym: strings.ToUpper(school.Acronym), Canteens: modelviewCanteens}

	return modelview
}
