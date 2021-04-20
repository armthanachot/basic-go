package main
import(
	"fmt"
)
func main() {
	var arr1 []int =  [] int {1,2,3,4}
	var arr2 [3]string
	arr2[0],arr2[1],arr2[2] = "a","b","c"
	fmt.Println(arr1[0])
	for _,item := range arr1 {
		fmt.Println(item, ",")
	}
	for _,item := range arr2 {
		fmt.Println(item, ",")
	}
}
