// ! https://leetcode.com/problems/reverse-linked-list/
package questions

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {

	var prev, next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
