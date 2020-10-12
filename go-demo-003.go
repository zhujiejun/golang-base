package main

import "fmt"

func cal(z float64) float64 {
	n := 1
	m := 3
	t := 2.0
	sum := 2.0
	for {
		if t <= z {
			break
		}
		t = t * float64(n) / float64(m)
		sum = sum + t
		n++
		m = m + 2
	}
	return sum
}

func main() {
	sum := cal(1e-1234567890)
	fmt.Println(sum)
}
