// ! https://leetcode.com/problems/palindrome-linked-list/description/
package questions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var stack []ListNode

func Push(str ListNode) {
	stack = append(stack, str)
}
func Pop() {
	n := len(stack) - 1
	stack = stack[:n]
}
func IsPalindrome(head *ListNode) bool {
	one := head
	two := head
	for two != nil && two.Next != nil {
		Push(*one)
		two = two.Next.Next
		if two != nil && two.Next == nil {
			one = one.Next.Next
		} else {
			one = one.Next
		}

	}
	for one != nil {
		if one.Val != stack[len(stack)-1].Val {
			return false
		}
		one = one.Next
		Pop()
	}
	return true
}
