package student

import (
	"github.com/01-edu/z01"
)

func ConvertBaseTo(nb int, base int, sequence []int) []int {
	if nb < 0 {
		return ConvertBaseTo(nb/-1, base, sequence)
	} else if nb > base {
		r := nb % base
		sequence = MyAppend([]int{r}, sequence)
		return ConvertBaseTo(nb/base, base, sequence)
	} else {
		r := nb % base
		sequence = MyAppend([]int{r}, sequence)
		sequence = MyAppend([]int{nb / base}, sequence)

		return sequence

	}
}

func PrintNbrBase(nb int, base string) string {
	lenBase := Length(base)
	sequence := []int{}
	overflow := 0
	if Valid(base) == false {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nb < 0 {
			z01.PrintRune('-')
		}
		if nb > 9223372036854775807 || nb < -9223372036854775807 {
			nb = nb - 1
			overflow = 1
		}
		p := ConvertBaseTo(nb, lenBase, sequence)
		base2 := []rune(base)
		if p[0] == 0 {
			p = p[1:]
		}
		pl := IntLen(p)
		p[pl-1] = p[pl-1] + overflow
		var s string
		for _, n := range p {
			s = s + string(base2[n])
		}
		return s
	}
	return ""
}
