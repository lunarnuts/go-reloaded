package student

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}
	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return root
	}

}
