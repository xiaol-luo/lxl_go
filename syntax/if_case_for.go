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

func case_example() {

}

func for_example() {
	// while(true)
	var sum int = 0
	for {
		sum ++
		fmt.Println("sum ", sum)
		if sum > 3 {
			break
		}
	}

	// while(condition)
	sum = 0
	for sum < 3 {
		sum ++
	}

	// for (i=0; i<10; i++)
	for sum=0; sum<3; sum++ {

	}
}

func main() {
	if_example()
	case_example()
	for_example()
}
