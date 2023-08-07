package main

import (
	"fmt"
	"log"
)

func main() {
	binarySum := addBinary("11", "10")
	fmt.Printf("binary sum calculated over string 11 and 10 is : %v", binarySum)

	fmt.Printf("Pow calculated recursively: %v\n", RecursivePow(2.0, 12))
	fmt.Printf("Pow calculated iteratively: %v\n", IterativePow(2.0, -12))

	fmt.Printf("Result of multiplicate of two string 123456789979 and 87765655429763: %v\n", multiply("123456789979", "87765655429763"))
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 5}
	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}
	list := mergeTwoLists(list1, list2)
	log.Printf("Merge List from list1 %v and list2 %v is: %v", list1, list2, list)

	fmt.Printf("reverse of List: %v", reverseList(list))

}
