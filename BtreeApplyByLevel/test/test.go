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
	fmt.Println("Tree:")
	student.PrintBST(root)
	fmt.Println("_______________")
	fmt.Println("After:")
	student.BTreeApplyByLevel(root, fmt.Println)
	fmt.Println("_______________")

	root = &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root, "07")
	student.BTreeInsertData(root, "12")
	student.BTreeInsertData(root, "10")
	student.BTreeInsertData(root, "05")
	student.BTreeInsertData(root, "02")
	student.BTreeInsertData(root, "03")
	fmt.Println("Tree:")
	student.PrintBST(root)
	fmt.Println("_______________")
	fmt.Println("After:")
	student.BTreeApplyByLevel(root, fmt.Print)
	fmt.Println("\n_______________")
}
