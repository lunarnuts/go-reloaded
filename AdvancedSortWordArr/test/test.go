package main

import (
	"fmt"
	"strings"

	student ".."
)

func f(a, b string) int {
	return strings.Compare(b, a)
}
func main() {
	array := "The earliest computing device undoubtedly consisted of the five fingers of each hand"

	result := student.SplitWhiteSpaces(array)
	student.AdvancedSortWordArr(result, strings.Compare)

	fmt.Println(result)
	array = "The word digital comesfrom \"digits\" or fingers"

	result = student.SplitWhiteSpaces(array)
	student.AdvancedSortWordArr(result, strings.Compare)

	fmt.Println(result)
	array = "Ta A 1 b B 2 c C 3"

	result = student.SplitWhiteSpaces(array)
	student.AdvancedSortWordArr(result, strings.Compare)

	fmt.Println(result)
	array = "The computing consisted device each earliest fingers five hand of of the undoubtedly"

	result = student.SplitWhiteSpaces(array)
	student.AdvancedSortWordArr(result, f)
	fmt.Println(result)

	array = "\"digits\" The comesfrom digital fingers or word"

	result = student.SplitWhiteSpaces(array)
	student.AdvancedSortWordArr(result, func(a, b string) int {
		return strings.Compare(b, a)
	})
	fmt.Println(result)

	array = "1 2 3 A B C a b c"

	result = student.SplitWhiteSpaces(array)
	student.AdvancedSortWordArr(result, func(a, b string) int {
		return strings.Compare(b, a)
	})
	fmt.Println(result)
}
