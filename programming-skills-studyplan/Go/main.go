package main

import "fmt"

func main() {
	binarySum := addBinary("11", "10")
	fmt.Println(fmt.Sprintf("binary sum calculated over string 11 and 10 is : %v", binarySum))

	fmt.Printf("Pow calculated recursively: %v\n", RecursivePow(2.0, 12))
	fmt.Printf("Pow calculated iteratively: %v\n", IterativePow(2.0, -12))

	fmt.Printf("Result of multiplicate of two string 123456789979 and 87765655429763: %v", multiply("123456789979", "87765655429763"))

}
