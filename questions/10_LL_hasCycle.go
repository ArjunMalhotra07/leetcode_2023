package questions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func HasCycle(head *ListNode) bool {
	return HasCycleHelperFunction(head)
}
func HasCycleHelperFunction(head *ListNode) bool {
	slow := head
	fast := head
	//! Check if Linked List is Empty or Linked List is just 1 Element long return false as no ring can be made
	if head == nil || head.Next == nil {
		return false
	}
	//! Fast ptr shouldn't be nil or the next of the fast shouldn't be nil (can be the last)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		//! Checking if fast == slow means we have a cycle
		if slow == fast {
			return true
		}
	}
	return false
}
