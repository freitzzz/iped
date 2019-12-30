package canteen

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/freitzzz/iped/model/customerror"
	"github.com/freitzzz/iped/model/menu"
)

// Canteen is a model that has the responsibility to inform the user which menus are available at the time
// A canteen has a unique name and is offered by a school
// A UML overview of this model can be found at https://github.com/freitzzz/iped-documentation/wiki/Architecture#models-structure
type Canteen struct {
	gorm.Model
	SchoolID uint `gorm:"type:int REFERENCES schools(id) ON UPDATE CASCADE ON DELETE CASCADE"`
	Name     string
	MenusMap []MenuEntry
	Location Location
}

// MenuEntry struct is an entry to canteen menus map
// Its declared as way to create a custom map with a list as Gorm does not permit the mapping of map types
type MenuEntry struct {
	gorm.Model
	CanteenID uint `gorm:"type:int REFERENCES canteens(id) ON UPDATE CASCADE ON DELETE CASCADE"`
	Time      time.Time
	Menus     []menu.Menu
}

// Location struct identifies the geographical location of a canteen
type Location struct {
	Latitude  float32
	Longitude float32
}

// New initializes a Canteen model using its name
// A FieldError is returned if the canteen name is empty
func New(Name string, Location Location) (Canteen, *customerror.FieldError) {

	canteen := Canteen{gorm.Model{}, 0, Name, []MenuEntry{}, Location}

	err := grantCanteenNameIsNotEmpty(Name)

	if err != nil {
		return canteen, err
	}

	return canteen, err
}

// AddTodayMenu allows the addition of a menu to today available menus
// A FieldError is returned if it was found to exist a menu that has the same type
// as the existing available menus
func (canteen *Canteen) AddTodayMenu(Menu menu.Menu) *customerror.FieldError {

	var err *customerror.FieldError

	entry := canteen.areThereMenusForToday()

	availableMenus := []menu.Menu{}

	if entry != -1 {
		availableMenus = canteen.MenusMap[entry].Menus
	}

	if lena := len(availableMenus); lena != 0 {
		index := 0
		for index < lena {
			if availableMenus[index].Type == Menu.Type {
				index = lena
				err = &customerror.FieldError{Field: "menus", Model: "canteen", Explanation: "canteen does not allow providing menus of the same type"}
			} else {
				index++
			}
		}
	}

	if err == nil {
		availableMenus = append(availableMenus, Menu)

		if entry == -1 {
			menuEntry := MenuEntry{Menus: availableMenus, Time: todayDateTime()}

			canteen.MenusMap = append(canteen.MenusMap, menuEntry)

		} else {
			canteen.MenusMap[entry].Menus = availableMenus
		}
	}

	return err

}

// AvailableMenus returns the menus which the canteen is providing at the time being asked
// as a slice
// If no menus are available an empty slice is returned
// The returned slice has different reference of the one in Canteen struct
// In order to prevent modifications
func (canteen Canteen) AvailableMenus() []menu.Menu {

	var availableMenus []menu.Menu

	exists := canteen.areThereMenusForToday()

	if exists == -1 {
		availableMenus = []menu.Menu{}
	} else {
		availableMenus = canteen.MenusMap[exists].Menus
		availableMenusCopy := make([]menu.Menu, len(availableMenus))

		copy(availableMenusCopy, availableMenus)
	}

	return availableMenus

}

// Equals compares equality between two canteens
// A canteen proves true equality to other canteen if both canteen names are equal
// Canteen names are case sensitive
func (canteen Canteen) Equals(comparingCanteen Canteen) bool {
	return strings.ToUpper(canteen.Name) == strings.ToUpper(comparingCanteen.Name)
}

// Checks availability of today's menus, and if available returns the index that points to menu entry on menu map
// If there are no available menus, -1 is returned
func (canteen Canteen) areThereMenusForToday() int {

	todaydate := todayDateTime()

	for index, entry := range canteen.MenusMap {
		if entry.Time.Equal(todaydate) {
			return index
		}
	}

	return -1
}

// Returns today date as a [time.Time] struct
// The struct returned is formatted to be DD-MM-YYYY 00:00:00
func todayDateTime() time.Time {
	datetime := time.Now()

	datetime = time.Date(datetime.Year(), datetime.Month(), datetime.Day(), int(0), int(0), int(0), int(0), datetime.Location())

	return datetime
}

// This function grants that a canteen name is not empty, and if empty returns an FieldError
func grantCanteenNameIsNotEmpty(name string) *customerror.FieldError {

	var err *customerror.FieldError

	if len(strings.TrimSpace(name)) == 0 {
		err = &customerror.FieldError{Field: "name", Model: "canteen", Explanation: "canteen name cannot be an empty string"}
	}

	return err

}
