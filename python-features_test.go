package pythonFeatures

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	var dummyIntSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reversedDummyIntSlice := reverse(dummyIntSlice)
	var expectedInt []interface{}
	expectedInt = append(expectedInt, 9,8,7,6,5,4,3,2,1)

	if !reflect.DeepEqual(reversedDummyIntSlice, expectedInt) {
		t.Errorf("Reversed int slice is not the same ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedInt, reversedDummyIntSlice)
	}

	var dummyStringSlice = []string{"a", "b", "c", "d"}
	reversedDummyStringSlice := reverse(dummyStringSlice)
	var expectedString []interface{}
	expectedString = append(expectedString, "d", "c", "b", "a")

	if !reflect.DeepEqual(reversedDummyStringSlice, expectedString) {
		t.Errorf("Reversed string slice is not the same ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedString, reversedDummyStringSlice)
	}

	var dummyInterface []interface{}
	dummyInterface = append(dummyInterface, 3,"a","b",4,"d",5, 3.50, "e", 2.50)
	reversedDummyInterface := reverse(dummyInterface)
	var expectedInterface []interface{}
	expectedInterface = append(expectedInterface, 2.50, "e", 3.50, 5, "d", 4, "b", "a", 3)

	if !reflect.DeepEqual(reversedDummyInterface, expectedInterface) {
		t.Errorf("Reversed interface slice is not the same ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedInterface, reversedDummyInterface)
	}
}

func TestMin(t *testing.T){
	var dummyInt = []int{1,2,3,-1,-2,4,5,-6,7,-20}
	minValue := min(dummyInt)
	expectedValue := -20
	if minValue != expectedValue{
		t.Errorf("Minimum value of the int slice is incorrect ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedValue, minValue)
	}

	var dummyFloat = []float64{1.0, 2.0, 0.15, 0.04, 3.2, 0.015}
	minValueFloat := min(dummyFloat)
	expectedValueFloat := 0.015
	if minValueFloat != expectedValueFloat{
		t.Errorf("Minimum value of the int slice is incorrect ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedValueFloat, minValueFloat)
	}

	var dummyString = []string{"4", "1", "5", "0","6","7", "a", "b", "c", "dd"}
	minValueString := min(dummyString)
	expectedValueString := "0"
	if minValueString != expectedValueString{
		t.Errorf("Minimum value of the int slice is incorrect ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedValueString, minValueString)
	}
}
