package student

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nb int, base string) {
	lenBase := Length(base)
	runeBase := []rune(base)
	if Valid(base) == false { //Valid is a checker func from my aux.go package
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nb < 0 {
			z01.PrintRune('-')
			nb = nb * -1
		}
		for nb != 0 {
			r := nb % lenBase
			if r < 0 {
				r = r * -1
			}
			defer z01.PrintRune(runeBase[r])
			nb = nb / lenBase
		}
	}
}
