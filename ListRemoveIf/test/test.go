package main

import (
	"fmt"

	student ".."
)

func PrintList(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &student.List{}
	link2 := &student.List{}

	fmt.Println("----normal state----")
	//student.ListPushBack(link2, 1)
	PrintList(link2)
	student.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()
	link = &student.List{}
	link2 = &student.List{}

	fmt.Println("----normal state----")
	//student.ListPushBack(link2, 1)
	PrintList(link2)
	student.ListRemoveIf(link2, 96)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	link = &student.List{}
	fmt.Println("----normal state----")
	student.ListPushBack(link, 98)
	student.ListPushBack(link, 98)
	student.ListPushBack(link, 33)
	student.ListPushBack(link, 34)
	student.ListPushBack(link, 33)
	student.ListPushBack(link, 34)
	student.ListPushBack(link, 33)
	student.ListPushBack(link, 89)
	student.ListPushBack(link, 33)
	PrintList(link)
	student.ListRemoveIf(link, 34)
	fmt.Println("------answer-----")
	PrintList(link)
	fmt.Println()

	link = &student.List{}
	fmt.Println("----normal state----")
	student.ListPushBack(link, 79)
	student.ListPushBack(link, 74)
	student.ListPushBack(link, 99)
	student.ListPushBack(link, 79)
	student.ListPushBack(link, 7)
	PrintList(link)
	student.ListRemoveIf(link, 99)
	fmt.Println("------answer-----")
	PrintList(link)
	fmt.Println()

	link = &student.List{}
	fmt.Println("----normal state----")
	student.ListPushBack(link, 56)
	student.ListPushBack(link, 93)
	student.ListPushBack(link, 68)
	student.ListPushBack(link, 56)
	student.ListPushBack(link, 87)
	student.ListPushBack(link, 68)
	student.ListPushBack(link, 56)
	student.ListPushBack(link, 68)
	PrintList(link)
	student.ListRemoveIf(link, 68)
	fmt.Println("------answer-----")
	PrintList(link)
	fmt.Println()

	link = &student.List{}
	fmt.Println("----normal state----")
	student.ListPushBack(link, "mvkUxbqhQve4l")
	student.ListPushBack(link, "4Zc4t hnf SQ")
	student.ListPushBack(link, "q2If E8BPuX ")
	PrintList(link)
	student.ListRemoveIf(link, "4Zc4t hnf SQ")
	fmt.Println("------answer-----")
	PrintList(link)
	fmt.Println()
}
