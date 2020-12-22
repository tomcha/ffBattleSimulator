package main

import (
	"math/rand"
	"time"
)

func main() {
}

func dice(times int, hedron int) int {
	rand.Seed(time.Now().UnixNano())
	total := 0
	for i := 0; i < times; i++ {
		total += rand.Intn(hedron) + 1
	}
	return total
}
