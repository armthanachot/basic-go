package main

import (
	"fmt"
)

func main() {
	var numbers []int
	// ไม่มีากรระบุขนาด เรียกว่า slide
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	showSlide(numbers)
	var numbers2 = make([]int, 0, 5) //จอง buffer ไว้ 5 ตัว
	numbers2 = append(numbers2, 6)
	numbers2 = append(numbers2, 3)
	numbers2 = append(numbers2, 9)
	numbers2 = append(numbers2, 9)
	numbers2 = append(numbers2, 9)
	numbers2 = append(numbers2, 9) //เราทำการ append เกิน 5 ตัว ระบบจะทำการ oversize แต่จะต้องเพิ่มทีละ 5 เท่ากับที่จองไว้ 
	showSlide(numbers2)

}

func showSlide(x []int) {
	fmt.Printf("len = %d , cap=%d , slice=%v\n", len(x), cap(x), x)
	// ขนาด , ความจุ , ค่า
}
