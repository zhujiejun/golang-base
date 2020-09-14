package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now().UnixNano()
	for i := 1; i <= 1234567890123; i++ {
		result += i
	}
	end := time.Now().UnixNano()
	//fmt.Println("the result is ", result, "and total time consumption is ", (end-start)/1e9, " seconds.")
	fmt.Printf("the result is %d and total time consumption is %d seconds.\n", result, ((end - start) / 1e9))
}
