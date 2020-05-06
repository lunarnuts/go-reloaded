package student

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}

	if root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}
