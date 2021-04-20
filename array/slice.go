package main

import "fmt"

func main() {
	var num = []int{1, 2, 3, 4, 5}
	showSlice(num)
	// ลบจากหัว
	num = num[1:len(num)]
	//เลือก ตั้งแต่  index 1 ถึง จนจบ แสดงว่า index 0 จะโดนตัดออก
	showSlice(num)
	// ลบจากท้าย
	num = num[0 : len(num)-1] //ลบตัวท้ายออก
	//เลือก ตั้งแต่  index 1 ถึง จนจบ แสดงว่า index 0 จะโดนตัดออก
	showSlice(num)

	// specific remove
	var numbers = []int{1, 2, 3, 4, 5}
	num = removeIndex(numbers, 2)
}
func showSlice(x []int) {
	fmt.Printf("len = %d , cap=%d , slice=%v\n", len(x), cap(x), x)
}

func removeIndex(num []int, index int) []int {
	return append(num[:index], num[index+1:]...)
	// append ข้อมูลทั้งหมดที่อยู่ใน array ใช้ ...
	//:index คือเริ่มจาก 0 ถึง index
}
