package student

import (
	"io/ioutil"
	"os"
)

func MyRead() ([]byte, error) {
	for {
		input, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			continue
		} else {
			return input, err
		}
	}
}
func CountWords(str string, charset string) int {
	l := Length(str)
	l2 := Length(charset)
	count := 0
	p := 0
	for n := range str {
		if n != 0 && n != l {
			j := ""
			if n+l2 <= l {
				j = string(str[n : n+l2])
			}
			if j == charset {
				str2 := string(str[p:n])
				if str2 != "" {
					count++
				}
				p = n + l2
			}
		}
	}
	str2 := string(str[p:l])
	if str2 != "" {
		count++
	}
	return count
}

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	power--
	return nb * RecursivePower(nb, power)
}

func Index(s string, c rune) int {
	str := []rune(s)
	index := -1
	for n := range str {
		if str[n] == c {
			index = n
			return index
		}
	}
	return index
}

func Valid(s string) bool {
	if Length(s) < 2 {
		return false
	}
	if (Index(s, '-') != -1) || (Index(s, '+') != -1) {
		return false
	}
	count := 0
	for _, a := range s {
		count++
		if Index(s[count:], a) >= 0 {
			return false
		}
	}
	return true
}

func Length(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func IntLen(str []int) int {
	s := 0
	for range str {
		s++
	}
	return s
}

func MyAppend(a, b []int) []int {
	p := [200]int{-1}
	count := 0
	for range a {
		p[count] = a[count]
		count++
	}
	len2 := IntLen(a)
	for range b {
		p[count] = b[count-len2]
		count++
	}
	return p[:count]
}
func MyAppendRune(a, b []rune) []rune {
	p := [200]rune{-1}
	count := 0
	for range a {
		p[count] = a[count]
		count++
	}
	len2 := 0
	for range a {
		len2++
	}
	for range b {
		p[count] = b[count-len2]
		count++
	}
	return p[:count]
}

func toInt(a rune) int {
	s := 0
	for i := '0'; i < a; i++ {
		s++
	}
	return s
}
func MyMake(f int) []byte {
	var buffer []byte
	for i := 0; i < f; i++ {
		buffer = append(buffer, 0)
	}
	return buffer
}
func Atoi(s string) int {
	num := 0
	sum := 0
	for _, n := range s {
		num = toInt(n)
		sum = sum*10 + num
	}
	return sum
}
func CheckS(s string) bool {
	check := []rune(s)
	for _, n := range check {
		if n > 47 && n < 57 || n == '+' || n == '-' {
			return true
		}
	}
	return false
}
