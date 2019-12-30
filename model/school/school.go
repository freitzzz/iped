package school

import (
	"strings"
	"unicode"

	"github.com/jinzhu/gorm"

	"github.com/freitzzz/iped/model/canteen"
	"github.com/freitzzz/iped/model/customerror"
)

// School is a model that offers canteens
// A school has a unique acronym, a descriptive name and needs to offer at least one canteen
// A UML overview of this model can be found at https://github.com/freitzzz/iped-documentation/wiki/Architecture#models-structure
type School struct {
	gorm.Model
	Acronym       string `gorm:"UNIQUE"`
	Name          string
	CanteensSlice []canteen.Canteen `gorm:"column:canteens"`
}

// New initializes a school model using its acronym, name and canteens
// A FieldError is returned if the canteen acronym is empty or has spaces between letters,
// name is empty, no canteens were provided or if it was found a duplicated canteen
func New(Acronym string, Name string, Canteens []canteen.Canteen) (School, *customerror.FieldError) {

	school := School{gorm.Model{}, Acronym, Name, Canteens}

	err := grantSchoolAcronymIsNotEmpty(Acronym)

	if err != nil {
		return school, err
	}

	err = grantSchoolAcronymDoesNotHaveSpacesBetweenLetters(Acronym)

	if err != nil {
		return school, err
	}

	err = grantSchoolNameIsNotEmpty(Name)

	if err != nil {
		return school, err
	}

	err = grantAtLeastOneCanteenIsProvided(Canteens)

	if err != nil {
		return school, err
	}

	err = grantNoDuplicatedCanteensExist(Canteens)

	if err != nil {
		return school, err
	}

	return school, err
}

// Canteens returns the available canteens provided by a school as a slice
// The returned slice has different reference of the one in School struct
// In order to prevent modifications
func (school School) Canteens() []canteen.Canteen {

	availableCanteens := make([]canteen.Canteen, len(school.CanteensSlice))

	copy(availableCanteens, school.CanteensSlice)

	return availableCanteens

}

// AddCanteen allows the addition of a new canteen to the already provided by the school
// An error is returned if the canteen being added already exists
func (school *School) AddCanteen(canteen canteen.Canteen) *customerror.FieldError {
	var err *customerror.FieldError

	schoolCanteens := school.Canteens()

	schoolCanteens = append(schoolCanteens, canteen)

	err = grantNoDuplicatedCanteensExist(schoolCanteens)

	if err == nil {
		school.CanteensSlice = schoolCanteens
	}

	return err
}

// This function grants that a school acronym is not empty, and if empty returns a FieldError
func grantSchoolAcronymIsNotEmpty(acronym string) *customerror.FieldError {

	var err *customerror.FieldError

	if len(strings.TrimSpace(acronym)) == 0 {
		err = &customerror.FieldError{Field: "acronym", Model: "school", Explanation: "school acronym cannot be an empty string"}
	}

	return err

}

// This function grants that a school acronym does not have spaces between letters, or else returns a FieldError
func grantSchoolAcronymDoesNotHaveSpacesBetweenLetters(acronym string) *customerror.FieldError {

	var err *customerror.FieldError

	acronymLength := len(acronym)

	acronymLength--

	for acronymLength >= 0 {
		if unicode.IsSpace(rune(acronym[acronymLength])) {
			acronymLength = -2
		} else {
			acronymLength--
		}
	}

	if acronymLength == -2 {
		err = &customerror.FieldError{Field: "acronym", Model: "school", Explanation: "school acronym cannot have spaces between letters"}
	}

	return err

}

// This function grants that a school name is not empty, and if empty returns a FieldError
func grantSchoolNameIsNotEmpty(name string) *customerror.FieldError {

	var err *customerror.FieldError

	if len(strings.TrimSpace(name)) == 0 {
		err = &customerror.FieldError{Field: "name", Model: "school", Explanation: "school name cannot be an empty string"}
	}

	return err

}

// This function grants that at least one canteen is provided in given canteen slice
// If the given canteen slice is nil or empty a FieldError is returned
func grantAtLeastOneCanteenIsProvided(canteens []canteen.Canteen) *customerror.FieldError {

	var err *customerror.FieldError

	if canteens == nil || len(canteens) == 0 {
		err = &customerror.FieldError{Field: "canteens", Model: "school", Explanation: "school requires at least one canteen"}
	}

	return err

}

// This function grants that all canteen given in a slice are unique
// If a canteen proves equality to any other canteen in the slice, a FieldError is returned
func grantNoDuplicatedCanteensExist(canteens []canteen.Canteen) *customerror.FieldError {

	var err *customerror.FieldError

	unique := true
	canteensLength := len(canteens)
	i := 0

	for i < canteensLength {
		j := i + 1
		for j < canteensLength {
			unique = !canteens[i].Equals(canteens[j])
			j++
			if !unique {
				i = canteensLength
				j = i
			}
		}
		i++
	}

	if !unique {
		err = &customerror.FieldError{Field: "canteens", Model: "school", Explanation: "school cannot have canteens with the same name"}
	}

	return err
}
