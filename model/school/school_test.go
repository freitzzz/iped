package school

import (
	"testing"

	"github.com/freitzzz/iped/model/canteen"
)

func TestEmptySchoolAcronymReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.Field != "acronym" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field acronym")
	}
}

func TestSchoolAcronymWithOnlySpacesReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New(" ", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.Field != "acronym" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field acronym")
	}
}

func TestSchoolAcronymWithSpacesBetweenLettersReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("IS EP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.Field != "acronym" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field acronym")
	}
}

func TestEmptySchoolNameReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", "", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.Field != "name" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field name")
	}
}

func TestSchoolNameWithOnlySpacesReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", " ", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.Field != "name" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field name")
	}
}

func TestIfNoSchoolCanteensAreProvidedAnErrorIsReturned(t *testing.T) {

	_canteens := []canteen.Canteen{}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.Field != "canteens" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field canteens")
	}
}

func TestIfDuplicatedSchoolCanteensAreProvidedAnErrorIsReturned(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")
	_canteens := []canteen.Canteen{_canteen, _canteen}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.Field != "canteens" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field canteens")
	}
}

func TestNotEmptyAndNoSpacesBetweenLettersSchoolAcronymDoesNotReturnError(t *testing.T) {
	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err != nil {
		t.Errorf("School initilization should have been successful but got %s", err)
	}
}

func TestNotEmptySchoolNameDoesNotReturnError(t *testing.T) {
	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err != nil {
		t.Errorf("School initilization should have been successful but got %s", err)
	}
}

func TestNotEmptyOrDuplicatedSchoolCanteensDoesNotReturnError(t *testing.T) {
	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err != nil {
		t.Errorf("School initilization should have been successful but got %s", err)
	}
}

func TestCanteensMethodReturnsSliceWithDifferentReference(t *testing.T) {
	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_school, _ := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	availableCanteens := _school.Canteens()

	// This verification is to grant that the returned available canteens slice length is 0

	if lenab := len(availableCanteens); lenab != 1 {
		t.Errorf("The length of availableCanteens slice should be 1 but got: %d", lenab)
	}

	// If the length of the slice is the same as the capacity the slice was successfuly copied from the original slice

	if capb := cap(availableCanteens); capb != 1 {
		t.Errorf("The capacitiy of availableCanteens should be the same as its length (1) but got %d", capb)
	}

	// If we add a new canteen to the the returned slice,
	// it should not modify the slice pointed on the school struct

	_differentCanteen, _ := canteen.New("Cantina do F")

	availableCanteens = append(availableCanteens, _differentCanteen)

	if lenam := len(availableCanteens); lenam != 2 {
		t.Errorf("The length of availableCanteens slice should now be 2 but got: %d", lenam)
	}

	availableCanteensAfterModification := _school.Canteens()

	if lenaam := len(availableCanteensAfterModification); lenaam != 1 {
		t.Errorf("The length of availableCanteensAfterModification slice should be 1 but got: %d", lenaam)
	}

	if capb := cap(availableCanteensAfterModification); capb != 1 {
		t.Errorf("The capacitiy of availableCanteensAfterModification should be the same as its length (1) but got %d", capb)
	}

}

func TestAddCanteenReturnsErrorOnDuplicatedCanteen(t *testing.T) {
	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_school, _ := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	availableCanteens := _school.Canteens()

	if lena := len(availableCanteens); lena != 1 {
		t.Errorf("It should only exist 1 school canteen but got: %d", lena)
	}

	_duplicatedCanteen, _ := canteen.New("cantina do h")

	err := _school.AddCanteen(_duplicatedCanteen)

	if err == nil {
		t.Errorf("The canteen with the name: %s was successfuly added but the school already provides a canteen with the name: %s", _duplicatedCanteen.Name, _canteen.Name)
	}

	availableCanteens = _school.Canteens()

	if lena := len(availableCanteens); lena != 1 {
		t.Errorf("The number of existing schools should still be 1 but got: %d", lena)
	}
}

func TestAddCanteenDoesNotReturnErrorOnCanteenWithDifferentName(t *testing.T) {
	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_school, _ := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	availableCanteens := _school.Canteens()

	if lena := len(availableCanteens); lena != 1 {
		t.Errorf("It should only exist 1 school canteen but got: %d", lena)
	}

	_differentNameCanteen, _ := canteen.New("cantina do f")

	err := _school.AddCanteen(_differentNameCanteen)

	if err != nil {
		t.Errorf("_canteen name is: %s and _differentNameCanteen name is: %s, which are different but the error: %s was returned while adding _differentNameCanteen", _canteen.Name, _differentNameCanteen.Name, err)
	}

	availableCanteens = _school.Canteens()

	if lena := len(availableCanteens); lena != 2 {
		t.Errorf("The number of existing schools should now be 2 but got: %d", lena)
	}
}
