package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randate() (int, int, int) {
	rand.Seed(time.Now().Unix())
	min_Y := 1971
	max_Y := 2070
	Y := rand.Intn(max_Y-min_Y) + min_Y
	min_M := 1
	max_M := 12
	M := rand.Intn(max_M-min_M) + min_M
	minD := 1
	maxD := 30
	D := rand.Intn(maxD-minD) + minD
	return Y, M, D
}

func sex() int {
	rand.Seed(time.Now().Unix())
	S := rand.Intn(8999) + 1000
	return S
}
func PESEL(Y int, M int, D int, S int) [11]int {
	var PESEL = [11]int{}
	var wage = [10]int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}
	PESEL[0] = (Y % 100) / 10
	PESEL[1] = Y % 10
	if M < 10 {
		PESEL[2] = 0
		PESEL[3] = M
	} else {
		PESEL[2] = (M % 100) / 10
		PESEL[3] = M % 10
	}
	if D < 10 {
		PESEL[4] = 0
		PESEL[5] = D
	} else {
		PESEL[4] = (D % 100) / 10
		PESEL[5] = D % 10
	}
	PESEL[6] = S / 1000
	PESEL[7] = S % 1000 / 100
	PESEL[8] = S % 100 / 10
	PESEL[9] = S % 10

	var c int
	for i := 0; i < len(wage); i++ {
		c += (PESEL[i] * wage[i]) % 10
	}
	PESEL[10] = c % 10
	return PESEL
}

func main() {
	Y, M, D := randate()
	fmt.Printf("Date of birth: %d-%d-%d\n", Y, M, D)
	S := sex()
	fmt.Println(S)
	if S%2 == 0 {
		fmt.Println("Your sex is woman")
	} else {
		fmt.Println("Your sex is man")
	}
	fmt.Println(PESEL(Y, M, D, S))
}
