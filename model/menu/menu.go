package menu

import (
	"github.com/freitzzz/iped/model/customerror"
	"github.com/freitzzz/iped/model/dish"
	"github.com/jinzhu/gorm"
)

// Menu is a model that contains a set of dishes available at either lunch or dinner
// A UML overview of this model can be found at https://github.com/freitzzz/iped-documentation/wiki/Architecture#models-structure
type Menu struct {
	gorm.Model
	MenuEntryID uint `gorm:"type:int REFERENCES menu_entries(id) ON UPDATE CASCADE ON DELETE CASCADE"`
	Type        MenuType
	DishesSlice []dish.Dish
}

// New initializes a Menu model using a menu type and a set of dishes
// A FieldError is returned either if the menu type is invalid, no dishes were provided or the dishes are not unique
func New(Type int, Dishes []dish.Dish) (Menu, *customerror.FieldError) {

	menu := Menu{gorm.Model{}, 0, MenuType(Type), Dishes}

	err := grantValidMenuType(Type)

	if err != nil {
		return menu, err
	}

	err = grantThatAtLeastOneDishWasProvided(Dishes)

	if err != nil {
		return menu, err
	}

	err = grantNoDuplicatedDishesExist(Dishes)

	return menu, err
}

// Dishes returns the available dishes on a menu as a slice
// The returned slice has different reference of the one in Menu struct
// In order to prevent modifications
func (menu Menu) Dishes() []dish.Dish {
	availableDishes := make([]dish.Dish, len(menu.DishesSlice))

	copy(availableDishes, menu.DishesSlice)

	return availableDishes
}

// This function grants that a menu type is valid, and if not returns a FieldError
// See [MenuType.Validate] for validation logic
func grantValidMenuType(menutype int) *customerror.FieldError {

	var err *customerror.FieldError

	if !Validate(menutype) {
		err = &customerror.FieldError{Field: "menutype", Model: "menu", Explanation: "specified menu type is not valid"}
	}

	return err
}

// This function grants that at least one dish is provided in given dish slice
// If the given dish slice is nil or empty a FieldError is returned
func grantThatAtLeastOneDishWasProvided(dishes []dish.Dish) *customerror.FieldError {

	var err *customerror.FieldError

	if dishes == nil || len(dishes) == 0 {
		err = &customerror.FieldError{Field: "dishes", Model: "menu", Explanation: "menu requires at least one dish to be provided"}
	}

	return err
}

// This function grants that all dishes given in a slice are unique
// If a dish proves equality to any other dish in the slice, a FieldError is returned
func grantNoDuplicatedDishesExist(dishes []dish.Dish) *customerror.FieldError {

	var err *customerror.FieldError

	unique := true
	dishesLength := len(dishes)
	i := 0

	for i < dishesLength {
		j := i + 1
		for j < dishesLength {
			unique = !dishes[i].Equals(dishes[j])
			j++
			if !unique {
				i = dishesLength
				j = i
			}
		}
		i++
	}

	if !unique {
		err = &customerror.FieldError{Field: "dishes", Model: "menu", Explanation: "menu cannot have duplicated dishes"}
	}

	return err
}
