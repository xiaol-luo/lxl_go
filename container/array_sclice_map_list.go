package main

import (
	"container/list"
	"fmt"
)

func test_array() {
	vec_1 := [3]int{1, 2, 3}
	fmt.Printf("test_array vec_1: %#v \n", vec_1)
	vec_1[1] = 10
	for k, v := range(vec_1) {
		fmt.Printf(fmt.Sprintf("vec_1[%d]= %v \n", k, v))
	}

	vec_2 := [3][2]int{ {1,2}, {3,4}, {5, 6} }
	fmt.Printf("test_array vec_2: %#v \n", vec_2)
	vec_2[2][1] = 10
	for k1, v1 := range(vec_2) {
		for k2, v2 := range(v1) {
			fmt.Printf("vec_2[%d][%d]=%v\n", k1, k2, v2)
		}
	}
}


func test_slice() {
	vec_1 := [3]int{1, 2, 3}
	s1 := vec_1[1:2] // 从数组中截取
	fmt.Printf("test_slice s1=%#v \n", s1)

	s2 := []int{1, 2} // 直接声明
	fmt.Printf("test_slice s2=%#v \n", s2)

	s3 := make([]int, 2, 10) // 通过make构造
	fmt.Printf("test_slice s3=%#v \n", s3)

	for i:=10; i<13; i++ {
		s1 = append(s1, i)
	}
	fmt.Printf("test_slice append s1=%#v \n", s1)
	copy(s3, s1)
	fmt.Printf("test_slice copy s3=%#v \n", s3)
}

func test_map() {
	m1 := map[int]string{1:"one", 2:"two"}
	fmt.Printf("test_map m1=%#v\n", m1)

	var m2 map[int]string
	fmt.Printf("test_map m2=%#v\n", m2)

	m3 := make(map[int]string, 10)
	m3[1] = "one"
	m3[2] = "two"
	m3[3] = "three"
	m3[4] = "four"
	m3[5] = "five"
	fmt.Printf("test_map m3=%#v\n", m3)

	delete(m3, 1)
	fmt.Printf("test_map delete m3=%#v\n", m3)
	if v, is_exist := m3[2]; is_exist {
		fmt.Printf("test_map m3[%d]=%v\n", 2, v)
	}
}

func test_list() {
	var l1 list.List
	// l2 := list.New()
	e1 := l1.PushBack(1)
	e2 := l1.PushFront(2)
	fmt.Printf("test list init l1=%v\n", l1)
	l1.MoveAfter(e2, e1)
	fmt.Printf("test list move after l1=%v\n", l1)
}

func main() {
	test_array()
	test_slice()
	test_map()
	test_list()
}
