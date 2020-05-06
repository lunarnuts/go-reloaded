package main

import (
	"fmt"

	student ".."
)

func main() {
	fmt.Println(student.ConvertBase("4506C", "0123456789ABCDEF", "choumi"))
	fmt.Println(student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF"))
	fmt.Println(student.ConvertBase("256850", "0123456789", "01"))
	fmt.Println(student.ConvertBase("uuhoumo", "choumi", "Zone01"))
	fmt.Println(student.ConvertBase("683241", "0123456789", "0123456789"))

}
