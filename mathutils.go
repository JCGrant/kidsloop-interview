package main

import "math"

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func sumOfDigits(n int) (result int) {
	for n > 0 {
		result += n % 10
		n = n / 10
	}
	return
}
