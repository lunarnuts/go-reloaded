package student

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	var base = []rune("0123456789")
	var base2 = []rune("0123456789")
	comb := base[:n]
	max := base2[10-n:]
	recursiveComb(comb, max, n)
}
func Check(s []rune) bool { //check for valid comb
	var c int
	for range s {
		c++
	}
	c--
	for i := 0; i < c; i++ {
		if s[i] >= s[i+1] {
			return false
		}
	}
	return true
}
func PrintC(a string) { //func to print to output
	arr := []rune(a)
	for i := range arr {
		z01.PrintRune(arr[i])
	}
}

func recursiveComb(comb []rune, max []rune, n int) bool {
	if string(comb) == string(max) { //base condition
		PrintC(string(comb) + "\n")
		return true
	}
	m := n - 1       //for indexing
	if Check(comb) { //check if suggested comb is valid
		PrintC(string(comb) + ", ")
	}
	for i, _ := range comb { //for each elem of array
		j := m - i
		if comb[j] < max[j] { //if elem exceeds 9
			comb[j]++
			break
		}
		if j > 0 {
			comb[j] = comb[j-1]
		}
	}
	return recursiveComb(comb, max, n)
}
