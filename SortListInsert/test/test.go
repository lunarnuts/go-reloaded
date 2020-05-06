package main

import (
	"fmt"

	student ".."
)

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {

	var link *student.NodeI

	link = listPushBack(link, 0)
	PrintList(link)
	link = student.SortListInsert(link, 39)
	PrintList(link)
	fmt.Println("____________________")

	link = &student.NodeI{Data: 0}
	link = listPushBack(link, 1)
	link = listPushBack(link, 2)
	link = listPushBack(link, 3)
	link = listPushBack(link, 4)
	link = listPushBack(link, 5)
	link = listPushBack(link, 24)
	link = listPushBack(link, 25)
	link = listPushBack(link, 54)
	PrintList(link)
	link = student.SortListInsert(link, 33)
	PrintList(link)
	fmt.Println("____________________")

	link = &student.NodeI{Data: 0}
	link = listPushBack(link, 2)
	link = listPushBack(link, 18)
	link = listPushBack(link, 33)
	link = listPushBack(link, 37)
	link = listPushBack(link, 37)
	link = listPushBack(link, 39)
	link = listPushBack(link, 52)
	link = listPushBack(link, 53)
	link = listPushBack(link, 57)
	PrintList(link)
	link = student.SortListInsert(link, 53)
	PrintList(link)
	fmt.Println("____________________")

	link = &student.NodeI{Data: 0}
	link = listPushBack(link, 5)
	link = listPushBack(link, 18)
	link = listPushBack(link, 24)
	link = listPushBack(link, 28)
	link = listPushBack(link, 35)
	link = listPushBack(link, 42)
	link = listPushBack(link, 52)
	PrintList(link)
	link = student.SortListInsert(link, 52)
	PrintList(link)
	fmt.Println("____________________")

	link = &student.NodeI{Data: 0}
	link = listPushBack(link, 12)
	link = listPushBack(link, 20)
	link = listPushBack(link, 23)
	link = listPushBack(link, 23)
	link = listPushBack(link, 24)
	link = listPushBack(link, 30)
	link = listPushBack(link, 41)
	link = listPushBack(link, 53)
	link = listPushBack(link, 57)
	link = listPushBack(link, 59)
	PrintList(link)
	link = student.SortListInsert(link, 38)
	PrintList(link)
	fmt.Println("____________________")
}
