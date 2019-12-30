// This gofile is the entrypoint for iped
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/freitzzz/iped/controller/canteen"
	"github.com/freitzzz/iped/controller/db"
	"github.com/freitzzz/iped/controller/dish"
	"github.com/freitzzz/iped/controller/menu"
	"github.com/freitzzz/iped/controller/middleware"
	"github.com/freitzzz/iped/controller/school"
	"github.com/labstack/echo"
)

func main() {

	echo.NotFoundHandler = middleware.NotFoundHandler()

	ech := echo.New()

	ech.Use(middleware.DbAccessMiddleware())

	ech.Use(middleware.ResourceIdentifierValidationMiddleware())

	// schools collection functionalities

	ech.GET("/schools", school.AvailableSchools)

	ech.GET("/schools/:id", school.DetailedSchoolInformation)

	ech.POST("/schools", school.CreateNewSchool)

	// canteens collection functionalities

	ech.GET("/schools/:id/canteens", canteen.AvailableCanteens)

	ech.GET("/schools/:id/canteens/:id2", canteen.DetailedCanteenInformation)

	ech.POST("/schools/:id/canteens", canteen.CreateNewCanteen)

	// menus collection functionalities

	ech.GET("/schools/:id/canteens/:id2/menus", menu.AvailableMenus)

	ech.GET("/schools/:id/canteens/:id2/menus/:id3", menu.DetailedMenuInformation)

	ech.POST("/schools/:id/canteens/:id2/menus", menu.CreateNewMenu)

	// dishes collection functionalities

	ech.GET("/schools/:id/canteens/:id2/menus/:id3/dishes", dish.AvailableDishes)

	ech.GET("/schools/:id/canteens/:id2/menus/:id3/dishes/:id4", dish.DetailedDishInformation)

	port, perr := strconv.Atoi(os.Getenv("PORT"))

	if perr != nil {
		panic(fmt.Sprint("Server couldn't be open as the specified port is not valid"))
	}

	ech.Start(fmt.Sprintf(":%d", port))

	defer db.Db.Close()
}
