package pythonFeatures

import (
	"fmt"
	"reflect"
	"sort"
)

// It reverses given array and returns it as an []interface
// Could take any type as the array parameter however if the parameter is not a slice/array it will raise an error & panic
func reverse(array interface{}) (arr []interface{}) {
	arr, ok := makeArray(array)

	if !ok {
		panic(fmt.Sprintf("Type %v is not an array. \n Expected: Any Array Type \n Got: %v", reflect.ValueOf(array).Kind(), reflect.ValueOf(array).Kind()))
		return nil
	}

	for fIndex, lIndex := 0, len(arr)-1; fIndex < lIndex; fIndex, lIndex = fIndex+1, lIndex-1 {
		arr[fIndex], arr[lIndex] = arr[lIndex], arr[fIndex]
	}

	return

}

// Can take any type as the inputSlice parameter, however if the parameter is not a slice/array it will raise an error & panic
// If the given parameter is an array/slice it will check which type of an array it is and then sort it by the type
// And finally find the minimum value by getting the index 0 since it will be the minimum value after sorting
// Panics if parameter's type is not known
func min(inputSlice interface{}) (minValue interface{}) {
	arr, ok := makeArray(inputSlice)

	if !ok {
		panic(fmt.Sprintf("Type %v is not an array. \n Expected: Any Array Type \n Got: %v", reflect.ValueOf(arr).Kind(), reflect.ValueOf(arr).Kind()))
		return nil
	}


	switch valueType := arr[0].(type) {
	case int:
		finalSlice := inputSlice.([]int)
		sort.Ints(finalSlice)
		minValue = finalSlice[0]
	case string:
		finalSlice := inputSlice.([]string)
		sort.Strings(finalSlice)
		minValue = finalSlice[0]
	case float64:
		finalSlice := inputSlice.([]float64)
		sort.Float64s(finalSlice)
		minValue = finalSlice[0]
	default:
		panic(fmt.Sprintf("Unknown type %v", valueType))
	}
	return
}

// Can take any type as the inputSlice parameter, however if the parameter is not a slice/array it will raise an error & panic
// If the given parameter is an array/slice it will check which type of an array it is and then sort it by the type
// And finally find the maximum value by getting the index 0 since it will be the maximum value after sorting
// Panics if parameter's type is not known
func max(inputSlice interface{}) (maxValue interface{}) {
	arr, ok := makeArray(inputSlice)

	if !ok {
		panic(fmt.Sprintf("Type %v is not an array. \n Expected: Any Array Type \n Got: %v", reflect.ValueOf(arr).Kind(), reflect.ValueOf(arr).Kind()))
		return nil
	}

	switch valueType := arr[0].(type) {
	case int:
		finalSlice := inputSlice.([]int)
		sort.Ints(finalSlice)
		maxValue = finalSlice[len(arr) - 1]
	case string:
		finalSlice := inputSlice.([]string)
		sort.Strings(finalSlice)
		maxValue = finalSlice[len(arr) - 1]
	case float64:
		finalSlice := inputSlice.([]float64)
		sort.Float64s(finalSlice)
		maxValue = finalSlice[len(arr) - 1]
	default:
		panic(fmt.Sprintf("Unknown type %v", valueType))
	}
	return
}

// Calls the takeArg func and if success bool is false return nil interface and false bool
// If true, creates an []interface and maps the values from any type array to []interface
func makeArray(array interface{}) (out []interface{}, ok bool) {
	slice, success := takeArg(array, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)

	for i := 0; i < c; i++ {

		out[i] = slice.Index(i).Interface()

	}
	return out, true
}

// Checks if given input is type of slice. If it is returns true, else it returns false as default
func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

