package menu

import "strings"

// MenuType is a enum representation of a menu type
// A UML overview of this enum can be found at https://github.com/freitzzz/iped-documentation/wiki/Architecture#models-structure
type MenuType int

// A menu is either available at lunch or dinner
const (
	Lunch MenuType = iota
	Dinner
)

// String is the implementation of stringer for MenuType enum
// If the MenuType being converted is not valid the string "nil" will be returned
func (Type MenuType) String() string {

	switch Type {
	case Lunch:
		return "lunch"
	case Dinner:
		return "dinner"
	default:
		return "nil"
	}
}

// Parse converts a string in a MenuType
// If string is not recognized as a menu type, -1 is returned
func Parse(menuAsString string) MenuType {
	switch strings.ToLower(menuAsString) {
	case "lunch":
		return MenuType(0)
	case "dinner":
		return MenuType(1)
	default:
		return -1
	}
}

// Validate check if a given integer is a valid menu type
func Validate(menutype int) bool {
	return menutype == 0 || menutype == 1
}
