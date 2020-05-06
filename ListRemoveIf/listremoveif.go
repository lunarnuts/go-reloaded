package student

type NodeI struct {
	Data interface{}
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	var prev *NodeI = nil
	curr := l.Head
	for curr != nil {
		if curr.Data != data_ref {
			prev, curr = curr, curr.Next
			continue
		}
		curr = curr.Next
		if prev == nil {
			l.Head = curr
		} else {
			prev.Next = curr
		}
	}
}
func ListPushBack(l *List, data interface{}) {
	n := &NodeI{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}
