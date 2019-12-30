package menu

import (
	"github.com/freitzzz/iped/model/menu"
)

// This file contains model views representation for GET functionalities of menus collection

// GetAvailableMenusModelView is the model view representation
// for the available menus functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/menus.md#available-menus
type GetAvailableMenusModelView []struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// GetDetailedMenuInformationModelView is the model view representation
// for the detailed menu information functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/menus.md#detailed-menu-information
type GetDetailedMenuInformationModelView struct {
	ID     int                 `json:"id"`
	Type   string              `json:"type"`
	Dishes []innerDishesStruct `json:"dishes"`
}

type innerDishesStruct struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// ToGetAvailableMenusModelView creates a GetAvailableMenusModelView using a slice of menu models
func ToGetAvailableMenusModelView(menus []menu.Menu) GetAvailableMenusModelView {
	modelview := make(GetAvailableMenusModelView, len(menus))

	for index, menu := range menus {
		element := &modelview[index]
		element.ID = int(menu.ID)
		element.Type = menu.Type.String()
	}

	return modelview
}

// ToGetDetailedMenuInformationModelView creates a GetDetailedMenuInformationModelView using a menu model
func ToGetDetailedMenuInformationModelView(menu menu.Menu) GetDetailedMenuInformationModelView {

	dishes := menu.Dishes()

	modelviewDishes := make([]innerDishesStruct, len(dishes))

	for index, dish := range dishes {
		element := &modelviewDishes[index]
		element.ID = int(dish.ID)
		element.Description = dish.Description
		element.Type = dish.Type.String()
	}

	modelview := GetDetailedMenuInformationModelView{ID: int(menu.ID), Type: menu.Type.String(), Dishes: modelviewDishes}

	return modelview
}
