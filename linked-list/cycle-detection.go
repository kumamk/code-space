package linkedlist

// node structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// 2 pointers
	var slow *ListNode
	var fast *ListNode

	slow = head
	fast = head
	found := false
	// check cycle in ll
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			found = true
			break
		}
	}

	if slow != fast || !found {
		return nil
	}

	var cycle *ListNode
	cycle = head
	// find cycle node in ll
	for cycle != slow {
		cycle = cycle.Next
		slow = slow.Next
	}
	return cycle
}
