package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 1234567890; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
