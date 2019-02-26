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
		t.Errorf("Reversed string slice is not the same ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedInterface, reversedDummyInterface)
	}
}
