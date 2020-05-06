package main

import (
	"os"

	student ".."
	"github.com/01-edu/z01"
)

//func to mirror/reverse array
func Mirror(v []rune) {
	len := 0
	for i := range v {
		len = i + 1
	}
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		temp := v[i]
		v[i] = v[j]
		v[j] = temp
	}
}

func main() {
	s := os.Args
	s = s[1:]
	len := 0
	for range s {
		len++
	}
	if len != 0 {
		str := ""
		first := false
		for _, arg := range s {
			if first {
				str += " "
			}
			first = true
			str += arg
		}

		runes := []rune(str)
		var p []int
		var v []rune
		//Find vowels, remember positions, write into array of runes v
		for i, r := range runes {
			if r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
				p = student.MyAppend(p, []int{i})
				v = student.MyAppendRune(v, []rune{runes[i]})
			}
		}
		Mirror(v)
		//replace original string vowels with mirrored vowels
		for i := range p {
			runes[p[i]] = v[i]
		}
		for _, r := range runes {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
