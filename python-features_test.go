package pythonFeatures

import (
	"reflect"
	"testing"
)

func assertPanic(t *testing.T, panicValue interface{}, f func(array interface{}) (arr []interface{})) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as the expected behaviour ;;; ")
		}
	}()
	f(panicValue)
}

func assertPanicMinMax(t *testing.T, panicValue interface{}, f func(array interface{}) (arr interface{})) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected ;;; ")
		}
	}()
	f(panicValue)
}

func TestReverse(t *testing.T) {
	var dummyIntSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reversedDummyIntSlice := reverse(dummyIntSlice)
	var expectedInt []interface{}
	expectedInt = append(expectedInt, 9, 8, 7, 6, 5, 4, 3, 2, 1)

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
	dummyInterface = append(dummyInterface, 3, "a", "b", 4, "d", 5, 3.50, "e", 2.50)
	reversedDummyInterface := reverse(dummyInterface)
	var expectedInterface []interface{}
	expectedInterface = append(expectedInterface, 2.50, "e", 3.50, 5, "d", 4, "b", "a", 3)

	if !reflect.DeepEqual(reversedDummyInterface, expectedInterface) {
		t.Errorf("Reversed interface slice is not the same ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedInterface, reversedDummyInterface)
	}
}

func TestMin(t *testing.T) {
	var dummyInt = []int{1, 2, 3, -1, -2, 4, 5, -6, 7, -20}
	minValue := min(dummyInt)
	expectedValue := -20
	if minValue != expectedValue {
		t.Errorf("Minimum value of the int slice is incorrect ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedValue, minValue)
	}

	var dummyFloat = []float64{1.0, 2.0, 0.15, 0.04, 3.2, 0.015}
	minValueFloat := min(dummyFloat)
	expectedValueFloat := 0.015
	if minValueFloat != expectedValueFloat {
		t.Errorf("Minimum value of the float slice is incorrect ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedValueFloat, minValueFloat)
	}

	var dummyString = []string{"4", "1", "5", "0", "6", "7", "a", "b", "c", "dd"}
	minValueString := min(dummyString)
	expectedValueString := "0"
	if minValueString != expectedValueString {
		t.Errorf("Minimum value of the string slice is incorrect ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedValueString, minValueString)
	}

	assertPanicMinMax(t, []float32{1, 2, 3, 4}, min)
}

func TestMax(t *testing.T) {
	var dummyInt = []int{1, 2, 3, -1, -2, 4, 5, -6, 7, -20}
	maxValue := max(dummyInt)
	expectedValue := 7
	if maxValue != expectedValue {
		t.Errorf("Maximum value of the int slice is incorrect ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedValue, maxValue)
	}

	var dummyFloat = []float64{1.0, 2.0, 0.15, 0.04, 3.2, 0.015}
	maxValueFloat := max(dummyFloat)
	expectedValueFloat := 3.2
	if maxValueFloat != expectedValueFloat {
		t.Errorf("Maximum value of the float slice is incorrect ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedValueFloat, maxValueFloat)
	}

	var dummyString = []string{"4", "1", "5", "0", "6", "7", "a", "b", "c", "dd"}
	maxValueString := max(dummyString)
	expectedValueString := "dd"
	if maxValueString != expectedValueString {
		t.Errorf("Maximum value of the string slice is incorrect ;;; \n"+
			" Expected : %v, \n Got : %v.", expectedValueString, maxValueString)
	}

	assertPanicMinMax(t, []float32{1, 2, 3, 4}, max)
}
func TestValidate(t *testing.T) {
	assertPanic(t, 5, validate)
}
func TestTakeArg(t *testing.T) {
	var dummyArg = []int{1, 2, 3}
	slice, ok := takeArg(dummyArg, reflect.Slice)
	if !ok || slice.Index(0).Kind() != reflect.Int {
		t.Errorf("Unexpected takeArg function behaviour ;;; \n"+
			" Expected : %v, \n Got : %v.", reflect.Int, slice.Index(0).Kind())
	}

	slice2, ok2 := takeArg(5, reflect.Slice)
	if ok2 {
		t.Errorf("Unexpected takeArg function behaviour, int is not type of slice ;;; \n"+
			" Expected : %v, \n Got : %v.", reflect.Slice, slice2.Kind())
	}
}

func TestMakeArray(t *testing.T) {
	dummyArray, ok := makeArray(5)
	if ok || dummyArray != nil {
		t.Errorf("Given input is not an array therefore dummyArray should be nil ;;; \n"+
			" Expected : %v, \n Got : %v.", nil, dummyArray)
	}

	dummyArray2, ok2 := makeArray([]int{1, 2, 3, 4, 5})
	if !ok2 || reflect.ValueOf(dummyArray2).Kind() != reflect.Slice {
		t.Errorf("Given input is a proper array, unexpected behaviour ;;; \n"+
			" Expected : %v, \n Got : %v.", reflect.Slice, reflect.ValueOf(dummyArray2).Kind())
	}

}
