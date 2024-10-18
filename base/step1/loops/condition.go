package main

import "fmt"

func InstanceCondition() {
	// if else
	var a int = 100
	if a < 20 {
		fmt.Println("a 小于 20")
	} else {
		fmt.Println("a 不小于 20")
	}

	// switch
	var animal string = "cat"
	var votes int = 0

	// 不需要break,fallthrough 强制执行下一个case代码（不判断条件）
	switch animal {
	case "cat":
		votes = 15
	default:
		votes = 10
	}
	fmt.Println(animal, ":", votes)

	var grade string = "A"
	// 未曾见过的swicth用法,像多个if else
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "A", grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}

}
