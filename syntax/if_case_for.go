package main

import "fmt"

func if_example() {
	var a bool = false
	var b bool = true
	if a {
		fmt.Println("reach if 1")
	} else if b {
		fmt.Println("reach if 2")
	} else {
		fmt.Println("reach if 3")
	}
}

func case_example(a string) {
	switch a {
	case "123", "223":
		fmt.Println("hit 123 or 223")
	case "aabb":
		fmt.Println("hit aabb")
	default:
		fmt.Println("hit default")
	}
}

func for_example() {
	// while(true)
	var sum int = 0
	for {
		sum ++
		fmt.Println("sum  a", sum)
		if sum > 3 {
			break
		}
	}

	// while(condition)
	sum = 0
	for sum < 3 {
		sum ++
		fmt.Println("sum  b", sum)
	}

	// for (i=0; i<10; i++)
	for sum=0; sum<3; sum++ {
		fmt.Println("sum  c", sum)
	}

	// for range
	for i, v := range [...]int{1, 2, 3} {
		fmt.Println(fmt.Sprintf("for range %d:%d", i, v))
	}
}

func main() {
	if_example()
	for _, v := range [...]string{"123", "223", "aabb", "ccdd"} {
		case_example(v)
	}
	for_example()
}
