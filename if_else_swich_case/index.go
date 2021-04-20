package main

import (
	"fmt"
)

func main() {
	someVal := 15
	if someVal == 10 {
		fmt.Println(" == 10 ")
	} else {
		fmt.Println(" != 10")
	}

	if someVal == 10 || someVal < 10 {
		fmt.Println("equals or less than 10")
	} else {
		fmt.Println("not equals or more than 10")
	}

	swCase()
}

func swCase() {
	index := 5
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Println("not entry case")
		break
	}
}
