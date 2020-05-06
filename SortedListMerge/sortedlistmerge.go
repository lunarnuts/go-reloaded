package piscine

type NodeI struct {
	Data int
	Next *NodeI
}
type List struct {
	head *NodeI
	tail *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	head := n1
	a := n1
	if n1 == nil && n2 == nil {
		return nil
	}
	if n1 != nil {
		for a.Next != nil {
			a = a.Next
		}
		a.Next = n2 //merge two lists
		a = n1
		head = n1
	} else {
		a = n2
		head = n2
	}
	for a.Next != nil {
		if a.Next.Data < a.Data { //bubble sort the list
			a.Next.Data, a.Data = a.Data, a.Next.Data
			a = head
			continue
		}
		a = a.Next
	}
	return head
}
