package main

/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
https://leetcode.com/problems/multiply-strings/?envType=study-plan-v2&envId=programming-skills
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" { // When either is zero result would be 0
		return "0"
	}
	len1, len2 := len(num1), len(num2)
	len := len1 + len2
	ans := make([]byte, len) // stores final result and also store carry and sum while calculation
	for i := len1 - 1; i >= 0; i-- {
		a := num1[i] - '0'
		for j := len2 - 1; j >= 0; j-- {
			b := num2[j] - '0'
			idx := i + j + 1 // index in ans where multiplication would be inserted
			p := a*b + ans[idx]
			ans[idx] = p % 10
			ans[idx-1] += p / 10
		}
	}
	for i := 0; i < len; i++ {
		ans[i] += '0'
	}
	if ans[0] == '0' {
		return string(ans[1:])
	}
	return string(ans)
}
