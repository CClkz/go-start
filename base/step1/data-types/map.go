package main

import "fmt"

func tryMap() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	m["c"] = 3

	fmt.Println(m)
}
