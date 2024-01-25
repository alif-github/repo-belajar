package main

import (
	"fmt"
	"regexp"
)

//// Arr ...
//type Arr struct {
//	value int32
//	index int
//}
//
//// ByValue ...
//type ByValue []Arr
//
//func (a ByValue) Len() int           { return len(a) }
//func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }
//
//func minimumSwaps(input []int32) int32 {
//	arr := make([]Arr, 0, len(input))
//
//	// step 1. Generate an array of sorted indexes
//	for i, v := range input {
//		arr = append(arr, Arr{v, i})
//	}
//
//	sort.Sort(ByValue(arr))
//
//	idx := make([]int, 0, len(input))
//	for _, ar := range arr {
//		idx = append(idx, ar.index)
//	}
//
//	// step 2. Sort the array by sorted indexes
//	var result int32
//	for i := 0; i < len(input); i++ {
//		if i == idx[i] {
//			continue
//		}
//
//		input[i], input[idx[i]] = input[idx[i]], input[i]
//		idx[i], idx[input[idx[i]]-1] = i, idx[i]
//		result++
//	}
//	return result
//}

func newMinimumSwaps(arr []int) int {
	count := 0
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			for arr[i] != i+1 {
				temp := 0
				temp = arr[arr[i]-1]
				arr[arr[i]-1] = arr[i]
				arr[i] = temp
				fmt.Println(arr)
				count++
			}
		}
	}

	return count
}

func reverseString(input string) string {
	storeSymb := ""
	byteStr := []rune(input)
	for i, j := 0, len(byteStr)-1; i < j; i, j = i+1, j-1 {
		rgx, _ := regexp.Compile("^[!%]+$")
		if j == len(byteStr)-1 && rgx.MatchString(string(byteStr[j])) {
			storeSymb = string(byteStr[j])
			byteStr = append(byteStr[:j], byteStr[j+1:]...)
			i, j = 0, len(byteStr)-1
		}

		byteStr[i], byteStr[j] = byteStr[j], byteStr[i]
	}

	if storeSymb != "" {
		s := []rune(storeSymb)
		byteStr = append(byteStr, s...)
	}

	return string(byteStr)
}

func main() {
	//fmt.Println(minimumSwaps([]int32{4, 3, 1, 2}))
	//fmt.Println(newMinimumSwaps([]int{2, 3, 4, 1, 5}))
	//fmt.Println(resultAddArray([]int{2, 3, 4, 1, 5}))
	fmt.Println(reverseString("ayalamih!"))
}

func resultAddArray(arr []int) int {
	result := 0
	for _, itemArr := range arr {
		result += itemArr
	}

	return result
}
