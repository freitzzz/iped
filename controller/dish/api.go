package dish

import (
	"net/http"
	"strconv"

	"github.com/freitzzz/iped/model/menu"

	"github.com/freitzzz/iped/model/canteen"

	model "github.com/freitzzz/iped/model/dish"
	view "github.com/freitzzz/iped/view/dish"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// AvailableDishes handles GET /dishes functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/dishes.md#available-dishes
func AvailableDishes(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	_schoolID, _ := strconv.Atoi(c.Param("id"))

	_canteenID, _ := strconv.Atoi(c.Param("id2"))

	_menuID, _ := strconv.Atoi(c.Param("id3"))

	_canteen := canteen.Canteen{}

	_canteen.SchoolID = uint(_schoolID)

	_canteen.ID = uint(_canteenID)

	_menu := menu.Menu{}

	err := db.Preload("DishesSlice").Find(&_menu, _menuID).Error

	// First check if menu with given id exists, and if not return Not Found

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	_menuEntry := canteen.MenuEntry{}

	err = db.Where(&_canteen).First(&_canteen).Error

	// Now lets grant that the menu belongs to the canteen, and if not return Not Found

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	_menuEntry.CanteenID = _canteen.ID

	_menuEntry.ID = _menu.MenuEntryID

	err = db.Where(&_menuEntry).First(&_menuEntry).Error

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	_dishes := _menu.Dishes()

	if len(_dishes) == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	modelview := view.ToGetAvailableDishesModelView(_dishes)

	return c.JSON(http.StatusOK, modelview)

}

// DetailedDishInformation handles GET /dishes/:id functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/dishes.md#detailed-dish-information
func DetailedDishInformation(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	_schoolID, _ := strconv.Atoi(c.Param("id"))

	_canteenID, _ := strconv.Atoi(c.Param("id2"))

	_menuID, _ := strconv.Atoi(c.Param("id3"))

	_dishID, _ := strconv.Atoi(c.Param("id4"))

	_canteen := canteen.Canteen{}

	_canteen.SchoolID = uint(_schoolID)

	_canteen.ID = uint(_canteenID)

	_menu := menu.Menu{}

	_dish := model.Dish{}

	err := db.Find(&_dish, _dishID).Error

	// First check if dish with given id exists, and if not return Not Found

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	err = db.Find(&_menu, _menuID).Error

	// Check if menu with given id exists, and if not return Not Found

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	_menuEntry := canteen.MenuEntry{}

	err = db.Where(&_canteen).First(&_canteen).Error

	// Now lets grant that the menu belongs to the canteen, and if not return Not Found

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	_menuEntry.CanteenID = _canteen.ID

	_menuEntry.ID = _menu.MenuEntryID

	err = db.Where(&_menuEntry).First(&_menuEntry).Error

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	modelview := view.ToGetDetailedDishInformationModelView(_dish)

	return c.JSON(http.StatusOK, modelview)

}
