package main

import "fmt"

func var_syntax()  {
	var nullStr string

	// empty_str_* is same
	var emptyStr1 string = ""
	var emptyStr2 = ""
	emptyStr3 := ""

	var pEmptyStr1 *string = &emptyStr1

	// array
	var zeroArrayOfInt [3]int;
	var initedArrayOfInt [3]int = [3]int {1, 2, 3}
	var sliceOfInt []int = initedArrayOfInt[:]

	fmt.Println("nullStr is ", nullStr)
	fmt.Println("emptyStr1 is ", emptyStr1)
	fmt.Println("emptyStr2 is ", emptyStr2)
	fmt.Println("emptyStr3 is ", emptyStr3)
	fmt.Println("pEmptyStr1 is ", pEmptyStr1)
	fmt.Println("zeroArrayOfInt is ", zeroArrayOfInt)
	fmt.Println("initedArrayOfInt is ", initedArrayOfInt)
	fmt.Println("sliceOfInt is ", sliceOfInt)
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var_syntax()

	var initedArrayOfInt [3]int = [3]int{1, 2, 3}
	var sliceOfInt []int = initedArrayOfInt[:]
	var sumOfArray = sum(initedArrayOfInt[:])
	var outStr string = fmt.Sprintf("sumOfArray is %d", sumOfArray)
	fmt.Print(outStr)
	var sumOfSlice = sum(sliceOfInt)
	outStr = fmt.Sprintf("sumOfSlice is %d", sumOfSlice)
	fmt.Print(outStr)
}
