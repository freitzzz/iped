package dish

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/freitzzz/iped/model/customerror"
)

// Dish is a model for what a person can choose to eat in canteen
// A UML overview of this model can be found at https://github.com/freitzzz/iped-documentation/wiki/Architecture#models-structure
type Dish struct {
	gorm.Model
	// MenuID field is a foreign key so gorm is able to map a has may association between menu and dish
	MenuID      uint `gorm:"type:int REFERENCES menus(id) ON UPDATE CASCADE ON DELETE CASCADE"`
	Type        DishType
	Description string
}

// New initializes a Dish model using a dish type and a description
// A FieldError is returned either if the dish description is empty or the dish type isn't valid
func New(Type int, Description string) (Dish, *customerror.FieldError) {

	dish := Dish{gorm.Model{}, 0, DishType(Type), Description}

	err := grantValidDishType(Type)

	if err != nil {
		return dish, err
	}

	err = grantDescriptionIsNotEmpty(Description)

	return dish, err
}

// Equals compares equality between two dishes
// A dish proves true equality to other dish if both dish types and description are equal
func (dish Dish) Equals(comparingDish Dish) bool {
	return dish.Type == comparingDish.Type && dish.Description == comparingDish.Description
}

// This function grants that a dish description is not empty, and if not returns a FieldError
func grantDescriptionIsNotEmpty(description string) *customerror.FieldError {

	var err *customerror.FieldError

	if len(strings.TrimSpace(description)) == 0 {
		err = &customerror.FieldError{Field: "description", Model: "dish", Explanation: "dish description cannot be an empty string"}
	}

	return err
}

// This function grants that a dish type is valid, and if not returns a FieldError
// See [DishType.Validate] for validation logic
func grantValidDishType(dishtype int) *customerror.FieldError {

	var err *customerror.FieldError

	if !Validate(dishtype) {
		err = &customerror.FieldError{Field: "dishtype", Model: "dish", Explanation: "dish type is not valid"}
	}

	return err
}
