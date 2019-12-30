package menu

import (
	"testing"

	"github.com/freitzzz/iped/model/dish"
)

func TestUnexistingMenuTypeReturnError(t *testing.T) {

	_dish := dish.Dish{Type: 0, Description: "Fried Noodles"}

	dishes := []dish.Dish{_dish}

	_, err := New(-1, dishes)

	if err == nil {
		t.Error("Menu initialization should have returned an error as there is no menu type for the value -1")
	}

	if err.Field != "menutype" {
		t.Error("Even though that menu initialization returned an error, the error should have been caused by the field menutype")
	}
}

func TestExistingMenuTypeDoesNotReturnError(t *testing.T) {
	_dish := dish.Dish{Type: 0, Description: "Fried Noodles"}

	dishes := []dish.Dish{_dish}

	_, err := New(0, dishes)

	if err != nil {
		t.Errorf("Menu initilization should have been successful but got %s", err)
	}
}

func TestNilDishListReturnError(t *testing.T) {

	_, err := New(0, nil)

	if err == nil {
		t.Error("Menu initilization should have returned an error but got nil")
	}

	if err.Field != "dishes" {
		t.Error("Even though that menu initialization returned an error, the error should have been caused by the field dishes")
	}
}

func TestEmptyDishListReturnError(t *testing.T) {
	_, err := New(0, []dish.Dish{})

	if err == nil {
		t.Error("Menu initilization should have returned an error but got nil")
	}

	if err.Field != "dishes" {
		t.Error("Even though that menu initialization returned an error, the error should have been caused by the field dishes")
	}
}

func TestDishListWithEqualDishesReturnError(t *testing.T) {
	_dish := dish.Dish{Type: 0, Description: "Fried Noodles"}

	dishes := []dish.Dish{_dish, _dish}

	_, err := New(0, dishes)

	if err == nil {
		t.Error("Menu initilization should have returned an error but got nil")
	}

	if err.Field != "dishes" {
		t.Error("Even though that menu initialization returned an error, the error should have been caused by the field dishes")
	}

	_dish2 := dish.Dish{Type: 1, Description: "Fried Noodles"}

	dishes = []dish.Dish{_dish, _dish2, _dish2}

	if err == nil {
		t.Error("Menu initilization should have returned an error but got nil")
	}

	if err.Field != "dishes" {
		t.Error("Even though that menu initialization returned an error, the error should have been caused by the field dishes")
	}
}

func TestExistingMenuTypeAndValidDishListDoesNotReturnError(t *testing.T) {
	_dish := dish.Dish{Type: 0, Description: "Fried Noodles"}

	dishes := []dish.Dish{_dish}

	_, err := New(0, dishes)

	if err != nil {
		t.Errorf("Menu initilization should have been successful but got %s", err)
	}
}

func TestDishesMethodReturnsSliceWithDifferentReference(t *testing.T) {
	_dish := dish.Dish{Type: 0, Description: "Fried Noodles"}

	dishes := []dish.Dish{_dish}

	menu, _ := New(0, dishes)

	availableDishes := menu.Dishes()

	// This verification is to grant that the returned available dishes slice length is 1

	if lenab := len(availableDishes); lenab != 1 {
		t.Errorf("The length of availableDishes slice should be 1 but got: %d", lenab)
	}

	// If the length of the slice is the same as the capacity the slice was successfuly copied from the original slice

	if capb := cap(availableDishes); capb != 1 {
		t.Errorf("The capacitiy of availableDishes should be the same as its length (1) but got %d", capb)
	}

	// If we add a new dish to the the returned slice,
	// it should not modify the slice pointed on the menu struct

	availableDishes = append(availableDishes, _dish)

	if lenam := len(availableDishes); lenam != 2 {
		t.Errorf("The length of availableDishes slice should now be 2 but got: %d", lenam)
	}

	availableDishesAfterModification := menu.Dishes()

	if lenaam := len(availableDishesAfterModification); lenaam != 1 {
		t.Errorf("The length of availableDishesAfterModification slice should be 1 but got: %d", lenaam)
	}

	if capb := cap(availableDishesAfterModification); capb != 1 {
		t.Errorf("The capacitiy of availableDishesAfterModification should be the same as its length (1) but got %d", capb)
	}
}

func TestDishesMethodReturnsDishesPassedOnInitialization(t *testing.T) {
	_dish := dish.Dish{Type: 0, Description: "Fried Noodles"}

	dishes := []dish.Dish{_dish}

	menu, _ := New(0, dishes)

	availableDishes := menu.Dishes()

	if lena := len(availableDishes); lena != 1 {
		t.Errorf("Available dishes slice length returned by Dishes method should be 1 but got: %d", lena)
	}

	equalDishType := _dish.Type == availableDishes[0].Type

	if !equalDishType {
		t.Errorf("Available dish type returned by Dishes method is not equal to the one passed on initialization")
	}

	equalDishDescription := _dish.Description == availableDishes[0].Description

	if !equalDishDescription {
		t.Errorf("Available dish description returned by Dishes method is not equal to the one passed on initialization")
	}

}
