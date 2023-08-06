package main

/*
Given two binary strings a and b, return their sum as a binary string.
https://leetcode.com/problems/add-binary/?envType=study-plan-v2&envId=programming-skills
*/

func addBinary(a string, b string) string {
	aright, bright := len(a)-1, len(b)-1
	carry := byte(0)
	max := aright + 2
	if max < bright+2 { // calculate max length of the given string
		max = bright + 2
	}
	ans := make([]byte, max)
	for aright >= 0 || bright >= 0 {
		sum := carry
		if aright >= 0 {
			sum += byte(a[aright] - '0')
		}
		if bright >= 0 {
			sum += byte(b[bright] - '0')
		}
		carry = sum / 2
		sum = (sum % 2)
		ans[max-1] = (sum + '0')
		max--
		aright--
		bright--
	}
	if carry > 0 {
		ans[0] = (carry + '0')
		return string(ans)
	}
	return string(ans[1:])
}
