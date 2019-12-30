package school

import (
	"net/http"
	"strconv"

	customerrormodel "github.com/freitzzz/iped/model/customerror"
	customerrorview "github.com/freitzzz/iped/view/customerror"

	"github.com/freitzzz/iped/model/canteen"

	model "github.com/freitzzz/iped/model/school"
	view "github.com/freitzzz/iped/view/school"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// AvailableSchools handles GET /schools functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/schools.md#available-schools
func AvailableSchools(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	schools := []model.School{}

	// Finds all available schools

	err := db.Find(&schools).Error

	if err != nil || len(schools) == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	modelview := view.ToGetAvailableSchoolsModelView(schools)

	return c.JSON(http.StatusOK, modelview)

}

// DetailedSchoolInformation handles GET /schools/:id functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/schools.md#detailed-school-information
func DetailedSchoolInformation(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	var school model.School

	// Finds school by ID

	err := db.Find(&school, id).Error

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	// Find school canteens

	db.Model(&school).Related(&school.CanteensSlice)

	modelview := view.ToGetDetailedSchoolInformationModelView(school)

	return c.JSON(http.StatusOK, modelview)

}

// CreateNewSchool handles POST /schools functionality
// See more info at: https://github.com/freitzzz/iped-documentation/blob/master/documentation/rest_api/schools.md#create-a-new-school
func CreateNewSchool(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	var modelview view.CreateNewSchoolModelView

	c.Bind(&modelview)

	canteens := make([]canteen.Canteen, len(modelview.Canteens))

	for index := range modelview.Canteens {

		location := canteen.Location{}

		location.Latitude = modelview.Canteens[index].Location.Latitude

		location.Longitude = modelview.Canteens[index].Location.Longitude

		canteen, cerr := canteen.New(modelview.Canteens[index].Name, location)
		if cerr != nil {

			modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(*cerr)

			return c.JSON(http.StatusBadRequest, modelview)
		}
		canteens[index] = canteen
	}

	school, serr := model.New(modelview.Acronym, modelview.Name, canteens)

	if serr != nil {

		modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(*serr)

		return c.JSON(http.StatusBadRequest, modelview)
	}

	var existingSchool model.School

	// Finds if school with same acronym already exists

	err := db.Where(map[string]interface{}{"acronym": modelview.Acronym}).First(&existingSchool).Error

	if err == nil {

		cerr := customerrormodel.FieldError{Field: "acronym", Model: "school", Explanation: "a school with the same acronym already exists"}

		modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(cerr)

		return c.JSON(http.StatusBadRequest, modelview)
	}

	// Creates school
	db.Create(&school)

	modelviewres := view.ToGetDetailedSchoolInformationModelView(school)

	return c.JSON(http.StatusCreated, modelviewres)

}
