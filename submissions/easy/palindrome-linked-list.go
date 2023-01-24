package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	middle := middleNode(head)
	halfReversed := reverseList(middle)

	for halfReversed != nil {
		if halfReversed.Val != head.Val {
			return false
		}

		halfReversed = halfReversed.Next
		head = head.Next
	}

	return true
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	current := head

	for current != nil {
		nxt := current.Next
		current.Next = prev
		prev = current
		current = nxt
	}

	return prev
}
