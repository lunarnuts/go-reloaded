package student

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func PrintBST(node *TreeNode) {
	var sb string
	traversePreOrder(&sb, "", "", node, node.Parent != nil)
	fmt.Println(sb)
}
func traversePreOrder(sb *string, padding, pointer string, node *TreeNode, hasLeftSib bool) {
	//fmt.Println(sb)
	if node != nil {
		*sb += "\n"
		*sb += padding
		*sb += pointer
		*sb += node.Data

		var paddingBuilder = padding
		if hasLeftSib {
			paddingBuilder += "│  "
		} else {
			if node.Parent != nil {

				paddingBuilder += "   "
			}
		}

		paddingForBoth := paddingBuilder
		pointerForLeft := "└──"
		pointerForRight := ""
		if node.Left != nil {
			pointerForRight = "├──"
		} else {
			pointerForRight = "└──"
		}

		traversePreOrder(sb, paddingForBoth, pointerForRight, node.Right, node.Left != nil)
		traversePreOrder(sb, paddingForBoth, pointerForLeft, node.Left, false)
	}
}
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	} else {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	}
	return root
}
