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
	var sliceOfInt2 []int = []int {1, 2, 3, 4}
	initedArrayOfIntWithVarNumbs := [...]int {1, 2, 3, 4, 5}

	var mapStrInt map[string]int = map[string]int {"one": 1, "two": 2}

	fmt.Println("nullStr is ", nullStr)
	fmt.Println("emptyStr1 is ", emptyStr1)
	fmt.Println("emptyStr2 is ", emptyStr2)
	fmt.Println("emptyStr3 is ", emptyStr3)
	fmt.Println("pEmptyStr1 is ", pEmptyStr1)
	fmt.Println("zeroArrayOfInt is ", zeroArrayOfInt)
	fmt.Println("initedArrayOfInt is ", initedArrayOfInt)
	fmt.Println("sliceOfInt is ", sliceOfInt)
	fmt.Println("sliceOfInt2 is ", sliceOfInt2)
	fmt.Println("initedArrayOfIntWithVarNumbs is ", initedArrayOfIntWithVarNumbs)
	fmt.Println("mapStrInt is ", mapStrInt)
}

// single return value
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

// multiple return values
func just_return(a int, b int) (int, int) {
	return a, b
}

// define struct Foo and struct function
type Foo struct {
	a int
	b string
}
// struct function
func (foo *Foo) ToString() string {
	var ret = fmt.Sprintf("Foo:: a=%d, b=%s", foo.a, foo.b)
	return ret
}

type Boo struct {
	c int
	d int
}

func (boo *Boo) ToString() string {
	var ret = fmt.Sprintf("Boo:: c=%d, d=%d", boo.c, boo.d)
	return ret
}

type TestInterface interface {
	ToString() string
}

func PrintStruct(struct_ins TestInterface) {
	fmt.Println(fmt.Sprintf("PrintStruct print struct : %s", struct_ins.ToString()))
}

func main() {
	var_syntax()
	var initedArrayOfInt [3]int = [3]int{1, 2, 3}
	var sliceOfInt []int = initedArrayOfInt[:]
	var sumOfArray = sum(initedArrayOfInt[:])
	var outStr string = fmt.Sprintf("sumOfArray is %d", sumOfArray)
	fmt.Println(outStr)
	var sumOfSlice = sum(sliceOfInt)
	outStr = fmt.Sprintf("sumOfSlice is %d", sumOfSlice)
	fmt.Println(outStr)

	var foo *Foo = new(Foo)
	foo.a = 100
	foo.b = "b"
	fmt.Println("foo is ", foo)

	ret_a, ret_b := just_return(100, 200)
	fmt.Println("just_return ret is", ret_a, ret_b)
	fmt.Println(fmt.Sprintf("foo.ToString() is %s", foo.ToString()))
	PrintStruct(foo)

	var boo = new(Boo)
	PrintStruct(boo)
}
