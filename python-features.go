package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

func reverse(array interface{}) []interface{} {
	arr, ok := makeArray(array)

	if !ok {
		panic(fmt.Sprintf("Type %v is not an array. \n Expected: Any Array Type \n Got: %v", reflect.ValueOf(array).Kind(), reflect.ValueOf(array).Kind()))
		return nil
	}

	for fIndex, lIndex := 0, len(arr)-1; fIndex < lIndex; fIndex, lIndex = fIndex+1, lIndex-1 {
		arr[fIndex], arr[lIndex] = arr[lIndex], arr[fIndex]
	}

	return arr

}

func min(inputSlice interface{}) (minValue interface{}) {
	arr, _ := makeArray(inputSlice)

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

func max(inputSlice interface{}) (minValue interface{}) {
	arr, _ := makeArray(inputSlice)

	switch valueType := arr[0].(type) {
	case int:
		finalSlice := inputSlice.([]int)
		sort.Ints(finalSlice)
		minValue = finalSlice[len(arr) - 1]
	case string:
		finalSlice := inputSlice.([]string)
		sort.Strings(finalSlice)
		minValue = finalSlice[len(arr) - 1]
	case float64:
		finalSlice := inputSlice.([]float64)
		sort.Float64s(finalSlice)
		minValue = finalSlice[len(arr) - 1]
	default:
		panic(fmt.Sprintf("Unknown type %v", valueType))
	}
	return
}

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

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

func main() {
	var numbers []int
	for i := 0; i < 1000000; i++ {
		numbers = append(numbers, i)
	}

	start := time.Now()
	numbers = append(numbers, -1, -2, -3, 999999)
	fmt.Println("Minimum of an array is ", max(numbers))
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed Time 1 : ", elapsed)

	start = time.Now()

	fmt.Println(reverse(numbers)[0])

	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("Elapsed Time 2 : ", elapsed)

	fmt.Println(reverse([]string{"me", "te", "me", "te", "me", "te", "me", "te", "me", "te", "me", "te", "me", "te"}))
	fmt.Println(reverse([]int{1, 2}))
	fmt.Println(reverse([]float64{1.5, 2.5}))
}
