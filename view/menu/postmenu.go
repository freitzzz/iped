package menu

// This file contains model views representation for POST functionalities of menus collection

// CreateNewMenuModelView is the model view representation
// for the create new menu functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/menus.md#create-a-new-menu
type CreateNewMenuModelView struct {
	Type   string `json:"type"`
	Dishes []struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"dishes"`
}
