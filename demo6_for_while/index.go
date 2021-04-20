package main

import (
	"fmt"
)

func main() {
	forLoop(10)
	forLoopWhile()
	courseArr := []string{"Android","ios","react"}
	forEach(courseArr)
	forLoopWhileBreak()
}

func forLoop(count int) {
	for index := 0; index < count; index++ {
		fmt.Printf("index : %d\n", index)
	}
}

func forLoopWhile() {
	index := 0
	for index < 5 {
		fmt.Printf("index while : %d\n", index)
		index++
	}
}

func forLoopWhileBreak() {
	index := 0
	for true {
		fmt.Printf("index while : %d\n", index)
		index++
		if index > 5{ 
			fmt.Println("5 already")
			break
		}
	}
}


func forEach(course []string){
	for index, item := range course{
		fmt.Printf("%d , %s\n",index,item)
		
	} 

	for _, item := range course {
		fmt.Printf("%s\n",item)
	}
}