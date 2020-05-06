package student

type NodeI struct {
	Data int
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	ins := &NodeI{Data: data_ref}
	ins.Next, l = l, ins //insert data-ref into the head of list
	a := l
	for a.Next != nil {
		if a.Next.Data < a.Data { //simple bubble sort in the list
			a.Next.Data, a.Data = a.Data, a.Next.Data
			a = l
			continue
		}
		a = a.Next
	}
	return l
}
