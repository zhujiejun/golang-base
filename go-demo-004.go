package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	inCircleCount := 0
	var x, y float64
	var Pi float64
	for i := 0; i < 123456789; i++ {
		x = rand.Float64()
		y = rand.Float64()

		if x*x+y*y < 1 {
			inCircleCount++
		}

		if i%1000 == 0 {
			Pi = (4.0 * float64(inCircleCount)) / float64(i)
			fmt.Println("calculated", i, "times, current PI is: ", Pi)
		}
	}
}
