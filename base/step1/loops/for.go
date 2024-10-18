package main

import "fmt"

func InstanceFor() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	arr := []string{"a", "b"}
	for i, v := range arr {
		fmt.Println(i, v)
	}

}
