package easy

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	merged := &ListNode{Val: 0, Next: l1}

	result := merged

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			merged.Next = &ListNode{Val: l1.Val, Next: nil}
			l1 = l1.Next
		} else {
			merged.Next = &ListNode{Val: l2.Val, Next: nil}
			l2 = l2.Next
		}
		merged = merged.Next
	}

	for l1 != nil {
		merged.Next = &ListNode{Val: l1.Val, Next: nil}
		l1 = l1.Next
		merged = merged.Next
	}

	for l2 != nil {
		merged.Next = &ListNode{Val: l2.Val, Next: nil}
		l2 = l2.Next
		merged = merged.Next
	}

	return result.Next
}
