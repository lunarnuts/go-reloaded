package student

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}

	if root.Data > node.Data {
		if root.Left.Data == node.Data {
			root.Left = rplc
			rplc.Parent = root
			return root
		} else {
			BTreeTransplant(root.Left, node, rplc)
		}
	} else if root.Data < node.Data {
		if root.Right.Data == node.Data {
			root.Right = rplc
			rplc.Parent = root
			return root

		} else {
			BTreeTransplant(root.Right, node, rplc)
		}
	} else {
		return rplc
	}
	return root
}
