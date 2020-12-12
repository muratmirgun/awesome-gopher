package main

import (
	"fmt"
	"math/rand"
	"time"
)

func threerandom() (int, int, int) {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)
	z := rand.Intn(10)

	return x, y, z
}

func main() {

	r1, r2, r3 := threerandom()
	fmt.Println(r1, r2, r3)
}
