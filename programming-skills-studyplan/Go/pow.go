package main

/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
https://leetcode.com/problems/powx-n/?envType=study-plan-v2&envId=programming-skills
*/

func IterativePow(x float64, n int) float64 {
	if n < 0 { // when n is negative
		n = -1 * n
		x = 1.0 / x
	}
	result := 1.0
	for n > 0 {
		if n%2 == 1 { // multiple x only when n is odd
			result *= x
			n--
		}
		x = x * x
		n /= 2
	}
	return result
}

func RecursivePow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1.0 / RecursivePow(x, -1*n)
	}
	if n%2 == 1 {
		return x * RecursivePow(x*x, (n-1)/2)
	}
	return RecursivePow(x*x, n/2)
}
