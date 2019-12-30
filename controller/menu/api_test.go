package menu

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/freitzzz/iped/model/dish"
	"github.com/freitzzz/iped/model/menu"

	"github.com/freitzzz/iped/model/canteen"
	"github.com/freitzzz/iped/model/school"

	"github.com/freitzzz/iped/controller/db"

	"github.com/labstack/echo"
)

var ech *echo.Echo

func init() {
	os.Setenv("IPEW_CONNECTION_STRING", ":memory:")

	ech = echo.New()

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_school, _ := school.New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	db.Db.Create(&_school)

	_canteen2, _ := canteen.New("Cantina do ESMAE")

	_dish, _ := dish.New(0, "Fried Noodles")

	_dishes := []dish.Dish{_dish}

	_menu, _ := menu.New(0, _dishes)

	_canteen2.AddTodayMenu(_menu)

	_canteens2 := []canteen.Canteen{_canteen2}

	_school2, _ := school.New("ESMAE", "Escola Superior Multimédia e Artes do Espetáculo", _canteens2)

	db.Db.Create(&_school2)

	_canteen3, _ := canteen.New("Cantina da ESS")

	_dish2, _ := dish.New(0, "Fried Noodles")

	_dishes2 := []dish.Dish{_dish2}

	_menu2, _ := menu.New(0, _dishes2)

	_canteen3.AddTodayMenu(_menu2)

	_canteens3 := []canteen.Canteen{_canteen3}

	_school3, _ := school.New("ESS", "Escola Superior de Saúde", _canteens3)

	db.Db.Create(&_school3)

}

func TestAvailableMenusReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/1/canteens/1/menus")

	// GET /menus
	AvailableMenus(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestAvailableMenusReturnsNotFoundIfSchoolDoesntExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("4", "1")
	ctx.Set("db", db.Db)

	// GET /menus
	AvailableMenus(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no school the resource identifier '4' so the response status code should be 404 but was: %d", rec.Code)
	}

}

func TestAvailableMenusReturnsNotFoundIfCanteenDoesntExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "2")
	ctx.Set("db", db.Db)

	// GET /menus
	AvailableMenus(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no canteen the resource identifier '2' so the response status code should be 404 but was: %d", rec.Code)
	}

}

func TestAvailableMenusReturnsNotFoundIfNoMenusAreAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "1")
	ctx.Set("db", db.Db)

	// GET /menus
	AvailableMenus(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("No menus are available so the response status code should be 404 but was: %d", rec.Code)
	}

}

func TestAvailableMenusReturnsOKIfMenusAreAvailable(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("2", "2")
	ctx.Set("db", db.Db)

	// GET /menus
	AvailableMenus(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("There is one menu available in the database so the response status code should be 200 but was: %d", rec.Code)
	}

}

func TestDetailedMenuInformationReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus/:id3")
	ctx.SetParamNames("id", "id2", "id3")
	ctx.SetParamValues("1", "1", "1")

	// GET /menus
	DetailedMenuInformation(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestDetailedMenuInformationReturnsNotFoundIfSchoolDoesntExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus/:id3")
	ctx.SetParamNames("id", "id2", "id3")
	ctx.SetParamValues("4", "2", "1")
	ctx.Set("db", db.Db)

	// GET /menus/:id
	DetailedMenuInformation(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no school with the resource id '4' so the response status code should be 404 but was: %d", rec.Code)
	}
}

func TestDetailedMenuInformationReturnsNotFoundIfCanteenDoesntExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus/:id3")
	ctx.SetParamNames("id", "id2", "id3")
	ctx.SetParamValues("2", "1", "1")
	ctx.Set("db", db.Db)

	// GET /menus/:id
	DetailedMenuInformation(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no canteen with the resource id '1' so the response status code should be 404 but was: %d", rec.Code)
	}
}

func TestDetailedMenuInformationReturnsNotFoundIfMenuDoesntExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus/:id3")
	ctx.SetParamNames("id", "id2", "id3")
	ctx.SetParamValues("3", "3", "1")
	ctx.Set("db", db.Db)

	// Menu that doens't exist in the collection

	// GET /menus/:id
	DetailedMenuInformation(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no menu with the resource id '2' so the response status code should be 404 but was: %d", rec.Code)
	}

	// Menu that doesn't exist in the database

	ctx.SetParamNames("id", "id2", "id3")
	ctx.SetParamValues("3", "3", "3")

	DetailedMenuInformation(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no menu with the resource id '2' so the response status code should be 404 but was: %d", rec.Code)
	}
}

func TestDetailedMenuInformationReturnsOKIfResourceWasFound(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus/:id3")
	ctx.SetParamNames("id", "id2", "id3")
	ctx.SetParamValues("2", "2", "1")
	ctx.Set("db", db.Db)

	// GET /menus/:id
	DetailedMenuInformation(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("There is a menu with the resource id %d so the response status code should be 200 but was: %d", 1, rec.Code)
	}
}

func TestCreateNewMenuReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "1")
	// POST /menus
	CreateNewMenu(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestCreateNewMenuReturnsBadRequestIfRequestBodyIsEmpty(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "1")
	ctx.Set("db", db.Db)

	// POST /menus
	CreateNewMenu(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("The request body is empty so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewMenuReturnsNotFoundIfSchoolDoesntExist(t *testing.T) {
	json := `{
    "type":"lunch",
    "dishes":[
        {
            "type":"meat",
            "description":"Arroz de pato"
        },
        {
            "type":"vegetarian",
            "description":"Tofu com tomate"
        }
    ]
}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("4", "1")
	ctx.Set("db", db.Db)

	// POST /menus
	CreateNewMenu(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no school with the resource identifier '4' so the response status code should be 404 but was : %d", rec.Code)
	}
}

func TestCreateNewMenuReturnsNotFoundIfCanteenDoesntExist(t *testing.T) {
	json := `{
    "type":"lunch",
    "dishes":[
        {
            "type":"meat",
            "description":"Arroz de pato"
        },
        {
            "type":"vegetarian",
            "description":"Tofu com tomate"
        }
    ]
}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "2")
	ctx.Set("db", db.Db)

	// POST /menus
	CreateNewMenu(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no canteen with the resource identifier '2' so the response status code should be 404 but was : %d", rec.Code)
	}
}

func TestCreateNewMenuReturnsBadRequestIfMenuFieldsAreInvalid(t *testing.T) {

	json := `{
    "type":"",
    "dishes":[
        {
            "type":"meat",
            "description":"Arroz de pato"
        },
        {
            "type":"vegetarian",
            "description":"Tofu com tomate"
        }
    ]
}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "1")
	ctx.Set("db", db.Db)

	// POST /menus
	CreateNewMenu(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("The menu type is invalid so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewMenuReturnsBadRequestIfMenuDishesAreInvalid(t *testing.T) {

	json := `{
    "type":"",
    "dishes":[
        {
            "type":"",
            "description":"Arroz de pato"
        },
        {
            "type":"vegetarian",
            "description":"Tofu com tomate"
        }
    ]
}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "1")
	ctx.Set("db", db.Db)

	// POST /menus
	CreateNewMenu(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("The menu type is invalid so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewMenuReturnsBadRequestIfMenuAlreadyExists(t *testing.T) {

	json := `{
    "type":"lunch",
    "dishes":[
        {
            "type":"meat",
            "description":"Arroz de pato"
        },
        {
            "type":"vegetarian",
            "description":"Tofu com tomate"
        }
    ]
}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("2", "2")
	ctx.Set("db", db.Db)

	// POST /menus
	CreateNewMenu(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Canteen with the resource identifier '2' already provides a 'lunch' type menu so the response status code should be 400 but was : %d", rec.Code)
	}
}

func TestCreateNewMenuReturnsCreatedIfAllMenuFieldsAreValid(t *testing.T) {
	json := `{
    "type":"dinner",
    "dishes":[
        {
            "type":"meat",
            "description":"Arroz de pato"
        },
        {
            "type":"vegetarian",
            "description":"Tofu com tomate"
        }
    ]
}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id/canteens/:id2/menus")
	ctx.SetParamNames("id", "id2")
	ctx.SetParamValues("1", "1")
	ctx.Set("db", db.Db)

	// POST /menus
	CreateNewMenu(ctx)

	if rec.Code != http.StatusCreated {
		t.Errorf("Canteen with the resource identifier '2' doesn't provide already any 'dinner' type menu so the response status code should be 201 but was : %d", rec.Code)
	}
}
