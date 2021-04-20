package main

import (
	"fmt"
)

func main() {
	fn1()
	fn2("hello fn 2", 10)
	var num1, num2 = 2, 5
	fmt.Println(sum(num1, num2))
	fmt.Printf("%d + %d = %d\n", num1, num2, sum(num1, num2))
	var a, b = 999, 111
	var swa, swb = swap(a, b)
	fmt.Printf("swap %d , %d <=> %d , %d\n", a, b, swa, swb)
	var x, y, title = swapV2(a, b)
	fmt.Printf("swap V2 %d , %d <=> %d , %d : title = %s", a, b, x, y, title)
}

func fn1() {
	fmt.Println("demo 5")

}

func fn2(msg string, version int) {
	fmt.Println(msg)
	fmt.Println(version)

}

// return function
func sum(num1 int, num2 int) int {
	return num1 + num2
}

//multiple return
func swap(a int, b int) (int, int) {
	return b, a
}

//multiple return and naming variable
func swapV2(x int, y int) (x1, y1 int, title string) {
	x1 = y
	y1 = x
	title = "swap title"
	return x1, y1, title
}


