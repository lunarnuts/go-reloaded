package main

import (
	"fmt"

	student ".."
)

func main() {
	root := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root, "07")
	student.BTreeInsertData(root, "12")
	student.BTreeInsertData(root, "10")
	student.BTreeInsertData(root, "05")
	student.BTreeInsertData(root, "02")
	student.BTreeInsertData(root, "03")
	node := student.BTreeSearchItem(root, "02")
	fmt.Println("Before delete:")
	student.PrintBST(root)
	fmt.Println("__________________\n")
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.PrintBST(root)
	fmt.Println("__________________\n")

	root = &student.TreeNode{Data: "33"}
	student.BTreeInsertData(root, "5")
	student.BTreeInsertData(root, "52")
	student.BTreeInsertData(root, "20")
	student.BTreeInsertData(root, "31")
	student.BTreeInsertData(root, "13")
	student.BTreeInsertData(root, "11")
	node = student.BTreeSearchItem(root, "20")
	fmt.Println("Before delete:")
	student.PrintBST(root)
	fmt.Println("__________________\n")
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.PrintBST(root)
	fmt.Println("__________________\n")

	root = &student.TreeNode{Data: "03"}
	student.BTreeInsertData(root, "39")
	student.BTreeInsertData(root, "99")
	student.BTreeInsertData(root, "44")
	student.BTreeInsertData(root, "11")
	student.BTreeInsertData(root, "14")
	student.BTreeInsertData(root, "11")
	node = student.BTreeSearchItem(root, "03")
	fmt.Println("Before delete:")
	student.PrintBST(root)
	fmt.Println("__________________\n")
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.PrintBST(root)
	fmt.Println("__________________\n")

	root = &student.TreeNode{Data: "03"}
	root.Right = &student.TreeNode{Data: "03"}
	root.Right.Parent = root
	student.BTreeInsertData(root, "01")
	student.BTreeInsertData(root, "94")
	student.BTreeInsertData(root, "19")
	student.BTreeInsertData(root, "24")
	student.BTreeInsertData(root, "111")
	node = student.BTreeSearchItem(root, "03")
	fmt.Println("Before delete:")
	student.PrintBST(root)
	fmt.Println("__________________\n")
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.PrintBST(root)
	fmt.Println("__________________\n")
}
