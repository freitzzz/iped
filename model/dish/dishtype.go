package dish

import "strings"

// DishType is a enum representation of a dish type
// A UML overview of this enum can be found at https://github.com/freitzzz/iped-documentation/wiki/Architecture#models-structure
type DishType int

// Possible dish types are the following
const (
	Meat DishType = iota
	Fish
	Vegetarian
	Diet
)

// String is the implementation of stringer for DishType enum
// If the DishType being converted is not valid the string "nil" will be returned
func (Type DishType) String() string {

	switch Type {
	case Meat:
		return "meat"
	case Fish:
		return "fish"
	case Vegetarian:
		return "vegetarian"
	case Diet:
		return "diet"
	default:
		return "nil"
	}
}

// Parse converts a string in a DishType
// If string is not recognized as a dish type, -1 is returned
func Parse(dishAsString string) DishType {
	switch strings.ToLower(dishAsString) {
	case "meat":
		return DishType(0)
	case "fish":
		return DishType(1)
	case "vegetarian":
		return DishType(2)
	case "diet":
		return DishType(3)
	default:
		return -1
	}
}

// Validate check if a given integer is a valid dish type
func Validate(dishtype int) bool {
	return dishtype >= 0 && dishtype <= 3
}
