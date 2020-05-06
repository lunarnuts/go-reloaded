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
	var link2 *student.NodeI
	link = nil
	link2 = nil
	PrintList(student.SortedListMerge(link, link2))
	fmt.Println("_________________________")

	link = nil
	link2 = nil
	link2 = listPushBack(link2, 2)
	link2 = listPushBack(link2, 2)
	link2 = listPushBack(link2, 4)
	link2 = listPushBack(link2, 9)
	link2 = listPushBack(link2, 12)
	link2 = listPushBack(link2, 12)
	link2 = listPushBack(link2, 19)
	link2 = listPushBack(link2, 20)

	PrintList(student.SortedListMerge(link, link2))
	fmt.Println("_________________________")

	link = nil
	link2 = nil
	link = listPushBack(link, 4)
	link = listPushBack(link, 4)
	link = listPushBack(link, 6)
	link = listPushBack(link, 9)
	link = listPushBack(link, 13)
	link = listPushBack(link, 18)
	link = listPushBack(link, 20)
	link = listPushBack(link, 20)

	PrintList(student.SortedListMerge(link, link2))
	fmt.Println("_________________________")

	link = nil
	link2 = nil
	link = listPushBack(link, 0)
	link = listPushBack(link, 7)
	link = listPushBack(link, 39)
	link = listPushBack(link, 92)
	link = listPushBack(link, 97)
	link = listPushBack(link, 93)
	link = listPushBack(link, 91)
	link = listPushBack(link, 28)
	link = listPushBack(link, 64)
	link2 = listPushBack(link2, 80)
	link2 = listPushBack(link2, 23)
	link2 = listPushBack(link2, 27)
	link2 = listPushBack(link2, 30)
	link2 = listPushBack(link2, 85)
	link2 = listPushBack(link2, 81)
	link2 = listPushBack(link2, 75)
	link2 = listPushBack(link2, 70)

	PrintList(student.SortedListMerge(link, link2))
	fmt.Println("_________________________")

	link = nil
	link2 = nil
	link = listPushBack(link, 0)
	link = listPushBack(link, 2)
	link = listPushBack(link, 11)
	link = listPushBack(link, 30)
	link = listPushBack(link, 54)
	link = listPushBack(link, 56)
	link = listPushBack(link, 70)
	link = listPushBack(link, 79)
	link = listPushBack(link, 99)
	link2 = listPushBack(link2, 23)
	link2 = listPushBack(link2, 28)
	link2 = listPushBack(link2, 38)
	link2 = listPushBack(link2, 67)
	link2 = listPushBack(link2, 67)
	link2 = listPushBack(link2, 79)
	link2 = listPushBack(link2, 95)
	link2 = listPushBack(link2, 97)

	PrintList(student.SortedListMerge(link, link2))
	fmt.Println("_________________________")

	link = nil
	link2 = nil
	link = listPushBack(link, 0)
	link = listPushBack(link, 3)
	link = listPushBack(link, 8)
	link = listPushBack(link, 8)
	link = listPushBack(link, 13)
	link = listPushBack(link, 19)
	link = listPushBack(link, 34)
	link = listPushBack(link, 38)
	link = listPushBack(link, 46)
	link2 = listPushBack(link2, 7)
	link2 = listPushBack(link2, 39)
	link2 = listPushBack(link2, 45)
	link2 = listPushBack(link2, 53)
	link2 = listPushBack(link2, 59)
	link2 = listPushBack(link2, 70)
	link2 = listPushBack(link2, 76)
	link2 = listPushBack(link2, 79)

	PrintList(student.SortedListMerge(link, link2))
	fmt.Println("_________________________")
}
