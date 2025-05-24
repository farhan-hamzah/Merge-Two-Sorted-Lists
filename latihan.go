package main

import "fmt"

// Definisi struct ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Fungsi untuk menggabungkan dua linked list terurut
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

// Helper untuk membuat linked list dari slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// Helper untuk mencetak linked list
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

// Fungsi main untuk demo
func main() {
	// Buat dua linked list
	list1 := createList([]int{1, 3, 5})
	list2 := createList([]int{2, 4, 6})

	fmt.Print("List 1: ")
	printList(list1)

	fmt.Print("List 2: ")
	printList(list2)

	// Gabungkan kedua list
	merged := mergeTwoLists(list1, list2)

	fmt.Print("Merged List: ")
	printList(merged)
}
