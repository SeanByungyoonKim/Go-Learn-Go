package main

import "fmt"

func main() {
	nums := make([]int, 11)
	for i := 0; i < 11; i++ {
		nums[i] = i
		if nums[i]%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
