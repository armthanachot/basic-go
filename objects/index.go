// objects in go call map
package main

import (
	"fmt"
)

func main() {
	var number = map[string]int{"one": 1, "two": 2, "three": 3} //string คือ type ของ key int คือ type ของ value
	fmt.Println("", number)
	fmt.Println("", number["two"])       //access key
	var colors = make(map[string]string) //ประกาศแบบเป็น map เปล่าๆ ไม่มีการ ใส่ key กับ value เริ่มต้นให้ เหมือนกับ x = {}
	colors["red"] = "#f00"               //สร้าง key red ลงไปที่ map colors
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println("", colors)
	var course = make(map[string]map[string]int)
	course["android"] = make(map[string]int) //ถ้าไม่ทำการ make ตรงนี้ จะไม่สามารถ เขียนซ้อนกันได้แบบด้านล่าง เพราะ ตัว root ไม่มีค่า
	course["android"]["price"] = 322123
	course["android"]["price"] = 322124
	course["ios"] = make(map[string]int) //เป็นการสร้าง root key อีกตัวใน level แรก ต้องมีการ make ก่อนเสมอ
	course["ios"]["price"] = 322124
	fmt.Println(course)
}
