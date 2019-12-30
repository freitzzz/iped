package school

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/freitzzz/iped/model/canteen"
	"github.com/freitzzz/iped/model/school"
	model "github.com/freitzzz/iped/model/school"

	"github.com/freitzzz/iped/controller/db"

	"github.com/labstack/echo"
)

var ech *echo.Echo

func init() {
	os.Setenv("IPEW_CONNECTION_STRING", ":memory:")

	ech = echo.New()

}

func TestAvailableSchoolsReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")

	// GET /schools
	AvailableSchools(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestAvailableSchoolsReturnsNotFoundIfNoSchoolsExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")
	ctx.Set("db", db.Db)

	// GET /schools
	AvailableSchools(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("As there are no schools available in the database the response status code should be 404 but was: %d", rec.Code)
	}

}

func TestAvailableSchoolsReturnsOKIfSchoolsExist(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_school, _ := model.New("ISEP", "Instituto Superior de Engenharia do Porto", []canteen.Canteen{_canteen})

	// Inserts a school in the database

	db.Db.Create(&_school)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")
	ctx.Set("db", db.Db)

	// GET /schools/:id
	AvailableSchools(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("There is one school available in the database so the response status code should be 200 but was: %d", rec.Code)
	}

}

func TestDetailedSchoolInformationReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	// GET /schools
	DetailedSchoolInformation(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestDetailedSchoolInformationReturnsNotFoundIfResourceWasNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("10")
	ctx.Set("db", db.Db)

	// GET /schools/:id
	DetailedSchoolInformation(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no school with the resource id '10' so the response status code should be 404 but was: %d", rec.Code)
	}
}

func TestDetailedSchoolInformationReturnsOKIfResourceWasFound(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do ESMAE")

	_school, _ := school.New("ESMAE", "Escola Superior de Música e Artes do Espetáculo", []canteen.Canteen{_canteen})

	db.Db.Create(&_school)

	id := _school.ID

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(id), 10))
	ctx.Set("db", db.Db)

	// GET /schools/:id
	DetailedSchoolInformation(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("There is a school with the resource id %d so the response status code should be 200 but was: %d", id, rec.Code)
	}
}

func TestCreateNewSchoolReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")

	// POST /schools
	CreateNewSchool(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestCreateNewSchoolReturnsBadRequestIfRequestBodyIsEmpty(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")
	ctx.Set("db", db.Db)

	// POST /schools
	CreateNewSchool(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("The request body is empty so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewSchoolReturnsBadRequestIfCanteenFieldsAreInvalid(t *testing.T) {

	json := `{
		"acronym":"ISEP", 
		"name":"Instituto Superior de Engenharia do Porto", 
		"canteens":[
			{
				"name":""
			}	
		]
	}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")
	ctx.Set("db", db.Db)

	// POST /schools
	CreateNewSchool(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("The request body an invalid canteen so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewSchoolReturnsBadRequestIfSchoolFieldsAreInvalid(t *testing.T) {
	json := `{
		"acronym":"ISEP", 
		"name":"Instituto Superior de Engenharia do Porto"`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")
	ctx.Set("db", db.Db)

	// POST /schools
	CreateNewSchool(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("The request body provides no canteens so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewSchoolReturnsBadRequestIfSchoolWithSameAcronymAlreadyExists(t *testing.T) {

	_canteen, _ := canteen.New("Cantina da ESS")

	_canteens := []canteen.Canteen{_canteen}

	_school, _ := model.New("ESS", "Escola Superior de Saúde", _canteens)

	db.Db.Create(&_school)

	json := `{
		"acronym":"ESS", 
		"name":"Escola Superior de Saúde", 
		"canteens":[
			{
				"name":"Cantina da ESS"
			}	
		]
	}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")
	ctx.Set("db", db.Db)

	// POST /schools
	CreateNewSchool(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("The school being created already exists so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewSchoolReturnsCreatedIfAllSchoolFieldsAreValid(t *testing.T) {
	json := `{
		"acronym":"IPP", 
		"name":"Instituto Politécnico do Porto", 
		"canteens":[
			{
				"name":"Cantina do IPP"
			}	
		]
	}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")
	ctx.Set("db", db.Db)

	// POST /schools
	CreateNewSchool(ctx)

	if rec.Code != http.StatusCreated {
		t.Errorf("All school fields are valid so the response status should be created: %d", rec.Code)
	}
}
