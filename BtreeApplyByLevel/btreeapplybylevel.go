package student

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	queue := make([]*TreeNode, 0)
	for root != nil {
		f(root.Data)
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if len(queue) != 0 {
			root = queue[0]
			queue = queue[1:]
		} else {
			root = nil
		}

	}

}
