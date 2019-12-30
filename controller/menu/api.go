package menu

import (
	"net/http"
	"strconv"

	"github.com/freitzzz/iped/model/dish"

	"github.com/freitzzz/iped/model/canteen"

	customerrorview "github.com/freitzzz/iped/view/customerror"

	model "github.com/freitzzz/iped/model/menu"
	view "github.com/freitzzz/iped/view/menu"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// AvailableMenus handles GET /menus functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/menus.md#available-menus
func AvailableMenus(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB) //schools/:id/menus

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	_schoolID, _ := strconv.Atoi(c.Param("id"))

	_canteenID, _ := strconv.Atoi(c.Param("id2"))

	_canteen := canteen.Canteen{}

	_canteen.SchoolID = uint(_schoolID)

	_canteen.ID = uint(_canteenID)

	ferr := db.Where(&_canteen, _canteenID).First(&_canteen).Error

	if ferr != nil {
		return c.NoContent(http.StatusNotFound)
	}

	db.Model(&_canteen).Preload("Menus").Related(&_canteen.MenusMap)

	_menus := _canteen.AvailableMenus()

	if len(_menus) == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	modelview := view.ToGetAvailableMenusModelView(_menus)

	return c.JSON(http.StatusOK, modelview)

}

// DetailedMenuInformation handles GET /menus/:id functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/menus.md#detailed-menu-information
func DetailedMenuInformation(c echo.Context) error {

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

	_menu := model.Menu{}

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

	modelview := view.ToGetDetailedMenuInformationModelView(_menu)

	return c.JSON(http.StatusOK, modelview)

}

// CreateNewMenu handles POST /menus functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/menus.md#create-a-new-menu
func CreateNewMenu(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	var modelview view.CreateNewMenuModelView

	c.Bind(&modelview)

	dishes := make([]dish.Dish, len(modelview.Dishes))

	for index, dishview := range modelview.Dishes {
		_dishtype := int(dish.Parse(dishview.Type))
		_dish, cerr := dish.New(_dishtype, dishview.Description)
		if cerr != nil {

			modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(*cerr)

			return c.JSON(http.StatusBadRequest, modelview)
		}
		dishes[index] = _dish
	}

	menutype := int(model.Parse(modelview.Type))
	menu, serr := model.New(menutype, dishes)

	if serr != nil {

		modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(*serr)

		return c.JSON(http.StatusBadRequest, modelview)
	}

	_schoolID, _ := strconv.Atoi(c.Param("id"))

	_canteenID, _ := strconv.Atoi(c.Param("id2"))

	_canteen := canteen.Canteen{}

	_canteen.SchoolID = uint(_schoolID)

	_canteen.ID = uint(_canteenID)

	ferr := db.Where(&_canteen, _canteenID).First(&_canteen).Error

	if ferr != nil {
		return c.NoContent(http.StatusNotFound)
	}

	db.Model(&_canteen).Preload("Menus").Related(&_canteen.MenusMap)

	err := _canteen.AddTodayMenu(menu)

	if err != nil {

		return c.JSON(http.StatusBadRequest, customerrorview.UsingFieldErrorToErrorMessageModelView(*err))
	}

	// Creates menu
	db.Save(&_canteen)

	_menu := model.Menu{}

	db.Where(&_canteen.MenusMap).Last(&_menu)

	db.Model(&_menu).Related(&_menu.DishesSlice)

	modelviewres := view.ToGetDetailedMenuInformationModelView(_menu)

	return c.JSON(http.StatusCreated, modelviewres)

}
