package main

import (
	"fmt"
	"time"
)

const cNumMax = 1234567890

func main() {
	sign := 1.0
	pi := 0.0
	t1 := time.Now()
	for i := 1; i < cNumMax+2; i += 2 {
		pi += (1.0 / float64(i)) * sign
		sign = -sign
	}
	pi *= 4
	t2 := time.Now()
	fmt.Printf("PI = %f; Time = %d\n", pi, t2.Sub(t1)/time.Millisecond)
}
