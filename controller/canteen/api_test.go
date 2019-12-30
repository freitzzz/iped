package canteen

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/freitzzz/iped/model/school"

	"github.com/freitzzz/iped/model/canteen"
	model "github.com/freitzzz/iped/model/canteen"

	"github.com/freitzzz/iped/controller/db"

	"github.com/labstack/echo"
)

var ech *echo.Echo

func init() {
	os.Setenv("IPEW_CONNECTION_STRING", ":memory:")

	ech = echo.New()

	_canteen, _ := model.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_school, _ := school.New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	db.Db.Create(&_school)

}

func TestAvailableCanteensReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/1/canteens")

	// GET /canteens
	AvailableCanteens(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestAvailableCanteensReturnsNotFoundIfSchoolDoesntExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens")
	ctx.SetParamNames("id")
	ctx.SetParamValues("2")
	ctx.Set("db", db.Db)

	// GET /canteens
	AvailableCanteens(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no school the resource identifier '2' so the response status code should be 404 but was: %d", rec.Code)
	}

}

func TestAvailableCanteensReturnsOKIfSchoolExists(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	ctx.Set("db", db.Db)

	// GET /canteens/:id
	AvailableCanteens(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("There is one canteen available in the database so the response status code should be 200 but was: %d", rec.Code)
	}

}

func TestDetailedCanteenInformationReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "1")

	// GET /canteens
	DetailedCanteenInformation(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestDetailedCanteenInformationReturnsNotFoundIfSchoolDoesntExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("2", "1")
	ctx.Set("db", db.Db)

	// GET /canteens/:id
	DetailedCanteenInformation(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no school with the resource id '2' so the response status code should be 404 but was: %d", rec.Code)
	}
}

func TestDetailedCanteenInformationReturnsNotFoundIfCanteenDoesntExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "2")
	ctx.Set("db", db.Db)

	// GET /canteens/:id
	DetailedCanteenInformation(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no canteen with the resource id '2' so the response status code should be 404 but was: %d", rec.Code)
	}
}

func TestDetailedCanteenInformationReturnsOKIfResourceWasFound(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "1")
	ctx.Set("db", db.Db)

	// GET /canteens/:id
	DetailedCanteenInformation(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("There is a canteen with the resource id %d so the response status code should be 200 but was: %d", 1, rec.Code)
	}
}

func TestCreateNewCanteenReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	// POST /canteens
	CreateNewCanteen(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestCreateNewCanteenReturnsBadRequestIfRequestBodyIsEmpty(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	ctx.Set("db", db.Db)

	// POST /canteens
	CreateNewCanteen(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("The request body is empty so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewCanteenReturnsNotFoundIfSchoolDoesntExist(t *testing.T) {
	json := `{
		"name":"Cantina do J"
	}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens")
	ctx.SetParamNames("id")
	ctx.SetParamValues("2")
	ctx.Set("db", db.Db)

	// POST /canteens
	CreateNewCanteen(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no school with the resource identifier '2' so the response status code should be 404 but was : %d", rec.Code)
	}
}

func TestCreateNewCanteenReturnsBadRequestIfCanteenFieldsAreInvalid(t *testing.T) {

	json := `{
		"name":""
	}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	ctx.Set("db", db.Db)

	// POST /canteens
	CreateNewCanteen(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("The canteen name is invalid so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewCanteenReturnsBadRequestIfCanteenAlreadyExists(t *testing.T) {

	json := `{
		"name":"Cantina do H"
	}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	ctx.Set("db", db.Db)

	// POST /canteens
	CreateNewCanteen(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("School with the resource identifier '1' already provides a canteen with the name 'Cantina do H' so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewCanteenReturnsCreatedIfAllCanteenFieldsAreValid(t *testing.T) {
	json := `{
		"name":"Cantina do J"
	}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	ctx.Set("db", db.Db)

	// POST /canteens
	CreateNewCanteen(ctx)

	if rec.Code != http.StatusCreated {
		t.Errorf("School with the resource identifier '1' doesn't provide any canteen with the name 'Cantinado J' so the response status code should be 201 but was : %d", rec.Code)
	}
}
