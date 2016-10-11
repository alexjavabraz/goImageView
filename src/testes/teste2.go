package main

import (
	"fmt"
	"runtime"
	"time"
	)

func main(){
	sum := 0
	for i := 0; i < 10; i++{
		fmt.Printf("%v",sum)
		fmt.Printf("+")
		fmt.Printf("%v", i)
		sum += i
		fmt.Println("=",sum)
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		
		sum += sum
		const total = 100

		if sum > total {
			fmt.Println("%g>%g\n", sum, total)
			break
		}

	}
	fmt.Println(sum)
	fmt.Println(runtime.GOOS)
	fmt.Println(time.Now().Weekday())
	

}