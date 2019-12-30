package canteen

import (
	"testing"

	"github.com/freitzzz/iped/model/dish"
	"github.com/freitzzz/iped/model/menu"
)

func TestEmptyCanteenNameReturnError(t *testing.T) {
	location := Location{}
	_, err := New("", location)

	if err == nil {
		t.Error("Canteen initilization should have returned an error but got nil")
	}

	if err.Field != "name" {
		t.Error("Even though that canteen initialization returned an error, the error should have been caused by the field name")
	}
}

func TestCanteenNameWithOnlySpacesReturnError(t *testing.T) {
	location := Location{}
	_, err := New(" ", location)

	if err == nil {
		t.Error("Canteen initilization should have returned an error but got nil")
	}

	if err.Field != "name" {
		t.Error("Even though that canteen initialization returned an error, the error should have been caused by the field name")
	}
}

func TestNotEmptyCanteenNameDoesNotReturnError(t *testing.T) {
	location := Location{}
	_, err := New("Cantina do H", location)

	if err != nil {
		t.Errorf("Canteen initilization should have been successful but got %s", err)
	}
}

func TestAvailableMenusMethodReturnsEmptySliceIfNoMenusAreAvailable(t *testing.T) {
	location := Location{}
	_canteen, _ := New("Cantina do H", location)

	availableMenus := _canteen.AvailableMenus()

	if lenab := len(availableMenus); lenab != 0 {
		t.Errorf("The length of availableMenus slice should be 0 but got: %d", lenab)
	}
}

func TestAvailableMenusMethodReturnsSliceWithDifferentReference(t *testing.T) {
	location := Location{}
	_canteen, _ := New("Cantina do H", location)

	availableMenus := _canteen.AvailableMenus()

	// This verification is to grant that the returned available menus slice length is 0

	if lenab := len(availableMenus); lenab != 0 {
		t.Errorf("The length of availableMenus slice should be 0 but got: %d", lenab)
	}

	// If the length of the slice is the same as the capacity the slice was successfuly copied from the original slice

	if capb := cap(availableMenus); capb != 0 {
		t.Errorf("The capacitiy of availableMenus should be the same as its length (0) but got %d", capb)
	}

	// If we add a new menu to the the returned slice,
	// it should not modify the slice pointed on the canteen struct

	_dish, _ := dish.New(0, "Fried Noodles")

	_menu, _ := menu.New(0, []dish.Dish{_dish})

	availableMenus = append(availableMenus, _menu)

	if lenam := len(availableMenus); lenam != 1 {
		t.Errorf("The length of availableMenus slice should now be 1 but got: %d", lenam)
	}

	availableMenusAfterModification := _canteen.AvailableMenus()

	if lenaam := len(availableMenusAfterModification); lenaam != 0 {
		t.Errorf("The length of availableMenusAfterModification slice should be 0 but got: %d", lenaam)
	}

	if capb := cap(availableMenusAfterModification); capb != 0 {
		t.Errorf("The capacitiy of availableMenusAfterModification should be the same as its length (0) but got %d", capb)
	}

}

func TestAddTodayMethodUpdatesAvailableMenus(t *testing.T) {
	location := Location{}
	_canteen, _ := New("Cantina do H", location)

	availableMenus := _canteen.AvailableMenus()

	// This verification is to grant that the returned available menus slice length is 0

	if lenab := len(availableMenus); lenab != 0 {
		t.Errorf("The length of availableMenus slice should be 0 but got: %d", lenab)
	}

	// If we add a new menu to today menus
	// It should update the available menus

	_dish, _ := dish.New(0, "Fried Noodles")

	_menu, _ := menu.New(0, []dish.Dish{_dish})

	_canteen.AddTodayMenu(_menu)

	availableMenus = _canteen.AvailableMenus()

	if lenam := len(availableMenus); lenam != 1 {
		t.Errorf("The length of availableMenus slice should now be 1 but got: %d", lenam)
	}
}

func TestAddTodayMethodReturnsErrorIfMenuOfTheSameTypeAlreadyExists(t *testing.T) {
	location := Location{}
	_canteen, _ := New("Cantina do H", location)

	availableMenus := _canteen.AvailableMenus()

	// This verification is to grant that the returned available menus slice length is 0

	if lenab := len(availableMenus); lenab != 0 {
		t.Errorf("The length of availableMenus slice should be 0 but got: %d", lenab)
	}

	// If we add a new menu to today menus
	// It should update the available menus

	_dish, _ := dish.New(0, "Fried Noodles")

	_menu, _ := menu.New(0, []dish.Dish{_dish})

	_canteen.AddTodayMenu(_menu)

	availableMenus = _canteen.AvailableMenus()

	if lenam := len(availableMenus); lenam != 1 {
		t.Errorf("The length of availableMenus slice should now be 1 but got: %d", lenam)
	}

	_differentTypeMenu, _ := menu.New(1, []dish.Dish{_dish})

	// By adding a menu of different type the number of available menus should now be 2

	_canteen.AddTodayMenu(_differentTypeMenu)

	availableMenus = _canteen.AvailableMenus()

	if lenad := len(availableMenus); lenad != 2 {
		t.Errorf("The length of availableMenus slice should now be 2 but got: %d", lenad)
	}

	// By adding a menu of the same type, an error should return and the number of available menus should still be 2

	_sameTypeMenu, _ := menu.New(0, []dish.Dish{_dish})

	err := _canteen.AddTodayMenu(_sameTypeMenu)

	if err == nil {
		t.Errorf("AddTodayMethod should have returned an error")
	}

	availableMenus = _canteen.AvailableMenus()

	if lenad := len(availableMenus); lenad != 2 {
		t.Errorf("The length of availableMenus slice should still be 2 but got: %d", lenad)
	}
}

func TestCanteensWithDifferentNameAreNotEqual(t *testing.T) {
	location := Location{}
	_canteenone, _ := New("Cantina do H", location)

	_canteentwo, _ := New("Cantina do F", location)

	if _canteenone.Equals(_canteentwo) {
		t.Errorf("_canteenone name is: %s and _canteentwo name is: %s, which are different, but were proved to be equal", _canteenone.Name, _canteentwo.Name)
	}
}

func TestCanteensWithEqualNameAreEqual(t *testing.T) {
	location := Location{}
	_canteenone, _ := New("Cantina do H", location)

	_canteentwo, _ := New("Cantina do H", location)

	if !_canteenone.Equals(_canteentwo) {
		t.Errorf("_canteenone name is: %s and _canteentwo name is: %s, which are equal, but were proved to not be equal", _canteenone.Name, _canteentwo.Name)
	}
}

func TestCanteensWithEqualNameButDifferentCaseAreEqual(t *testing.T) {
	location := Location{}
	_canteenone, _ := New("Cantina do H", location)

	_canteentwo, _ := New("cantina do h", location)

	if !_canteenone.Equals(_canteentwo) {
		t.Errorf("_canteenone name is: %s and _canteentwo name is: %s, which are equal, but were proved to not be equal", _canteenone.Name, _canteentwo.Name)
	}
}
