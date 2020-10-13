package main

import (
	"fmt"
	"sync"
	"time"
)

func calculate(max int) {
	var lock sync.Mutex
	result := 0
	start := time.Now().UnixNano()
	for i := 1; i <= max; i++ {
		lock.Lock()
		result += i
		lock.Unlock()
	}
	end := time.Now().UnixNano()
	fmt.Printf("the result is %d and total time consumption is %d seconds.\n", result, ((end - start) / 1e9))
}

func main() {
	//input := bufio.NewScanner(os.Stdin)
	//for input.Scan() {
	//	in := input.Text()
	//	//fmt.Println("in:", in)
	//	calculate(
	//}
	var input int
	fmt.Println("please input the max number:")
	fmt.Scanln(&input)
	for k := 1; k <= 5; k++ {
		go calculate(input)
	}
	time.Sleep(time.Second)
}
