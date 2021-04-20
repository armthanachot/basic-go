package main

//ใน go ไม่มี class เราใข้ struct แทน
import (
	"fmt"
)

func main() {
	var p1 product
	p1.name = "Arduino"
	p1.price = 200
	p1.stock = 20
	fmt.Println(p1)
	fmt.Println(p1.name) //access key name
	p1.price = updatePrice(p1.price , 10)
	fmt.Println(p1)
	var discount float64 = p1.discount(50)
	fmt.Println("discount : ",discount)
	p1 = p1.clear()
	fmt.Println(p1)
}

type product struct {
	name  string
	price int
	stock int
} //ถ้าตั้งชื่อเป็นตัวใหญ่จะ warning แต่ ไม่ error มันแปลว่า เราจะแชร์ให้ ที่อื่นเข้ามาใช้ได้

func updatePrice(price int , add int) int{ //ถ้ามีการ return ก็ใส่ type ได้
	newprice := price+add
	return newprice
}

// ต้องประกาศว่าเป็น function ของ object ไหน ในที่นี้เป็น function ของ object product
func (p product) clear() product{ //return product ใหม่กลับไป
	p.price = 0
	p.stock = 0
	return p
} //สร้าง function ให้ object เวลาเรียกใช้ก็ .clear()

func (p product) discount(percent float64) float64{
	var productprice float64 = float64(p.price)
	var convertPercent float64 = percent/100
	var disc float64 = convertPercent
	var result float64 = 0.00
	result = productprice - (productprice * disc)
	fmt.Println(result)
	return result
}