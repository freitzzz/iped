package canteen

import (
	"net/http"
	"strconv"

	"github.com/freitzzz/iped/model/school"

	"github.com/freitzzz/iped/model/canteen"
	customerrorview "github.com/freitzzz/iped/view/customerror"

	model "github.com/freitzzz/iped/model/canteen"
	view "github.com/freitzzz/iped/view/canteen"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// AvailableCanteens handles GET /canteens functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/canteens.md#available-canteens
func AvailableCanteens(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB) //schools/:id/canteens

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	_schoolID, _ := strconv.Atoi(c.Param("id"))

	_school := school.School{}

	err := db.Find(&_school, _schoolID).Related(&_school.CanteensSlice).Error

	// No need to check slice length as a school requires at least one canteen

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	_canteens := _school.Canteens()

	modelview := view.ToGetAvailableCanteensModelView(_canteens)

	return c.JSON(http.StatusOK, modelview)

}

// DetailedCanteenInformation handles GET /canteens/:id functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/canteens.md#detailed-canteen-information
func DetailedCanteenInformation(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	_schoolID, _ := strconv.Atoi(c.Param("id"))

	_canteenID, _ := strconv.Atoi(c.Param("id2"))

	_canteen := canteen.Canteen{}

	_canteen.SchoolID = uint(_schoolID)

	_canteen.ID = uint(_canteenID)

	err := db.Where(&_canteen).First(&_canteen).Error

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	modelview := view.ToGetDetailedCanteenInformationModelView(_canteen)

	return c.JSON(http.StatusOK, modelview)

}

// CreateNewCanteen handles POST /canteens functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/canteens.md#create-a-new-canteen
func CreateNewCanteen(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	var modelview view.CreateNewCanteenModelView

	c.Bind(&modelview)

	location := model.Location{}

	location.Latitude = modelview.Location.Latitude

	location.Longitude = modelview.Location.Longitude

	canteen, serr := model.New(modelview.Name, location)

	if serr != nil {

		modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(*serr)

		return c.JSON(http.StatusBadRequest, modelview)
	}

	_schoolID, _ := strconv.Atoi(c.Param("id"))

	_school := school.School{}

	ferr := db.Find(&_school, _schoolID).Related(&_school.CanteensSlice).Error

	if ferr != nil {
		return c.NoContent(http.StatusNotFound)
	}

	err := _school.AddCanteen(canteen)

	if err != nil {

		return c.JSON(http.StatusBadRequest, customerrorview.UsingFieldErrorToErrorMessageModelView(*err))
	}

	// Creates canteen
	db.Save(&_school)

	_canteen := model.Canteen{}

	_canteen.SchoolID = _school.ID

	_canteen.Name = canteen.Name

	db.Where(&_canteen).First(&_canteen)

	modelviewres := view.ToGetDetailedCanteenInformationModelView(_canteen)

	return c.JSON(http.StatusCreated, modelviewres)

}
